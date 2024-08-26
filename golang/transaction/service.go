package transaction

type Database interface {
	GetItem(tableName string, id string) (interface{}, bool)
	Save(tableName string, index string, item interface{})
}

type AccountService interface {
	DebitBalance(id string, value float32) bool
	CreditBalance(id string, value float32)
}

type Service struct {
	db Database
	as AccountService
}

func (s *Service) ProcessTransaction(input string) {
	transaction, err := NewTransaction(input)

	if err != nil {
		return
	}

	ok := s.as.DebitBalance(transaction.Origin, transaction.Value)
	if ok {
		s.as.CreditBalance(transaction.Destination, transaction.Value)
		s.db.Save("transactions", transaction.ID.String(), transaction)
	}
}

func NewService(database Database, as AccountService) Service {
	return Service{
		database,
		as,
	}
}
