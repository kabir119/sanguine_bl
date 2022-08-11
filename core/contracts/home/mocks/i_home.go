// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"

	event "github.com/ethereum/go-ethereum/event"

	home "github.com/synapsecns/sanguine/core/contracts/home"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// IHome is an autogenerated mock type for the IHome type
type IHome struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *IHome) Address() common.Address {
	ret := _m.Called()

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// AllGuards provides a mock function with given fields: opts
func (_m *IHome) AllGuards(opts *bind.CallOpts) ([]common.Address, error) {
	ret := _m.Called(opts)

	var r0 []common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) []common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllNotaries provides a mock function with given fields: opts
func (_m *IHome) AllNotaries(opts *bind.CallOpts) ([]common.Address, error) {
	ret := _m.Called(opts)

	var r0 []common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) []common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Count provides a mock function with given fields: opts
func (_m *IHome) Count(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Dispatch provides a mock function with given fields: opts, _destinationDomain, _recipientAddress, _optimisticSeconds, _tips, _messageBody
func (_m *IHome) Dispatch(opts *bind.TransactOpts, _destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, _destinationDomain, _recipientAddress, _optimisticSeconds, _tips, _messageBody)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, uint32, [32]byte, uint32, []byte, []byte) *types.Transaction); ok {
		r0 = rf(opts, _destinationDomain, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, uint32, [32]byte, uint32, []byte, []byte) error); ok {
		r1 = rf(opts, _destinationDomain, _recipientAddress, _optimisticSeconds, _tips, _messageBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterDispatch provides a mock function with given fields: opts, messageHash, leafIndex, destinationAndNonce
func (_m *IHome) FilterDispatch(opts *bind.FilterOpts, messageHash [][32]byte, leafIndex []*big.Int, destinationAndNonce []uint64) (*home.HomeDispatchIterator, error) {
	ret := _m.Called(opts, messageHash, leafIndex, destinationAndNonce)

	var r0 *home.HomeDispatchIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, [][32]byte, []*big.Int, []uint64) *home.HomeDispatchIterator); ok {
		r0 = rf(opts, messageHash, leafIndex, destinationAndNonce)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeDispatchIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, [][32]byte, []*big.Int, []uint64) error); ok {
		r1 = rf(opts, messageHash, leafIndex, destinationAndNonce)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterDomainNotaryAdded provides a mock function with given fields: opts
func (_m *IHome) FilterDomainNotaryAdded(opts *bind.FilterOpts) (*home.HomeDomainNotaryAddedIterator, error) {
	ret := _m.Called(opts)

	var r0 *home.HomeDomainNotaryAddedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *home.HomeDomainNotaryAddedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeDomainNotaryAddedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterDomainNotaryRemoved provides a mock function with given fields: opts
func (_m *IHome) FilterDomainNotaryRemoved(opts *bind.FilterOpts) (*home.HomeDomainNotaryRemovedIterator, error) {
	ret := _m.Called(opts)

	var r0 *home.HomeDomainNotaryRemovedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *home.HomeDomainNotaryRemovedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeDomainNotaryRemovedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterGuardAdded provides a mock function with given fields: opts
func (_m *IHome) FilterGuardAdded(opts *bind.FilterOpts) (*home.HomeGuardAddedIterator, error) {
	ret := _m.Called(opts)

	var r0 *home.HomeGuardAddedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *home.HomeGuardAddedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeGuardAddedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterGuardRemoved provides a mock function with given fields: opts
func (_m *IHome) FilterGuardRemoved(opts *bind.FilterOpts) (*home.HomeGuardRemovedIterator, error) {
	ret := _m.Called(opts)

	var r0 *home.HomeGuardRemovedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *home.HomeGuardRemovedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeGuardRemovedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterImproperAttestation provides a mock function with given fields: opts
func (_m *IHome) FilterImproperAttestation(opts *bind.FilterOpts) (*home.HomeImproperAttestationIterator, error) {
	ret := _m.Called(opts)

	var r0 *home.HomeImproperAttestationIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *home.HomeImproperAttestationIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeImproperAttestationIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterInitialized provides a mock function with given fields: opts
func (_m *IHome) FilterInitialized(opts *bind.FilterOpts) (*home.HomeInitializedIterator, error) {
	ret := _m.Called(opts)

	var r0 *home.HomeInitializedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *home.HomeInitializedIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeInitializedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterNewUpdaterManager provides a mock function with given fields: opts
func (_m *IHome) FilterNewUpdaterManager(opts *bind.FilterOpts) (*home.HomeNewUpdaterManagerIterator, error) {
	ret := _m.Called(opts)

	var r0 *home.HomeNewUpdaterManagerIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts) *home.HomeNewUpdaterManagerIterator); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeNewUpdaterManagerIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterOwnershipTransferred provides a mock function with given fields: opts, previousOwner, newOwner
func (_m *IHome) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*home.HomeOwnershipTransferredIterator, error) {
	ret := _m.Called(opts, previousOwner, newOwner)

	var r0 *home.HomeOwnershipTransferredIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *home.HomeOwnershipTransferredIterator); ok {
		r0 = rf(opts, previousOwner, newOwner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeOwnershipTransferredIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, previousOwner, newOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterUpdaterSlashed provides a mock function with given fields: opts, updater, reporter
func (_m *IHome) FilterUpdaterSlashed(opts *bind.FilterOpts, updater []common.Address, reporter []common.Address) (*home.HomeUpdaterSlashedIterator, error) {
	ret := _m.Called(opts, updater, reporter)

	var r0 *home.HomeUpdaterSlashedIterator
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []common.Address, []common.Address) *home.HomeUpdaterSlashedIterator); ok {
		r0 = rf(opts, updater, reporter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeUpdaterSlashedIterator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, updater, reporter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGuard provides a mock function with given fields: opts, _index
func (_m *IHome) GetGuard(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	ret := _m.Called(opts, _index)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) common.Address); ok {
		r0 = rf(opts, _index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, _index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotary provides a mock function with given fields: opts, _index
func (_m *IHome) GetNotary(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	ret := _m.Called(opts, _index)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) common.Address); ok {
		r0 = rf(opts, _index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, _index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GuardsAmount provides a mock function with given fields: opts
func (_m *IHome) GuardsAmount(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HistoricalRoots provides a mock function with given fields: opts, arg0
func (_m *IHome) HistoricalRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	ret := _m.Called(opts, arg0)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) [32]byte); ok {
		r0 = rf(opts, arg0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, arg0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImproperAttestation provides a mock function with given fields: opts, _attestation
func (_m *IHome) ImproperAttestation(opts *bind.TransactOpts, _attestation []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, _attestation)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []byte) *types.Transaction); ok {
		r0 = rf(opts, _attestation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []byte) error); ok {
		r1 = rf(opts, _attestation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: opts, _updaterManager
func (_m *IHome) Initialize(opts *bind.TransactOpts, _updaterManager common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, _updaterManager)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, _updaterManager)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, _updaterManager)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LocalDomain provides a mock function with given fields: opts
func (_m *IHome) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	ret := _m.Called(opts)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint32); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MAXMESSAGEBODYBYTES provides a mock function with given fields: opts
func (_m *IHome) MAXMESSAGEBODYBYTES(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Nonce provides a mock function with given fields: opts
func (_m *IHome) Nonce(opts *bind.CallOpts) (uint32, error) {
	ret := _m.Called(opts)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint32); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotariesAmount provides a mock function with given fields: opts
func (_m *IHome) NotariesAmount(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Owner provides a mock function with given fields: opts
func (_m *IHome) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseDispatch provides a mock function with given fields: log
func (_m *IHome) ParseDispatch(log types.Log) (*home.HomeDispatch, error) {
	ret := _m.Called(log)

	var r0 *home.HomeDispatch
	if rf, ok := ret.Get(0).(func(types.Log) *home.HomeDispatch); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeDispatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseDomainNotaryAdded provides a mock function with given fields: log
func (_m *IHome) ParseDomainNotaryAdded(log types.Log) (*home.HomeDomainNotaryAdded, error) {
	ret := _m.Called(log)

	var r0 *home.HomeDomainNotaryAdded
	if rf, ok := ret.Get(0).(func(types.Log) *home.HomeDomainNotaryAdded); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeDomainNotaryAdded)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseDomainNotaryRemoved provides a mock function with given fields: log
func (_m *IHome) ParseDomainNotaryRemoved(log types.Log) (*home.HomeDomainNotaryRemoved, error) {
	ret := _m.Called(log)

	var r0 *home.HomeDomainNotaryRemoved
	if rf, ok := ret.Get(0).(func(types.Log) *home.HomeDomainNotaryRemoved); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeDomainNotaryRemoved)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseGuardAdded provides a mock function with given fields: log
func (_m *IHome) ParseGuardAdded(log types.Log) (*home.HomeGuardAdded, error) {
	ret := _m.Called(log)

	var r0 *home.HomeGuardAdded
	if rf, ok := ret.Get(0).(func(types.Log) *home.HomeGuardAdded); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeGuardAdded)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseGuardRemoved provides a mock function with given fields: log
func (_m *IHome) ParseGuardRemoved(log types.Log) (*home.HomeGuardRemoved, error) {
	ret := _m.Called(log)

	var r0 *home.HomeGuardRemoved
	if rf, ok := ret.Get(0).(func(types.Log) *home.HomeGuardRemoved); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeGuardRemoved)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseImproperAttestation provides a mock function with given fields: log
func (_m *IHome) ParseImproperAttestation(log types.Log) (*home.HomeImproperAttestation, error) {
	ret := _m.Called(log)

	var r0 *home.HomeImproperAttestation
	if rf, ok := ret.Get(0).(func(types.Log) *home.HomeImproperAttestation); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeImproperAttestation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInitialized provides a mock function with given fields: log
func (_m *IHome) ParseInitialized(log types.Log) (*home.HomeInitialized, error) {
	ret := _m.Called(log)

	var r0 *home.HomeInitialized
	if rf, ok := ret.Get(0).(func(types.Log) *home.HomeInitialized); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeInitialized)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseNewUpdaterManager provides a mock function with given fields: log
func (_m *IHome) ParseNewUpdaterManager(log types.Log) (*home.HomeNewUpdaterManager, error) {
	ret := _m.Called(log)

	var r0 *home.HomeNewUpdaterManager
	if rf, ok := ret.Get(0).(func(types.Log) *home.HomeNewUpdaterManager); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeNewUpdaterManager)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseOwnershipTransferred provides a mock function with given fields: log
func (_m *IHome) ParseOwnershipTransferred(log types.Log) (*home.HomeOwnershipTransferred, error) {
	ret := _m.Called(log)

	var r0 *home.HomeOwnershipTransferred
	if rf, ok := ret.Get(0).(func(types.Log) *home.HomeOwnershipTransferred); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeOwnershipTransferred)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseUpdaterSlashed provides a mock function with given fields: log
func (_m *IHome) ParseUpdaterSlashed(log types.Log) (*home.HomeUpdaterSlashed, error) {
	ret := _m.Called(log)

	var r0 *home.HomeUpdaterSlashed
	if rf, ok := ret.Get(0).(func(types.Log) *home.HomeUpdaterSlashed); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*home.HomeUpdaterSlashed)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Parser provides a mock function with given fields:
func (_m *IHome) Parser() home.Parser {
	ret := _m.Called()

	var r0 home.Parser
	if rf, ok := ret.Get(0).(func() home.Parser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(home.Parser)
		}
	}

	return r0
}

// RenounceOwnership provides a mock function with given fields: opts
func (_m *IHome) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Root provides a mock function with given fields: opts
func (_m *IHome) Root(opts *bind.CallOpts) ([32]byte, error) {
	ret := _m.Called(opts)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) [32]byte); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetSystemMessenger provides a mock function with given fields: opts, _systemMessenger
func (_m *IHome) SetSystemMessenger(opts *bind.TransactOpts, _systemMessenger common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, _systemMessenger)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, _systemMessenger)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, _systemMessenger)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetUpdater provides a mock function with given fields: opts, _updater
func (_m *IHome) SetUpdater(opts *bind.TransactOpts, _updater common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, _updater)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, _updater)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, _updater)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetUpdaterManager provides a mock function with given fields: opts, _updaterManager
func (_m *IHome) SetUpdaterManager(opts *bind.TransactOpts, _updaterManager common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, _updaterManager)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, _updaterManager)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, _updaterManager)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// State provides a mock function with given fields: opts
func (_m *IHome) State(opts *bind.CallOpts) (uint8, error) {
	ret := _m.Called(opts)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint8); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestUpdate provides a mock function with given fields: opts
func (_m *IHome) SuggestUpdate(opts *bind.CallOpts) (struct {
	Nonce uint32
	Root  [32]byte
}, error) {
	ret := _m.Called(opts)

	var r0 struct {
		Nonce uint32
		Root  [32]byte
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) struct {
		Nonce uint32
		Root  [32]byte
	}); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(struct {
			Nonce uint32
			Root  [32]byte
		})
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemMessenger provides a mock function with given fields: opts
func (_m *IHome) SystemMessenger(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferOwnership provides a mock function with given fields: opts, newOwner
func (_m *IHome) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, newOwner)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address) *types.Transaction); ok {
		r0 = rf(opts, newOwner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address) error); ok {
		r1 = rf(opts, newOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tree provides a mock function with given fields: opts
func (_m *IHome) Tree(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdaterManager provides a mock function with given fields: opts
func (_m *IHome) UpdaterManager(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VERSION provides a mock function with given fields: opts
func (_m *IHome) VERSION(opts *bind.CallOpts) (uint8, error) {
	ret := _m.Called(opts)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) uint8); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchDispatch provides a mock function with given fields: opts, sink, messageHash, leafIndex, destinationAndNonce
func (_m *IHome) WatchDispatch(opts *bind.WatchOpts, sink chan<- *home.HomeDispatch, messageHash [][32]byte, leafIndex []*big.Int, destinationAndNonce []uint64) (event.Subscription, error) {
	ret := _m.Called(opts, sink, messageHash, leafIndex, destinationAndNonce)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *home.HomeDispatch, [][32]byte, []*big.Int, []uint64) event.Subscription); ok {
		r0 = rf(opts, sink, messageHash, leafIndex, destinationAndNonce)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *home.HomeDispatch, [][32]byte, []*big.Int, []uint64) error); ok {
		r1 = rf(opts, sink, messageHash, leafIndex, destinationAndNonce)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchDomainNotaryAdded provides a mock function with given fields: opts, sink
func (_m *IHome) WatchDomainNotaryAdded(opts *bind.WatchOpts, sink chan<- *home.HomeDomainNotaryAdded) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *home.HomeDomainNotaryAdded) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *home.HomeDomainNotaryAdded) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchDomainNotaryRemoved provides a mock function with given fields: opts, sink
func (_m *IHome) WatchDomainNotaryRemoved(opts *bind.WatchOpts, sink chan<- *home.HomeDomainNotaryRemoved) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *home.HomeDomainNotaryRemoved) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *home.HomeDomainNotaryRemoved) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchGuardAdded provides a mock function with given fields: opts, sink
func (_m *IHome) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *home.HomeGuardAdded) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *home.HomeGuardAdded) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *home.HomeGuardAdded) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchGuardRemoved provides a mock function with given fields: opts, sink
func (_m *IHome) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *home.HomeGuardRemoved) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *home.HomeGuardRemoved) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *home.HomeGuardRemoved) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchImproperAttestation provides a mock function with given fields: opts, sink
func (_m *IHome) WatchImproperAttestation(opts *bind.WatchOpts, sink chan<- *home.HomeImproperAttestation) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *home.HomeImproperAttestation) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *home.HomeImproperAttestation) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInitialized provides a mock function with given fields: opts, sink
func (_m *IHome) WatchInitialized(opts *bind.WatchOpts, sink chan<- *home.HomeInitialized) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *home.HomeInitialized) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *home.HomeInitialized) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchNewUpdaterManager provides a mock function with given fields: opts, sink
func (_m *IHome) WatchNewUpdaterManager(opts *bind.WatchOpts, sink chan<- *home.HomeNewUpdaterManager) (event.Subscription, error) {
	ret := _m.Called(opts, sink)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *home.HomeNewUpdaterManager) event.Subscription); ok {
		r0 = rf(opts, sink)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *home.HomeNewUpdaterManager) error); ok {
		r1 = rf(opts, sink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchOwnershipTransferred provides a mock function with given fields: opts, sink, previousOwner, newOwner
func (_m *IHome) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *home.HomeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, previousOwner, newOwner)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *home.HomeOwnershipTransferred, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, previousOwner, newOwner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *home.HomeOwnershipTransferred, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, previousOwner, newOwner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchUpdaterSlashed provides a mock function with given fields: opts, sink, updater, reporter
func (_m *IHome) WatchUpdaterSlashed(opts *bind.WatchOpts, sink chan<- *home.HomeUpdaterSlashed, updater []common.Address, reporter []common.Address) (event.Subscription, error) {
	ret := _m.Called(opts, sink, updater, reporter)

	var r0 event.Subscription
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *home.HomeUpdaterSlashed, []common.Address, []common.Address) event.Subscription); ok {
		r0 = rf(opts, sink, updater, reporter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *home.HomeUpdaterSlashed, []common.Address, []common.Address) error); ok {
		r1 = rf(opts, sink, updater, reporter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
