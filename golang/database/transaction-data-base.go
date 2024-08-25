package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/HenriquePalote/transactions-processing-per-stacks/golang/account"
	"github.com/HenriquePalote/transactions-processing-per-stacks/golang/transaction"
)

type Database struct {
	mu              sync.Mutex
	transactionList []*transaction.Transaction
	accountList     map[string]*account.Account
}

func (db *Database) AddTransaction(transaction transaction.Transaction) {
	db.mu.Lock()
	defer db.mu.Unlock()

	origin, has := db.GetAccount(transaction.Origin)
	if !has {
		fmt.Fprintf(os.Stderr, "Account %s doesn't exist\n", origin.Name)
	}

	if origin.Balance >= transaction.Value {
		origin.Balance -= transaction.Value
	} else {
		fmt.Fprintf(os.Stderr, "Account %s hasn't balance\n", origin.Name)
		return
	}

	destination, has := db.GetAccount(transaction.Destination)
	if !has {
		fmt.Fprintf(os.Stderr, "Account %s doesn't exist\n", origin.Name)
	}

	destination.Balance += transaction.Value

	db.transactionList = append(db.transactionList, &transaction)
}

func (db *Database) AddAccount(account account.Account) {
	db.mu.Lock()
	defer db.mu.Unlock()

	_, has := db.accountList[account.Name]
	if !has {
		db.accountList[account.Name] = &account
	}
}

func (db Database) GetAccount(name string) (*account.Account, bool) {
	account, has := db.accountList[name]
	return account, has
}

func (db Database) GetAccountTransactions(accountName string) []transaction.Transaction {
	list := make([]transaction.Transaction, 0)

	for _, tr := range db.transactionList {
		if tr.Destination == accountName || tr.Origin == accountName {
			list = append(list, *tr)
		}
	}

	return list
}

func (db Database) Print() {
	fmt.Println("ACCOUNTS")
	for _, account := range db.accountList {
		fmt.Println(account)
	}
	fmt.Println("--------------------")
	fmt.Println("TRANSACTIONS")
	for _, transaction := range db.transactionList {
		fmt.Println(transaction)
	}
}

func NewDatabase() Database {
	return Database{
		transactionList: make([]*transaction.Transaction, 0),
		accountList:     make(map[string]*account.Account),
	}
}
