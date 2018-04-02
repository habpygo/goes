package main

import (
	"TheGoProgrammingLanguage/gosnippets/banking/bankaccounts"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// open the file or, if it doesn't exist, create one with that name
	jsonLedger, err := os.OpenFile("bankledger.json", os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("could not decode bytefile: ", err)
	}

	defer jsonLedger.Close()

	// read the opened xmlFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonLedger)

	// create the ledger variable
	var cl bankaccounts.CustodianLedger
	json.Unmarshal(byteValue, &cl)

	var newAccount *bankaccounts.Bankaccount
	/*
		Use newAccount(s) below to test the application. First add them one by one
		Then check if double accounts will be added or not
	*/

	//newAccount = bankaccounts.NewBankAccount("BTCcom", 5378, "HerenStreet", "17a", "Amsterdam")
	//newAccount = bankaccounts.NewBankAccount("Harry Boer", 10480, "Friesestraat", "1143", "Amersfoort")
	//newAccount = bankaccounts.NewBankAccount("Pipo de Clown", 1234, "Clown Street", "34", "New York")
	///newAccount = bankaccounts.NewBankAccount("Mama Lou", 3478, "Lou Street", "117", "Melbourne")
	//newAccount = bankaccounts.NewBankAccount("Kees van Dalen", 5378, "Long Lane", "456", "San Francisco")
	//newAccount = bankaccounts.NewBankAccount("Gerrie de Groot", 5378, "Blaak", "56", "Rotterdam")
	//newAccount = bankaccounts.NewBankAccount("Eva Nierling", 12897, "Coolsingel", "126", "Rotterdam-Zuid")

	if newAccount != nil {
		if !bankaccounts.CheckForExistingAccount(newAccount, cl) {

			cl.AddAccountToLedger(newAccount)

			bankLedger, _ := json.Marshal(cl) // the []byte bankLedger is created to be written to the json file
			writeErr := ioutil.WriteFile("bankledger.json", bankLedger, 0644)
			if err != nil {
				log.Fatal("Could not write to JSON file on disk: ", writeErr)
			}
		}
	}
	// reading from the ledger
	for k, _ := range cl.Ledger {
		fmt.Printf("for k is %x, cl.Ledger[0].Name is: %s\n", k, cl.Ledger[k].Name)

	}
}
