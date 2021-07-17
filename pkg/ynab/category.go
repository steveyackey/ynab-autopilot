package ynab

type CategoriesResponse struct {
	Data CategoriesResponseData `json:"data"`
}

type CategoriesResponseData struct {
	CategoryGroups []CategoryGroupWithCategories `json:"category_groups"`
	// The knowledge of the server
	ServerKnowledge int64 `json:"server_knowledge"`
}

type CategoryGroupWithCategories struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// Whether or not the category group is hidden
	Hidden bool `json:"hidden"`
	// Whether or not the category group has been deleted.  Deleted category groups will only be included in delta requests.
	Deleted bool `json:"deleted"`
	// Category group categories.  Amounts (budgeted, activity, balance, etc.) are specific to the current budget month (UTC).
	Categories []Category `json:"categories"`
}

type Category struct {
	Id              string `json:"id"`
	CategoryGroupId string `json:"category_group_id"`
	Name            string `json:"name"`
	// Whether or not the category is hidden
	Hidden bool `json:"hidden"`
	// If category is hidden this is the id of the category group it originally belonged to before it was hidden.
	OriginalCategoryGroupId string `json:"original_category_group_id,omitempty"`
	Note                    string `json:"note,omitempty"`
	// Budgeted amount in milliunits format
	Budgeted int64 `json:"budgeted"`
	// Activity amount in milliunits format
	Activity int64 `json:"activity"`
	// Balance in milliunits format
	Balance int64 `json:"balance"`
	// The type of goal, if the category has a goal (TB='Target Category Balance', TBD='Target Category Balance by Date', MF='Monthly Funding', NEED='Plan Your Spending')
	GoalType string `json:"goal_type,omitempty"`
	// The month a goal was created
	GoalCreationMonth string `json:"goal_creation_month,omitempty"`
	// The goal target amount in milliunits
	GoalTarget int64 `json:"goal_target,omitempty"`
	// The original target month for the goal to be completed.  Only some goal types specify this date.
	GoalTargetMonth string `json:"goal_target_month,omitempty"`
	// The percentage completion of the goal
	GoalPercentageComplete int32 `json:"goal_percentage_complete,omitempty"`
	// The number of months, including the current month, left in the current goal period.
	GoalMonthsToBudget int32 `json:"goal_months_to_budget,omitempty"`
	// The amount of funding still needed in the current month to stay on track towards completing the goal within the current goal period.  This amount will generally correspond to the 'Underfunded' amount in the web and mobile clients except when viewing a category with a Needed for Spending Goal in a future month.  The web and mobile clients will ignore any funding from a prior goal period when viewing category with a Needed for Spending Goal in a future month.
	GoalUnderFunded int64 `json:"goal_under_funded,omitempty"`
	// The total amount funded towards the goal within the current goal period.
	GoalOverallFunded int64 `json:"goal_overall_funded,omitempty"`
	// The amount of funding still needed to complete the goal within the current goal period.
	GoalOverallLeft int64 `json:"goal_overall_left,omitempty"`
	// Whether or not the category has been deleted.  Deleted categories will only be included in delta requests.
	Deleted bool `json:"deleted"`
}
