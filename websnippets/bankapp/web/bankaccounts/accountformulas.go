package bankaccounts

// NewBankAccount takes in the new data and returns a new bank account
func NewBankAccount(name string, accountnumber string, street string, number string, city string) *BankAccount {
	ba := &BankAccount{
		Name:          name,
		AccountNumber: accountnumber,
		Street:        street,
		HouseNumber:   number,
		City:          city,
	}
	return ba
}

// AddAccountToLedger adds a new account to the ledger
func (cl *CustodianLedger) AddAccountToLedger(account *BankAccount) {
	newAccount := account
	cl.Ledger = append(cl.Ledger, newAccount)
}

// CheckForExistingAccount checks if the new account already exists
func CheckForExistingAccount(account *BankAccount, cl CustodianLedger) bool {

	for _, value := range cl.Ledger {
		if value.AccountNumber == account.AccountNumber {
			return true
		}
	}

	return false
}
