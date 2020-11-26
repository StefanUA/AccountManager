package transaction

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"time"
)

type (
	//IService interface establishes the contract for handling
	//transactions
	IService interface {
		ReadTransactionFile(string) error
	}

	//Service implements the contract for handling
	//transactions
	Service struct{}

	//Transaction represents the structure of a requested transaction
	Transaction struct {
		ID         int       `json:"id"`
		CustomerID int       `json:"customer_id"`
		LoadAmount string    `json:"load_amount"`
		Time       time.Time `json:"time"`
	}
)

var transactionService IService

//ReadTransactionFile receives an input file location
//and executes transactions written in the file
func (*Service) ReadTransactionFile(inputFile *string) ([]*Transaction, error) {
	file, err := os.Open(*inputFile)
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

		json.Unmarshal([]byte(line), transaction)
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
