package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPrices(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TEST") == "" {
		t.Skip("run with RUN_INTEGRATION_TEST=1 go test")
	}

	api := NewElPrisenLigeNuAPI(nil)
	prices, err := api.GetPrices(time.Now())
	assert.NoError(t, err)
	assert.NotEmpty(t, prices)
}
