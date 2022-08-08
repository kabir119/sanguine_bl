// autogenerated file

package home

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IHomeTransactor ...
type IHomeTransactor interface {
	// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
	//
	// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns()
	Dispatch(opts *bind.TransactOpts, _destinationDomain uint32, _recipientAddress [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error)
	// ImproperAttestation is a paid mutator transaction binding the contract method 0x88a278ec.
	//
	// Solidity: function improperAttestation(address _updater, bytes _attestation) returns(bool)
	ImproperAttestation(opts *bind.TransactOpts, _updater common.Address, _attestation []byte) (*types.Transaction, error)
	// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
	//
	// Solidity: function initialize(address _updaterManager) returns()
	Initialize(opts *bind.TransactOpts, _updaterManager common.Address) (*types.Transaction, error)
	// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
	//
	// Solidity: function renounceOwnership() returns()
	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)
	// SetSystemMessenger is a paid mutator transaction binding the contract method 0xb7bc563e.
	//
	// Solidity: function setSystemMessenger(address _systemMessenger) returns()
	SetSystemMessenger(opts *bind.TransactOpts, _systemMessenger common.Address) (*types.Transaction, error)
	// SetUpdater is a paid mutator transaction binding the contract method 0x9d54f419.
	//
	// Solidity: function setUpdater(address _updater) returns()
	SetUpdater(opts *bind.TransactOpts, _updater common.Address) (*types.Transaction, error)
	// SetUpdaterManager is a paid mutator transaction binding the contract method 0x9776120e.
	//
	// Solidity: function setUpdaterManager(address _updaterManager) returns()
	SetUpdaterManager(opts *bind.TransactOpts, _updaterManager common.Address) (*types.Transaction, error)
	// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
	//
	// Solidity: function transferOwnership(address newOwner) returns()
	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)
}
