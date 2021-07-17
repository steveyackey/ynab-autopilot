package ynab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Client struct {
	BearerToken string
}

func NewClient(bearerToken string) (*Client, error) {
	return &Client{BearerToken: fmt.Sprintf("Bearer %s", bearerToken)}, nil
}

var YNABApiV1 = "https://api.youneedabudget.com/v1"

func (c *Client) NewGetRequest(method string, path string) ([]byte, error) {

	fullUrl := fmt.Sprintf("%s%s", YNABApiV1, path)
	req, err := http.NewRequest(method, fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error while creating request: %v", err)
	}
	req.Header.Add("Authorization", c.BearerToken)
	req.Header.Add("Accept", "application/json")
	client := &http.Client{Timeout: 2 * time.Minute}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error on response.\n[ERROR] - %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the response bytes: %v", err)
	}
	return body, nil
}

func (c *Client) NewPatchRequest(method string, path string, payload []byte) ([]byte, error) {

	fullUrl := fmt.Sprintf("%s%s", YNABApiV1, path)
	req, err := http.NewRequest(method, fullUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error while creating request: %v", err)
	}
	req.Header.Add("Authorization", c.BearerToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 2 * time.Minute}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error on response.\n[ERROR] - %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the response bytes: %v", err)
	}
	return body, nil
}

func (c *Client) ListBudgets() {
	result, err := c.NewGetRequest(http.MethodGet, "/budgets")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	var budgets BudgetResponse
	json.Unmarshal(result, &budgets)
	for _, budget := range budgets.Data.Budgets {
		fmt.Printf("ID: %s\tName:%s\n", budget.Id, budget.Name)
	}
}

func (c *Client) GetCategories(budgetID string) *FlatCategories {
	categoriesEndpoint := fmt.Sprintf("/budgets/%s/categories", budgetID)
	result, err := c.NewGetRequest(http.MethodGet, categoriesEndpoint)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	var categories CategoriesResponse
	err = json.Unmarshal(result, &categories)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	return categories.MapToFlatCategories()
}

type UpdateCategoryPayload struct {
	Category UpdateCategoryPayloadCategory `json:"category"`
}

type UpdateCategoryPayloadCategory struct {
	Budgeted int64 `json:"budgeted"`
}

func (c *Client) UpdateCategory(budgetID string, month string, categoryID string, amount float64) (CategoryUpdate, error) {
	updateCategoryEndpoint := fmt.Sprintf("/budgets/%s/months/%s/categories/%s", budgetID, month, categoryID)
	var payload, err = json.Marshal(UpdateCategoryPayload{UpdateCategoryPayloadCategory{Budgeted: YNABMoney(amount)}})
	if err != nil {
		return CategoryUpdate{}, err
	}

	result, err := c.NewPatchRequest(http.MethodPatch, updateCategoryEndpoint, payload)
	if err != nil {
		return CategoryUpdate{}, err
	}
	var categoryResult CategoryUpdate
	err = json.Unmarshal(result, &categoryResult)
	if err != nil {
		return CategoryUpdate{}, err
	}
	return categoryResult, nil
}
