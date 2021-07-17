package ynab

type Account struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// The type of account. Note: payPal, merchantAccount, investmentAccount, and mortgage types have been deprecated and will be removed in the future.
	Type string `json:"type"`
	// Whether this account is on budget or not
	OnBudget bool `json:"on_budget"`
	// Whether this account is closed or not
	Closed bool   `json:"closed"`
	Note   string `json:"note,omitempty"`
	// The current balance of the account in milliunits format
	Balance int64 `json:"balance"`
	// The current cleared balance of the account in milliunits format
	ClearedBalance int64 `json:"cleared_balance"`
	// The current uncleared balance of the account in milliunits format
	UnclearedBalance int64 `json:"uncleared_balance"`
	// The payee id which should be used when transferring to this account
	TransferPayeeId string `json:"transfer_payee_id"`
	// Whether or not the account is linked to a financial institution for automatic transaction import.
	DirectImportLinked bool `json:"direct_import_linked,omitempty"`
	// If an account linked to a financial institution (direct_import_linked=true) and the linked connection is not in a healthy state, this will be true.
	DirectImportInError bool `json:"direct_import_in_error,omitempty"`
	// Whether or not the account has been deleted.  Deleted accounts will only be included in delta requests.
	Deleted bool `json:"deleted"`
}
