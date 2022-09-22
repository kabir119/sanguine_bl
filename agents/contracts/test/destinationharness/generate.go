// Package destinationharness generates abi data for contract DestinationHarness.sol
package destinationharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/DestinationHarness.sol --pkg destinationharness --sol-version 0.8.13 --filename destinationharness
// line after go:generate cannot be left blank
