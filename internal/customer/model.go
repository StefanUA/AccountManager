package customer

type (
	//IService interface establishes the contract for interacting
	//with cutomers
	IService interface {
		validateTransaction(string) (bool, error)
		Load(Transaction) (bool, error)
	}

	//Service implements the contract for interacting
	//with cutomers
	Service struct{}

	//Customer
	Customer struct {
		CustomerID string
		Weekly
	}
)
