package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	accountModule "github.com/HenriquePalote/transactions-processing-per-stacks/golang/account"
	"github.com/HenriquePalote/transactions-processing-per-stacks/golang/database"
	transactionModule "github.com/HenriquePalote/transactions-processing-per-stacks/golang/transaction"
)

func processFile(filename string, cb func(line string)) {
	var wg sync.WaitGroup

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		wg.Add(1)
		line := fileScanner.Text()
		go func(line string) {
			fmt.Println(line)
			cb(line)
			wg.Done()
		}(line)
	}
	wg.Wait()
}

func exec() {
	database := database.NewDatabase()

	processFile("./seeds/account.seed.txt", func(line string) {
		account, err := accountModule.NewAccount(line)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		database.AddAccount(account)
	})

	processFile("./seeds/transaction.seed.txt", func(line string) {
		transaction, err := transactionModule.NewTransaction(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		database.AddTransaction(transaction)
	})

	database.Print()
}

func main() {
	exec()
}
