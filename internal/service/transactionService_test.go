package service

import (
	"bytes"
	"encoding/json"
	"github.com/StefanUA/AccountManager/internal/model"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func Test_ProcessTransactions(t *testing.T) {
	assert := assert.New(t)

	ts := &TransactionService{}
	customerService = serviceMock{}
	validInput := []model.TransactionRequest{
		model.TransactionRequest{
			ID:         "testID",
			CustomerID: "testCustID",
			LoadAmount: "1000",
			Time: model.JSONTime{
				Time: time.Now(),
			},
		},
		model.TransactionRequest{
			ID:         "testID-1",
			CustomerID: "testCustID-1",
			LoadAmount: "1005",
			Time: model.JSONTime{
				Time: time.Now().AddDate(1, 0, 0),
			},
		},
	}
	validOutput := model.NewOrderedResponseMap()
	validOutput.Set("testID-testCustID",
		model.TransactionResponse{
			ID:         "testID",
			CustomerID: "testCustID",
			Accepted:   true,
		})
	validOutput.Set("testID-1-testCustID-1",
		model.TransactionResponse{
			ID:         "testID-1",
			CustomerID: "testCustID-1",
			Accepted:   false,
		})

	loadMock = func(request model.TransactionRequest) bool {
		return true
	}
	result := ts.ProcessTransactions(validInput[:0])
	assert.Equal(result, model.NewOrderedResponseMap(), "Empty transactions should return empty response map")

	loadMock = func(request model.TransactionRequest) bool {
		if request.ID == "testID-1" {
			return false
		}
		return true
	}
	result = ts.ProcessTransactions(validInput)
	assert.Equal(result, validOutput, "Valid transactions should return matching responses")
}

func Test_readTransactionFile(t *testing.T) {
	assert := assert.New(t)

	ts := &TransactionService{}
	const layout = "2006-01-02T15:04:05Z07:00"
	transactionTime, err := time.Parse(layout, "2006-01-02T15:04:05Z")
	assert.NoError(err, "Valid time format should be parsed without error")
	validResponse := []model.TransactionRequest{{
		ID:         "testID",
		CustomerID: "testCustID",
		LoadAmount: "$1000.54",
		Time: model.JSONTime{
			Time: transactionTime,
		},
	}}
	jsonBytes, err := json.Marshal(validResponse[0])
	assert.NoError(err, "Valid response should serialize without error")
	validReader := strings.NewReader(string(jsonBytes))
	invalidReader := strings.NewReader("")

	result := ts.readTransactionFile(validReader)
	assert.Equal(result, validResponse, "Valid reader should return transaction content")

	result = ts.readTransactionFile(invalidReader)
	assert.Nil(result, "Invalid reader should return nil response")
}

func Test_writeResponseFile(t *testing.T) {
	assert := assert.New(t)

	ts := &TransactionService{}
	validResponse := model.NewOrderedResponseMap()
	responseItem := model.TransactionResponse{
		ID:         "testID",
		CustomerID: "testCustID",
		Accepted:   false,
	}
	validResponse.Set("testID-testCustID", responseItem)
	assert.Equal(validResponse.Size(), 1)
	var validWriter bytes.Buffer

	jsonBytes, err := json.Marshal(validResponse.GetByIndex(0))
	assert.NoError(err, "Valid response should serialize without error")

	err = ts.writeResponseFile(&validWriter, validResponse)
	assert.NoError(err, "Valid response should write without error")
	assert.Equal(string(jsonBytes)+"\n", validWriter.String(), "Valid writer should print transaction response")
}

//Mocks
var isValidTransactionRequestMock func(model.Customer, model.TransactionRequest) bool
var loadMock func(model.TransactionRequest) bool
var getAmountMock func(string) (float64, bool)
var getWeekKeyMock func(time.Time) string

type serviceMock struct{}

func (sm serviceMock) isValidTransactionRequest(customer model.Customer, request model.TransactionRequest) bool {
	return isValidTransactionRequestMock(customer, request)
}
func (sm serviceMock) Load(request model.TransactionRequest) bool {
	return loadMock(request)
}
func (sm serviceMock) getAmount(cashAmount string) (float64, bool) {
	return getAmountMock(cashAmount)
}
func (sm serviceMock) getWeekKey(inputTime time.Time) string {
	return getWeekKeyMock(inputTime)
}
