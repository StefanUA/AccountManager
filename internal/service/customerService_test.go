package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getAmount(t *testing.T) {
	assert := assert.New(t)

	var cs ICustomerService = &CustomerService{}
	testAmounts := []struct {
		errorMsg          string
		amounts           []string
		response          []float64
		boolResponseValue bool
	}{
		{
			amounts:           []string{"$300", "$0.35", "$5,000.09", "$8942.24", "$0"},
			response:          []float64{300.0, 0.35, 5000.09, 8942.24, 0},
			boolResponseValue: true,
			errorMsg:          "Valid dollar amount '%s' should return numeric value '%f'",
		},
		{
			amounts:           []string{"text", "$$300", "$$0300", "$0300", "$0.3", "$5,00.09", "$8942.424"},
			response:          []float64{0, 0, 0, 0, 0, 0, 0},
			errorMsg:          "Invalid dollar amount '%s' should return numeric value '0'",
			boolResponseValue: false,
		},
	}

	for _, test := range testAmounts {
		for i, amount := range test.amounts {
			result, ok := cs.getAmount(amount)
			assert.Equal(test.response[i], result, fmt.Sprintf(test.errorMsg, amount, test.response[i]))
			assert.Equal(test.boolResponseValue, ok, fmt.Sprintf(test.errorMsg, amount, test.response[i]))
		}
	}
}

func Test_Load(t *testing.T) {

}
