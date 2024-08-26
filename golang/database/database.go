package database

import (
	"fmt"
	"os"
	"sync"
)

type Table struct {
	data map[string]interface{}
}

type Database struct {
	mu     sync.Mutex
	tables map[string]Table
}

func (db *Database) GetItem(tableName string, id string) (interface{}, bool) {
	table, has := db.tables[tableName]
	if !has {
		fmt.Fprintf(os.Stderr, "Table %s doesn't exist \n", table)
		return nil, false
	}

	item, has := table.data[id]

	return item, has
}

func (db *Database) Save(tableName string, index string, item interface{}) {
	db.mu.Lock()
	defer db.mu.Unlock()
	table, has := db.tables[tableName]
	if !has {
		fmt.Fprintf(os.Stderr, "Table %s doesn't exist \n", table)
		return
	}

	table.data[index] = item
}

func (db Database) Print() {
	for tableName, table := range db.tables {
		fmt.Printf("---------%s---------\n", tableName)

		for key, value := range table.data {
			fmt.Printf("%s = %v\n", key, value)
		}
	}
}

func NewDatabase() Database {
	return Database{
		tables: map[string]Table{
			"transactions": Table{data: make(map[string]interface{})},
			"accounts":     Table{data: make(map[string]interface{})},
		},
	}
}
