type (
	//IService interface establishes the contract for handling
	//transactions
	IService interface {
		ReadTransactionFile(string) ([]*Transaction, error)
		ProcessTransactions([]*Transaction) ([]*Response, error)
		WriteTransactionOutput([]*Response, string) error
	}

	//Service implements the contract for handling
	//transactions
	Service struct{}

	//Transaction represents the structure of a requested transaction
	Transaction struct {
		ID         string           `json:"id"`
		CustomerID string           `json:"customer_id"`
		LoadAmount string           `json:"load_amount"`
		Time       structs.JSONTime `json:"time"`
	}

	//Response represents the result of an executed transaction
	Response struct {
		ID         int  `json:"id"`
		CustomerID int  `json:"customer_id"`
		Accepted   bool `json:"accepted"`
	}
)