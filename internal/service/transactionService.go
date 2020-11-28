package service

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/StefanUA/AccountManager/internal/model"
)

type (
	//ITransactionService interface establishes the contract for handling
	//transactions
	ITransactionService interface {
		ReadTransactionFile(string) ([]model.TransactionRequest, error)
		ProcessTransactions([]model.TransactionRequest) map[string]model.TransactionResponse
		WriteTransactionOutput(map[string]model.TransactionResponse, string) error
	}

	//TransactionService implements the contract for handling
	//transactions
	TransactionService struct{}
)

var customerService ICustomerService = CustomerService{}
var osOpen = os.Open
var bufioNewScanner = bufio.NewScanner

//ReadTransactionFile receives an input file location
//and executes transactions written in the file
func (ts TransactionService) ReadTransactionFile(inputFile string) ([]model.TransactionRequest, error) {
	file, err := osOpen(inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return nil, err
	}
	defer file.Close()
	return ts.readTransactionFile(file), err
}

func (ts TransactionService) readTransactionFile(reader io.Reader) []model.TransactionRequest {
	fileScanner := bufioNewScanner(reader)

	var transactionRequests []model.TransactionRequest
	for fileScanner.Scan() {
		line := fileScanner.Text()
		transactionRequest := &model.TransactionRequest{}

		json.Unmarshal([]byte(line), transactionRequest)
		transactionRequests = append(transactionRequests, *transactionRequest)
	}
	return transactionRequests
}

//ProcessTransactions receives a list of transactions
//and executes transactions written in the file
func (ts TransactionService) ProcessTransactions(transactionRequests []model.TransactionRequest) map[string]model.TransactionResponse {
	transactionResponses := make(map[string]model.TransactionResponse)
	for _, transactionRequest := range transactionRequests {
		transactionResponse := model.NewTransactionResponse(transactionRequest.ID, transactionRequest.CustomerID)
		if _, ok := transactionResponses[transactionResponse.ID]; !ok {
			transactionResponse.Accepted = customerService.Load(transactionRequest)
		} else {
			transactionResponse.Accepted = false
		}
		transactionResponses[transactionResponse.ID] = transactionResponse
	}
	log.Printf("%v\n", transactionResponses)

	return transactionResponses
}

//WriteTransactionOutput receives a list of transaction responses
//and outputs the data into an output file
func (ts TransactionService) WriteTransactionOutput(responses map[string]model.TransactionResponse, outputFile string) error {
	return nil
}
