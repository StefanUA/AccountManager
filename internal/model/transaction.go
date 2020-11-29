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
		keys []string
		data map[string]TransactionResponse
	}
)

//NewTransactionResponse initialiazes a new transaction response for a transaction with default accepted value.
//Does not support item removal
func NewTransactionResponse(id string, customerID string) TransactionResponse {
	return TransactionResponse{
		ID:         id,
		CustomerID: customerID,
		Accepted:   false,
	}
}

//NewOrderedResponseMap initialiazes an OrderedResponseMap
func NewOrderedResponseMap() OrderedResponseMap {
	return OrderedResponseMap{
		keys: make([]string, 0),
		data: make(map[string]TransactionResponse, 0),
	}
}

//Set adds a key-value pair
func (orm *OrderedResponseMap) Set(key string, value TransactionResponse) {
	orm.keys = append(orm.keys, key)
	orm.data[key] = value
}

//Get returns a value when the key is passed in
func (orm *OrderedResponseMap) Get(key string) TransactionResponse {
	return orm.data[key]
}

//GetByIndex returns a value when the index is passed in
func (orm *OrderedResponseMap) GetByIndex(index int) TransactionResponse {
	if index < len(orm.keys) {
		key := orm.keys[index]
		return orm.data[key]
	}
	return TransactionResponse{}
}

//Size returns the number of items stored
func (orm *OrderedResponseMap) Size() int {
	return len(orm.keys)
}
