package backfill

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/scribe/db"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"golang.org/x/sync/errgroup"
)

// ContractBackfiller is a backfiller that fetches logs for a specific contract.
type ContractBackfiller struct {
	// contract is the contract to get logs for
	contract contracts.DeployedContract
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// client is the client for filtering
	client client.EVMClient
	// cache is a cache for txHashes
	cache *lru.Cache
}

// NewContractBackfiller creates a new backfiller for a contract.
func NewContractBackfiller(contract contracts.DeployedContract, eventDB db.EventDB, client client.EVMClient) (*ContractBackfiller, error) {
	// initialize the cache for the txHashes
	cache, err := lru.New(500)
	if err != nil {
		return nil, fmt.Errorf("could not initialize cache: %w", err)
	}

	return &ContractBackfiller{
		contract: contract,
		eventDB:  eventDB,
		client:   client,
		cache:    cache,
	}, nil
}

// Backfill takes in a channel of logs, uses each log to get the receipt from its txHash,
// gets all of the logs from the receipt, then stores the receipt, the logs from the
// receipt, and the last indexed block for hte contract in the EventDB.
//
//nolint:gocognit, cyclop
func (c ContractBackfiller) Backfill(ctx context.Context, givenStart uint64, endHeight uint64) error {
	// initialize the channel for the logs
	startHeight, err := c.StartHeightForBackfill(ctx, givenStart)
	if err != nil {
		return fmt.Errorf("could not get start height: %w", err)
	}
	// in the case of a failed backfill, we want to start from the last indexed block - 1
	if startHeight != 0 {
		startHeight--
	}
	logChan, errChan, doneChan := c.GetLogs(ctx, startHeight, endHeight)
	// start listening for logs
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case log := <-logChan:
				// check if the txHash has already been stored in the cache
				if _, ok := c.cache.Get(log.TxHash); ok {
					continue
				}
				err = c.Store(ctx, log)
				if err != nil {
					return fmt.Errorf("could not store data: %w", err)
				}
			case err := <-errChan:
				return fmt.Errorf("could not get logs: %w", err)
			case <-doneChan:
				return nil
			}
		}
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not backfill contract: %w", err)
	}
	return nil
}

// Store stores the logs, receipts, and transactions for a tx hash.
//
//nolint:cyclop
func (c ContractBackfiller) Store(ctx context.Context, log types.Log) error {
	receipt, err := c.client.TransactionReceipt(ctx, log.TxHash)
	if err != nil {
		return fmt.Errorf("could not get transaction receipt for txHash: %w", err)
	}

	// parallelize storing logs, receipts, and transactions
	g, groupCtx := errgroup.WithContext(ctx)
	if err != nil {
		return fmt.Errorf("could not create errgroup: %w", err)
	}

	g.Go(func() error {
		// get the logs from the receipt and store them in the db
		for _, log := range receipt.Logs {
			if log == nil {
				return fmt.Errorf("log is nil")
			}
			err = c.eventDB.StoreLog(groupCtx, *log, uint32(c.contract.ChainID().Uint64()))
			if err != nil {
				return fmt.Errorf("could not store log: %w", err)
			}
		}
		return nil
	})

	g.Go(func() error {
		// store the receipt in the db
		err = c.eventDB.StoreReceipt(groupCtx, *receipt, uint32(c.contract.ChainID().Uint64()))
		if err != nil {
			return fmt.Errorf("could not store receipt: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		// store the transaction in the db
		txn, isPending, err := c.client.TransactionByHash(groupCtx, receipt.TxHash)
		if err != nil {
			return fmt.Errorf("could not get transaction by hash: %w", err)
		}
		if isPending {
			return fmt.Errorf("transaction is pending")
		}
		err = c.eventDB.StoreEthTx(groupCtx, txn, uint32(c.contract.ChainID().Uint64()))
		if err != nil {
			return fmt.Errorf("could not store transaction: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not store data: %w", err)
	}

	// store the last indexed block in the db
	err = c.eventDB.StoreLastIndexed(ctx, c.contract.Address(), uint32(c.contract.ChainID().Uint64()), receipt.BlockNumber.Uint64())
	if err != nil {
		return fmt.Errorf("could not store last indexed block: %w", err)
	}

	// store the txHash in the cache
	c.cache.Add(log.TxHash, true)

	return nil
}

// chunkSize is how big to make the chunks when fetching.
const chunkSize = 1024

// GetLogs gets all logs for the contract.
func (c ContractBackfiller) GetLogs(ctx context.Context, startHeight, endHeight uint64) (logsChan <-chan types.Log, errsChan <-chan error, completeChan <-chan bool) {
	// initialize the channel
	logChan := make(chan types.Log)
	errChan := make(chan error)
	doneChan := make(chan bool)

	// start the filterer. This filters the range and sends the logs to the logChan.
	rangeFilter := NewRangeFilter(c.contract.Address(), c.client, big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), chunkSize, true)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// start the range filterer, return any errors to an error channel
		err := rangeFilter.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not filter range: %w", err)
		}
		return nil
	})

	// take the logs and put them in the log channel
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case logInfos := <-rangeFilter.GetLogChan():
				for _, log := range logInfos.logs {
					logChan <- log
				}
				if rangeFilter.Done() {
					doneChan <- true
				}
			}
		}
	})

	// return errors to the channel when done filtering
	go func() {
		err := g.Wait()
		if err != nil {
			errChan <- err
		}
	}()

	return logChan, errChan, doneChan
}

// StartHeightForBackfill gets the startHeight for backfilling. This is the maximum
// of the most recent block for the contract and the startHeight given in the config.
func (c ContractBackfiller) StartHeightForBackfill(ctx context.Context, givenStart uint64) (startHeight uint64, err error) {
	lastBlock, err := c.eventDB.RetrieveLastIndexed(ctx, c.contract.Address(), uint32(c.contract.ChainID().Uint64()))
	if err != nil {
		return 0, fmt.Errorf("could not retrieve last indexed block for contract: %w", err)
	}

	if lastBlock > givenStart {
		return lastBlock, nil
	}
	return givenStart, nil
}