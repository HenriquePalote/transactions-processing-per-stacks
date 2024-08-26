package account

import (
	"fmt"
	"os"
)

type Database interface {
	GetItem(tableName string, id string) (interface{}, bool)
	Save(tableName string, index string, item interface{})
}

type Service struct {
	db Database
}

func (s Service) SeedAccount(accountInput string) {
	account, err := NewAccount(accountInput)

	if err != nil {
		return
	}
	s.db.Save("accounts", account.Name, account)
}

func (s Service) GetAccount(name string) (Account, bool) {
	item, _ := s.db.GetItem("accounts", name)
	account, ok := item.(Account)

	if ok {
		return account, true
	}

	return Account{}, false
}

func (s Service) DebitBalance(id string, value float32) bool {
	account, has := s.GetAccount(id)

	if !has {
		fmt.Fprintf(os.Stderr, "Account %s doesn't exist\n", id)
		return false
	}

	if account.Balance >= value {
		account.Balance -= value
		s.db.Save("accounts", account.Name, account)
		return true
	} else {
		fmt.Fprintf(os.Stderr, "Account %s hasn't balance\n", account.Name)
		return false
	}
}

func (s Service) CreditBalance(id string, value float32) {
	account, has := s.GetAccount(id)

	if !has {
		fmt.Fprintf(os.Stderr, "Account %s doesn't exist\n", id)
	}

	account.Balance += value
	s.db.Save("accounts", account.Name, account)
}

func NewService(db Database) Service {
	return Service{
		db,
	}
}
