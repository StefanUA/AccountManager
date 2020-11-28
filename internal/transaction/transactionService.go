package transaction

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

var transactionService IService

//ReadTransactionFile receives an input file location
//and executes transactions written in the file
func (Service) ReadTransactionFile(inputFile string) ([]*Transaction, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var transactions []*Transaction
	for fileScanner.Scan() {
		line := fileScanner.Text()
		transaction := &Transaction{}

		err = json.Unmarshal([]byte(line), transaction)
		transactions = append(transactions, transaction)
		log.Printf("%v\n", transaction.Time)
	}

	return transactions, nil
}

//ProcessTransactions receives a list of transactions
//and executes transactions written in the file
func (Service) ProcessTransactions(transactions []*Transaction) ([]*Response, error) {

	return nil, nil
}

//WriteTransactionOutput receives a list of transaction responses
//and outputs the data into an output file
func (Service) WriteTransactionOutput(responses []*Response, outputFile string) error {
	return nil
}
