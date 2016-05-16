package square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentWasFullyRefundedPARTIAL(t *testing.T) {
	p := Payment{Refunds: []Refund{Refund{Type: "PARTIAL"}}}

	assert.False(t, p.WasFullyRefunded())
}

func TestPaymentWasFullyRefundedFULL(t *testing.T) {
	p := Payment{Refunds: []Refund{Refund{Type: "FULL"}}}

	assert.True(t, p.WasFullyRefunded())
}

func TestItemizationQuantity(t *testing.T) {
	i := Itemization{QuantityString: "1.00"}
	assert.Equal(t, 1.00, i.Quantity())
}

func TestItemizationQuantityEmptyDefaultsTo0(t *testing.T) {
	i := Itemization{}
	assert.Equal(t, 0.00, i.Quantity())
}
