package accountmanager

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

//NewCommand creates a new Account manager executable for the cli
func TestNewCommand(t *testing.T) {
	assert := assert.New(t)

	accountManager := NewCommand()

	assert.NotNil(t, accountManager)
	assert.NotEmpty(t, accountManager.Usage)
}

//Execute runs the account manger application
func TestExecuteValidFlag(t *testing.T) {
	assert := assert.New(t)

	accountManager := NewCommand()
	flag.Set("inputFile", "testFile")

	accountManager.Execute()
}
