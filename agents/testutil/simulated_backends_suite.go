package testutil

import (
	"testing"

	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

// SimulatedBackendsTestSuite can be used as the base for any test needing simulated backends
// that have an origin, destination and attestation collector and a guard and notary
// added to each.
// TODO (joe): For tests that do not need all 3 simulated backends, allow them to pass in
// flags indicating the subset of backends desired. Some tests might only want
// an attestation collector, others might only want an origin and an attestation collector,
// others might want just a destination, etc.
type SimulatedBackendsTestSuite struct {
	*testsuite.TestSuite
	OriginContract              *originharness.OriginHarnessRef
	DestinationContract         *destinationharness.DestinationHarnessRef
	DestinationContractMetadata contracts.DeployedContract
	AttestationHarness          *attestationharness.AttestationHarnessRef
	AttestationContract         *attestationcollector.AttestationCollectorRef
	AttestationContractMetadata contracts.DeployedContract
	TestBackendOrigin           backends.SimulatedTestBackend
	TestBackendDestination      backends.SimulatedTestBackend
	TestBackendAttestation      backends.SimulatedTestBackend
	NotaryWallet                wallet.Wallet
	GuardWallet                 wallet.Wallet
	NotarySigner                signer.Signer
	GuardSigner                 signer.Signer
	OriginWallet                wallet.Wallet
	DestinationWallet           wallet.Wallet
	AttestationWallet           wallet.Wallet
	OriginSigner                signer.Signer
	DestinationSigner           signer.Signer
	AttestationSigner           signer.Signer
}

// NewSimulatedBackendsTestSuite creates an end-to-end test suite with simulated
// backends set up.
func NewSimulatedBackendsTestSuite(tb testing.TB) *SimulatedBackendsTestSuite {
	tb.Helper()
	return &SimulatedBackendsTestSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// SetupOrigin sets up the backend that will have the origin contract deployed on it.
func (a *SimulatedBackendsTestSuite) SetupOrigin(deployManager *DeployManager) {
	_, a.OriginContract = deployManager.GetOriginHarness(a.GetTestContext(), a.TestBackendOrigin)
	originOwnerPtr, err := a.OriginContract.OriginHarnessCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	originOwnerAuth := a.TestBackendOrigin.GetTxContext(a.GetTestContext(), &originOwnerPtr)

	txOriginNotaryAdd, err := a.OriginContract.AddAgent(originOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendOrigin.WaitForConfirmation(a.GetTestContext(), txOriginNotaryAdd)
	txOriginGuardAdd, err := a.OriginContract.AddAgent(originOwnerAuth.TransactOpts, uint32(0), a.GuardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendOrigin.WaitForConfirmation(a.GetTestContext(), txOriginGuardAdd)
}

// SetupDestination sets up the backend that will have the destination contract deployed on it.
func (a *SimulatedBackendsTestSuite) SetupDestination(deployManager *DeployManager) {
	a.DestinationContractMetadata, a.DestinationContract = deployManager.GetDestinationHarness(a.GetTestContext(), a.TestBackendDestination)

	destOwnerPtr, err := a.DestinationContract.DestinationHarnessCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}

	destOwnerAuth := a.TestBackendDestination.GetTxContext(a.GetTestContext(), &destOwnerPtr)
	txDestinationNotaryAdd, err := a.DestinationContract.AddAgent(destOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendDestination.WaitForConfirmation(a.GetTestContext(), txDestinationNotaryAdd)
	txDestinationGuardAdd, err := a.DestinationContract.AddAgent(destOwnerAuth.TransactOpts, uint32(0), a.GuardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendDestination.WaitForConfirmation(a.GetTestContext(), txDestinationGuardAdd)
}

// SetupAttestation sets up the backend that will have the attestation collector contract deployed on it.
func (a *SimulatedBackendsTestSuite) SetupAttestation(deployManager *DeployManager) {
	_, a.AttestationHarness = deployManager.GetAttestationHarness(a.GetTestContext(), a.TestBackendAttestation)
	a.AttestationContractMetadata, a.AttestationContract = deployManager.GetAttestationCollector(a.GetTestContext(), a.TestBackendAttestation)

	attestOwnerPtr, err := a.AttestationContract.AttestationCollectorCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	attestOwnerAuth := a.TestBackendAttestation.GetTxContext(a.GetTestContext(), &attestOwnerPtr)

	txAddNotary, err := a.AttestationContract.AddAgent(attestOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendAttestation.WaitForConfirmation(a.GetTestContext(), txAddNotary)
	txAddGuard, err := a.AttestationContract.AddAgent(attestOwnerAuth.TransactOpts, uint32(0), a.GuardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendAttestation.WaitForConfirmation(a.GetTestContext(), txAddGuard)
}

// SetupTest sets up the test.
func (a *SimulatedBackendsTestSuite) SetupTest() {
	a.TestSuite.SetupTest()

	a.TestBackendOrigin = preset.GetRinkeby().Geth(a.GetTestContext(), a.T())
	a.TestBackendDestination = preset.GetBSCTestnet().Geth(a.GetTestContext(), a.T())
	a.TestBackendAttestation = preset.GetMaticMumbai().Geth(a.GetTestContext(), a.T())

	var err error
	a.NotaryWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.GuardWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.OriginWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.DestinationWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.AttestationWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}

	a.NotarySigner = localsigner.NewSigner(a.NotaryWallet.PrivateKey())
	a.GuardSigner = localsigner.NewSigner(a.GuardWallet.PrivateKey())

	deployManager := NewDeployManager(a.T())
	a.SetupOrigin(deployManager)
	a.SetupDestination(deployManager)
	a.SetupAttestation(deployManager)
}