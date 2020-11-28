package model

type (

	//Customer represents the shape of an account capable of holding moneytary value
	Customer struct {
		CustomerID         string
		WeeklyTransactions map[string]WeeklyTransaction
	}

	//WeeklyTransaction represents a customers transaction record for a given week
	WeeklyTransaction struct {
		Total float64
		Days  [7]DailylyTransaction
	}

	//DailylyTransaction represents a customers transaction record for a given day
	DailylyTransaction struct {
		Total float64
		Count int
	}
)

//MaxDailyLimit represents the maximum amount that can be deposited per day
const MaxDailyLimit = 5000.00

//MaxWeeklyLimit represents the maximum amount that can be deposited per week
const MaxWeeklyLimit = 20000.00

//MaxDailyTransationCount represents the maximum amount that can be deposited per day
const MaxDailyTransationCount = 3

//DollarAmountRegex regext to retrieve a valid dollar amount from a string
const DollarAmountRegex = ""

//NewCustomer creates a Customer with all default limits
func NewCustomer(id string) Customer {
	return Customer{
		CustomerID:         id,
		WeeklyTransactions: make(map[string]WeeklyTransaction),
	}
}

//NewWeeklyTransaction creates a WeeklyTransaction
func NewWeeklyTransaction() WeeklyTransaction {
	return WeeklyTransaction{}
}

//NewDailyTransaction creates a DailylyTransaction
func NewDailyTransaction() DailylyTransaction {
	return DailylyTransaction{}
}
