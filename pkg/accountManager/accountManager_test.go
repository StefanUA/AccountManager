package accountmanager

import (
	"errors"
	"testing"

	"github.com/StefanUA/AccountManager/internal/transaction"
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

	readTransactionFileMock = func(s string) ([]*transaction.Transaction, error) {
		return make([]*transaction.Transaction, 0), nil
	}

	err := accountManager.Execute("testFile")
	assert.Nil(err, "Valid file should not return error")

	readTransactionFileMock = func(s string) ([]*transaction.Transaction, error) {
		return nil, errors.New("File does not exist")
	}
	err = accountManager.Execute("testFile")
	assert.Error(err, "Invalid file should return error")

	readTransactionFileMock = func(s string) ([]*transaction.Transaction, error) {
		return nil, errors.New("File does not exist")
	}
	err = accountManager.Execute("testFile")
	assert.Error(err, "Invalid file should return error")
}

//Mocks
var readTransactionFileMock func(string) ([]*transaction.Transaction, error)

type serviceMock struct{}

func (sm serviceMock) ReadTransactionFile(inputFile string) ([]*transaction.Transaction, error) {
	return readTransactionFileMock(inputFile)
}
