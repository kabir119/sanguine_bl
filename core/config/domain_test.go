package config_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/synapse-node/testutils/utils"
)

func domainConfigFixture() config.DomainConfig {
	return config.DomainConfig{
		DomainID:              gofakeit.Uint32(),
		Type:                  types.AllChainTypes()[0].String(),
		RequiredConfirmations: gofakeit.Uint32(),
		HomeAddress:           utils.NewMockAddress().String(),
		RPCUrl:                gofakeit.URL(),
	}
}

func (c ConfigSuite) TestDomainConfigChainType() {
	domainConfig := domainConfigFixture()
	domainConfig.Type = gofakeit.StreetName()

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrInvalidChainType)
}

func (c ConfigSuite) TestDomainConfigID() {
	domainConfig := domainConfigFixture()
	domainConfig.DomainID = 0

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrInvalidDomainID)
}

func (c ConfigSuite) TestXappConfigAddressBlank() {
	domainConfig := domainConfigFixture()
	domainConfig.HomeAddress = ""

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrRequiredField)
}

func (c ConfigSuite) TestXappRPCddressBlank() {
	domainConfig := domainConfigFixture()
	domainConfig.RPCUrl = ""

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrRequiredField)
}

func (c ConfigSuite) TestDomainConfigsDuplicateDomainID() {
	domainConfigA := domainConfigFixture()
	domainConfigB := domainConfigFixture()

	// manually set these to the same id
	domainConfigB.DomainID = domainConfigA.DomainID

	domainConfigs := config.DomainConfigs{
		"a": domainConfigA,
		"b": domainConfigB,
	}

	ok, err := domainConfigs.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrInvalidDomainID)
}
