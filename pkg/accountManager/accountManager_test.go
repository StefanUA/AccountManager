package accountmanager

import (
	"errors"
	"testing"

	"github.com/StefanUA/AccountManager/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	assert := assert.New(t)

	accountManager := NewCommand(serviceMock{})

	assert.NotNil(t, accountManager)
	assert.NotEmpty(t, accountManager.Usage, "Usage should contain valid usage explanation, not empty.")
	assert.NotEmpty(t, accountManager.transactionService, "transactionService should contain a valid service reference when service is passed in, not empty.")
}

func TestExecute(t *testing.T) {
	assert := assert.New(t)

	accountManager := NewCommand(serviceMock{})

	readTransactionFileMock = func(s string) ([]model.TransactionRequest, error) {
		return []model.TransactionRequest{}, nil
	}
	processTransactionsMock = func(s []model.TransactionRequest) map[string]model.TransactionResponse {
		return make(map[string]model.TransactionResponse, 0)
	}
	writeTransactionsMock = func(responses map[string]model.TransactionResponse, outputFile string) error {
		return nil
	}
	err := accountManager.Execute("testFile", "testOutputFile")
	assert.Nil(err, "Valid file should not return error")

	readTransactionFileMock = func(s string) ([]model.TransactionRequest, error) {
		return nil, errors.New("File does not exist")
	}

	err = accountManager.Execute("testFile", "testOutputFile")
	assert.Error(err, "Non-existent file should return error")

	readTransactionFileMock = func(s string) ([]model.TransactionRequest, error) {
		return []model.TransactionRequest{}, nil
	}
	processTransactionsMock = func(s []model.TransactionRequest) map[string]model.TransactionResponse {
		return make(map[string]model.TransactionResponse, 0)
	}
	writeTransactionsMock = func(responses map[string]model.TransactionResponse, outputFile string) error {
		return errors.New("Error writing file")
	}
	err = accountManager.Execute("testFile", "testOutputFile")
	assert.Error(err, "Unsucessful file write should return error")
}

//Mocks
var readTransactionFileMock func(string) ([]model.TransactionRequest, error)
var processTransactionsMock func([]model.TransactionRequest) map[string]model.TransactionResponse
var writeTransactionsMock func(map[string]model.TransactionResponse, string) error

type serviceMock struct{}

func (sm serviceMock) ReadTransactionFile(inputFile string) ([]model.TransactionRequest, error) {
	return readTransactionFileMock(inputFile)
}
func (sm serviceMock) ProcessTransactions(requests []model.TransactionRequest) map[string]model.TransactionResponse {
	return processTransactionsMock(requests)
}
func (sm serviceMock) WriteTransactionOutput(responses map[string]model.TransactionResponse, outputFile string) error {
	return writeTransactionsMock(responses, outputFile)
}
