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
		ProcessTransactions([]model.TransactionRequest) model.OrderedResponseMap
		WriteTransactionOutput(model.OrderedResponseMap, string) error
	}

	//TransactionService implements the contract for handling
	//transactions
	TransactionService struct{}
)

var customerService ICustomerService = &CustomerService{}
var osOpen = os.Open
var osCreate = os.Create
var bufioNewScanner = bufio.NewScanner

//ReadTransactionFile receives an input file location
//and executes transactions written in the file
func (ts TransactionService) ReadTransactionFile(inputFile string) ([]model.TransactionRequest, error) {
	file, err := osOpen(inputFile)
	if err != nil {
		log.Fatalf("Error reading file '%s': %v", inputFile, err)
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
func (ts TransactionService) ProcessTransactions(transactionRequests []model.TransactionRequest) model.OrderedResponseMap {
	transactionResponses := model.NewOrderedResponseMap()
	for _, transactionRequest := range transactionRequests {
		transactionResponse := model.NewTransactionResponse(transactionRequest.ID, transactionRequest.CustomerID)
		if existingTransaction := transactionResponses.Get(transactionResponse.ID + "-" + transactionResponse.CustomerID); existingTransaction.ID == "" {
			transactionResponse.Accepted = customerService.Load(transactionRequest)
			transactionResponses.Set(transactionResponse.ID+"-"+transactionResponse.CustomerID, transactionResponse)
		}
	}

	return transactionResponses
}

//WriteTransactionOutput receives a list of transaction responses
//and outputs the data into an output file
func (ts TransactionService) WriteTransactionOutput(responses model.OrderedResponseMap, outputFile string) error {
	file, err := osCreate(outputFile)
	if err != nil {
		log.Fatalf("Error writing file '%s': %v", outputFile, err)
		return err
	}
	defer file.Close()
	return ts.writeResponseFile(file, responses)
}

func (ts TransactionService) writeResponseFile(writer io.Writer, responses model.OrderedResponseMap) error {
	w := bufio.NewWriter(writer)
	var err error
	for i := 0; i < responses.Size(); i++ {
		value := responses.GetByIndex(i)
		jsonBytes, err := json.Marshal(value)
		if err != nil {
			log.Fatalf("Error marshalling input: %v", value)
			jsonBytes = []byte("")
		}
		_, err = w.Write(jsonBytes)
		err = w.WriteByte('\n')
		if err != nil {
			log.Fatalf("Error writing input: %v", value)
		}
	}
	w.Flush()
	return err
}
