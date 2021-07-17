package ynab

import (
	"time"
)

type BudgetResponse struct {
	Data BudgetResponseData `json:"data"`
}

type BudgetResponseData struct {
	Budgets []Budget `json:"budgets"`
}

type Budget struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// The last time any changes were made to the budget from either a web or mobile client
	LastModifiedOn time.Time `json:"last_modified_on,omitempty"`
	// The earliest budget month
	FirstMonth string `json:"first_month,omitempty"`
	// The latest budget month
	LastMonth  string `json:"last_month,omitempty"`
	DateFormat string `json:"date_format,omitempty"`
	// The budget accounts (only included if `include_accounts=true` specified as query parameter)
	Accounts *[]Account `json:"accounts,omitempty"`
}
