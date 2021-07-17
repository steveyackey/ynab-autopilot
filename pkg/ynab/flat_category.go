package ynab

import "fmt"

type FlatCategory struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Budgeted  int64  `json:"budgeted"`
	Balance   int64  `json:"balance"`
	GroupID   string `json:"category_group_id"`
	GroupName string `json:"category_group_name"`
}

type FlatCategories struct {
	Data map[string]FlatCategory
}

func (c *CategoriesResponse) MapToFlatCategories() *FlatCategories {
	var categories = FlatCategories{Data: make(map[string]FlatCategory)}
	for _, group := range c.Data.CategoryGroups {
		for _, category := range group.Categories {
			categories.Data[category.Name] = FlatCategory{
				Id:        category.Id,
				Name:      category.Name,
				Budgeted:  category.Budgeted,
				Balance:   category.Balance,
				GroupID:   group.Id,
				GroupName: group.Name,
			}
		}
	}
	return &categories
}

func (f *FlatCategories) ToYAML() string {
	result := "categories: \n"
	yamlTemplate := `
    - categoryId: %s
      categoryName: %s
      budgeted: %d
      balance: %d
      category_group_id: %s
      category_group_name: %s
	`
	for _, cat := range f.Data {
		result += fmt.Sprintf(yamlTemplate, cat.Id, cat.Name, cat.Balance, cat.Budgeted, cat.GroupID, cat.GroupName)
	}

	return result

}
