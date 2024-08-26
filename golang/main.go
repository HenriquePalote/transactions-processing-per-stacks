package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/HenriquePalote/transactions-processing-per-stacks/golang/account"
	"github.com/HenriquePalote/transactions-processing-per-stacks/golang/database"
	"github.com/HenriquePalote/transactions-processing-per-stacks/golang/transaction"
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
			cb(line)
			wg.Done()
		}(line)
	}
	wg.Wait()
}

func exec() {
	database := database.NewDatabase()
	accountService := account.NewService(&database)
	transaction := transaction.NewService(&database, accountService)

	processFile("./seeds/account.seed.txt", func(line string) {
		accountService.SeedAccount(line)
	})

	processFile("./seeds/transaction.seed.txt", func(line string) {
		transaction.ProcessTransaction(line)
	})

	database.Print()
}

func main() {
	exec()
}
