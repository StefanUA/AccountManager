package accountmanager

import "github.com/StefanUA/AccountManager/internal/service"

// AccountManager is the main application for running the account manager
type AccountManager struct {
	Usage              string
	transactionService service.ITransactionService
}

//NewCommand creates a new Account manager executable for the cli
func NewCommand(transactionService service.ITransactionService) AccountManager {
	accountManager := AccountManager{
		Usage: `Account manager takes an input file of transactions and executes each transaction against banking rules. The output is the result of each inputed transaction. 
		Each formatted like so: {"id":"15887","customer_id":"528","load_amount":"$3318.47", "time":"2000-01-01T00:00:00Z"}
		
		Input:
		- inputFile: file to be processed`,
		transactionService: transactionService,
	}

	return accountManager
}

//Execute runs the account manger application
func (am AccountManager) Execute(intputFile string, outputFile string) error {
	transactions, err := am.transactionService.ReadTransactionFile(intputFile)
	if err != nil {
		return err
	}

	responses := am.transactionService.ProcessTransactions(transactions)

	err = am.transactionService.WriteTransactionOutput(responses, outputFile)
	if err != nil {
		return err
	}
	return nil
}
