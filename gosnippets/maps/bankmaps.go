package main

import (
	"fmt"
)

func createBankMap(account string, credit float32, debit float32) map[string]float32 {
	bankMap := make(map[string]float32)

	credit = credit - debit
	bankMap[account] = credit

	return bankMap
}

func main() {
	account := "4567"
	var accountNumber string

	myWallet := createBankMap(account, 100, 50)
	for k, _ := range myWallet {
		accountNumber = k
	}
	fmt.Printf("On my bank account no. %x, I have %v Euro", accountNumber, myWallet[account])
}
