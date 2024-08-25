package account

import (
	"strconv"
	"strings"
)

type Account struct {
	Name    string
	Balance float32
}

func NewAccount(account string) (Account, error) {
	splittedAccount := strings.Split(account, " ")

	balance, err := strconv.ParseFloat(splittedAccount[1], 32)

	if err != nil {
		return Account{}, err
	}

	return Account{splittedAccount[0], float32(balance)}, nil
}
