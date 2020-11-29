package service

import (
	"regexp"
	"strconv"
	"time"

	"github.com/StefanUA/AccountManager/internal/model"
)

type (

	//ICustomerService interface establishes the contract for interacting
	//with cutomers
	ICustomerService interface {
		isValidTransactionRequest(model.Customer, model.TransactionRequest) bool
		Load(model.TransactionRequest) bool
	}

	//CustomerService implements the contract for interacting
	//with cutomers
	CustomerService struct{}
)

var customers = make(map[string]model.Customer)

func (cs CustomerService) isValidTransactionRequest(customer model.Customer, transactionRequest model.TransactionRequest) bool {
	result := true

	weekKey := cs.getWeekKey(transactionRequest.Time.Time)
	weeklyTransaction, _ := customer.WeeklyTransactions[weekKey]
	amount, isValidAmount := cs.getAmount(transactionRequest.LoadAmount)

	if !isValidAmount {
		result = false
	} else if (weeklyTransaction.Total + amount) > model.MaxWeeklyLimit {
		result = false
	} else if dayTransaction := weeklyTransaction.Days[transactionRequest.Time.Time.Day()]; dayTransaction.Total+amount > model.MaxDailyLimit {
		result = false
	} else if dayTransaction := weeklyTransaction.Days[transactionRequest.Time.Time.Day()]; dayTransaction.Count+1 > model.MaxDailyTransationCount {
		result = false
	}

	return result
}

func (cs CustomerService) getWeekKey(inputTime time.Time) string {
	year, week := inputTime.ISOWeek()
	return strconv.Itoa(year) + "-" + strconv.Itoa(week)
}

func (cs CustomerService) getAmount(cashAmount string) (float64, bool) {
	regx := regexp.MustCompile(`^\$(([1-9]\d*)?\d)(\.\d\d)?$`)
	match := regx.FindStringSubmatch(cashAmount)
	value, err := strconv.ParseFloat(match[1], 64)
	if cap(match) != 1 && err != nil {
		return 0, false
	}
	return value, true
}

//Load takes in a transaction and performs it against a customer account
//returns a boolean to represent the successful execution of the transaction
func (cs CustomerService) Load(transactionRequest model.TransactionRequest) bool {
	customer, ok := customers[transactionRequest.CustomerID]
	result := false

	if !ok {
		customer = model.NewCustomer(transactionRequest.ID)
		customers[customer.CustomerID] = customer
	}

	if cs.isValidTransactionRequest(customer, transactionRequest) {
		weekKey := cs.getWeekKey(transactionRequest.Time.Time)
		weeklyTransaction, weekExists := customer.WeeklyTransactions[weekKey]
		if !weekExists {
			weeklyTransaction = model.NewWeeklyTransaction()
			customer.WeeklyTransactions[weekKey] = weeklyTransaction
		}
		amount, _ := cs.getAmount(transactionRequest.LoadAmount)
		weeklyTransaction.Total += amount
		dailyTransaction := &weeklyTransaction.Days[transactionRequest.Time.Time.Day()]
		dailyTransaction.Count++
		dailyTransaction.Total += amount
		customer.WeeklyTransactions[weekKey] = model.NewWeeklyTransaction()
		customers[transactionRequest.CustomerID] = customer
		result = true
	}

	return result
}
