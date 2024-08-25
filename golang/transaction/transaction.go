package transaction

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID
	Origin      string
	Destination string
	Value       float32
}

func NewTransaction(transaction string) (Transaction, error) {
	splitedTransaction := strings.Split(transaction, " ")

	value, err := strconv.ParseFloat(splitedTransaction[2], 32)
	if err != nil {
		return Transaction{}, err
	}

	return Transaction{
		ID:          uuid.New(),
		Origin:      splitedTransaction[0],
		Destination: splitedTransaction[1],
		Value:       float32(value),
	}, nil

}
