package ynab

import "testing"

func TestYNABMoney(t *testing.T) {
	tests := []struct {
		name   string
		amount float64
		want   int64
	}{
		{name: "One dollar", amount: 1.00, want: 1000},
		{name: "Zero dollars", amount: 0.00, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YNABMoney(tt.amount); got != tt.want {
				t.Errorf("YNABMoney() = %v, want %v", got, tt.want)
			}
		})
	}
}
