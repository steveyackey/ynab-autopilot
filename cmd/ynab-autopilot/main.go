package main

import (
	"fmt"
	"os"

	"github.com/steveyackey/ynab-autopilot/internal/config"
	"github.com/steveyackey/ynab-autopilot/pkg/ynab"
)

func UpdateAllCategories(settings *config.Settings, client *ynab.Client, categories *ynab.FlatCategories) {
	for _, action := range settings.CategoryActions {
		initialBudgeted := ynab.ConvertFromYNABMoney(categories.Data[action.CategoryName].Budgeted)
		result, err := client.UpdateCategory(settings.Budget, "current", action.CategoryID, initialBudgeted+action.AmountToBudget)
		if err != nil {
			fmt.Printf("Error updating %s, %v\n", action.CategoryName, err)
		}
		fmt.Printf("UPDATE RESULTS:\n")
		fmt.Printf("%q:\n  Initial: $%g\n  Added:   $%g\n  Current: $%g\n", action.CategoryName, initialBudgeted, action.AmountToBudget, ynab.ConvertFromYNABMoney(result.Data.Category.Balance))
	}
}
func main() {
	settings := config.NewConfig()

	client, err := ynab.NewClient(settings.BearerToken)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	if settings.Budget == "" {
		fmt.Println("Please add budgetID: your-budget-id to your config. Here's a current list of budgets:")
		client.ListBudgets()
		os.Exit(1)

	}

	categories := client.GetCategories(settings.Budget)
	if len(os.Args) < 2 {
		fmt.Println("Please add a command: list, run")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		fmt.Println(categories.ToYAML())
		os.Exit(0)

	case "run":
		UpdateAllCategories(settings, client, categories)
	default:
		fmt.Println("Please add a command: list, run")
	}

}
