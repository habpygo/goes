package bankaccounts

// BankAccount is the root datastructure
type BankAccount struct {
	Name          string
	AccountNumber string
	Street        string
	HouseNumber   string
	City          string
}

// CustodianLedger takes all the bankaccounts
type CustodianLedger struct {
	Ledger []*BankAccount
}
