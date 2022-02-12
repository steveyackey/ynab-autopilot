package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Settings struct {
	BearerToken     string
	Budget          string
	CategoryActions []CategoryAction
}

type CategoryAction struct {
	CategoryID     string
	CategoryName   string
	AmountToBudget float64
}

func NewConfig() *Settings {
	viper.SetConfigName("config")                // name of config file (without extension)
	viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/ynab-autopilot/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.ynab-autopilot") // call multiple times to add many search paths
	viper.AddConfigPath(".")                     // optionally look for config in the working directory
	err := viper.ReadInConfig()                  // Find and read the config file
	if err != nil {                              // Handle errors reading the config file
		panic(fmt.Errorf("fatal error in config file: %w", err))
	}

	if !viper.IsSet("bearerToken") {
		panic("Please create a config file and add:\n bearerToken: YourTokenHere")
	}

	var categoryActions []CategoryAction
	err = viper.UnmarshalKey("categoryActions", &categoryActions)
	if err != nil {
		panic(fmt.Errorf("category actions cannot be unmarshalled, %v", err))
	}

	settings := Settings{
		BearerToken:     viper.GetString("bearerToken"),
		Budget:          viper.GetString("budgetID"),
		CategoryActions: categoryActions,
	}
	return &settings
}
