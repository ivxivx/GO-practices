package number

import (
	"math/big"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_deciamlToFloat(t *testing.T) {
	t.Parallel()

	dec := decimal.RequireFromString("1.234")

	flt, accuracy := dec.BigFloat().Float64()

	assert.Equal(t, big.Below, accuracy)

	newDec := decimal.NewFromFloat(flt)

	assert.Equal(t, dec, newDec)
}
