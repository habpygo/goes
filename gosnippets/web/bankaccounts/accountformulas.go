package bankaccounts

// NewBankAccount takes in the new data and returns a new bank account
func NewBankAccount(name string, accountnumber string, street string, number string, city string) *Bankaccount {
	ba := &Bankaccount{
		Name:          name,
		AccountNumber: accountnumber,
		Street:        street,
		HouseNumber:   number,
		City:          city,
	}
	return ba
}

// AddAccountToLedger adds a new account to the ledger
func (cl *CustodianLedger) AddAccountToLedger(account *Bankaccount) {
	newAccount := account
	cl.Ledger = append(cl.Ledger, newAccount)
}

// CheckForExistingAccount checks if the new account already exists
func CheckForExistingAccount(account *Bankaccount, cl CustodianLedger) bool {

	for _, value := range cl.Ledger {
		if value.AccountNumber == account.AccountNumber {
			return true
		}
	}

	return false
}
