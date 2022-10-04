// Package internal contains dev dependencies required for running developer tasks (coverage testing, etc)
// that are not required by the project itself. In order to enforce this constraint, this module panics upon
// being imported. Dependencies here are not included in produced binaries and won't affect the dev build
//
// test-data is the synapse-contracts repo, used for testing deployments
package internal

import "github.com/dgraph-io/ristretto"

func init() {
	panic("could not import dev package: this package is meant to define dependencies, not be imported.")
}

// required by mockery.
var _ = ristretto.Config{}