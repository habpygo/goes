package bankaccounts

// Bankaccount is the root datastructure
type Bankaccount struct {
	Name          string
	AccountNumber string
	Address
}

// Address is put into a separate struct; for KYC procedures later on
type Address struct {
	Street      string
	HouseNumber string
	City        string
}

// CustodianLedger takes all the bankaccounts
type CustodianLedger struct {
	Ledger []*Bankaccount
}
