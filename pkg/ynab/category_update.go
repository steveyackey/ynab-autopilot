package ynab

type CategoryUpdate struct {
	Data CategoryUpdateData `json:"data"`
}

type CategoryUpdateData struct {
	Category Category `json:"category"`
	// The knowledge of the server
	ServerKnowledge int64 `json:"server_knowledge"`
}
