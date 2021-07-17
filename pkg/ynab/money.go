package ynab

func YNABMoney(amount float64) int64 {
	return int64(amount * 1000)
}

func ConvertFromYNABMoney(amount int64) float64 {
	return float64(amount) / 1000
}
