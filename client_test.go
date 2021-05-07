package mpmapi

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIClient_GetLineItems(t *testing.T) {
	assert := assert.New(t)
	c := NewClient(os.Getenv("MOPUB_PM_API_KEY"))
	lineItems, err := c.GetAllLineItems()
	assert.NoError(err)
	assert.NotEmpty(lineItems)
	assert.Greater(len(lineItems), 2000)
}

func TestAPIClient_GetAdUnits(t *testing.T) {
	assert := assert.New(t)
	c := NewClient(os.Getenv("MOPUB_PM_API_KEY"))
	adUnits, err := c.GetAllAdUnits()
	assert.NoError(err)
	assert.NotEmpty(adUnits)
	assert.Greater(len(adUnits), 40)
}
