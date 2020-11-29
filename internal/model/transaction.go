package model

type (

	//TransactionRequest represents the structure of a requested transaction
	TransactionRequest struct {
		ID         string   `json:"id"`
		CustomerID string   `json:"customer_id"`
		LoadAmount string   `json:"load_amount"`
		Time       JSONTime `json:"time"`
	}

	//TransactionResponse represents the result of an executed transaction
	TransactionResponse struct {
		ID         string `json:"id"`
		CustomerID string `json:"customer_id"`
		Accepted   bool   `json:"accepted"`
	}

	//OrderedResponseMap stores key-value pairs in the order in which they are entered
	OrderedResponseMap struct {
		keys   []string
		values []TransactionResponse
	}
)

//NewTransactionResponse initialiazes a new transaction response for a transaction with default accepted value
func NewTransactionResponse(id string, customerID string) TransactionResponse {
	return TransactionResponse{
		ID:         id,
		CustomerID: customerID,
		Accepted:   false,
	}
}
