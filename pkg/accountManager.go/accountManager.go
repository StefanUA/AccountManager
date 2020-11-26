package accountmanager

import (
	"flag"
	"os"

	"github.com/StefanUA/AccountManager/internal/transaction"
)

// AccountManager is the main application for running the account manager
type AccountManager struct {
	Usage string
	// Description string
}

//NewCommand creates a new Account manager executable for the cli
func NewCommand() *AccountManager {
	accountManager := &AccountManager{
		Usage: `Account manager takes an input file of transactions and executes each transaction against banking rules. The output is the result of each inputed transaction. 
		Each formatted like so: {"id":"15887","customer_id":"528","load_amount":"$3318.47", "time":"2000-01-01T00:00:00Z"}
		
		Input:
		- inputFile: file to be processed`}

	return accountManager
}

//Execute runs the account manger application
func (*AccountManager) Execute() {
	inputFilePtr := flag.String("input", "", "Input file to process (Required)")
	flag.Parse()

	if *inputFilePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	transaction.ProcessTransactionFile(inputFilePtr)
}
