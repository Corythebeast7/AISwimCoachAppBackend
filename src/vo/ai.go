package vo

// AiQueryRequest represents the input payload for the AI query endpoint
type AiQueryRequest struct {
	Prompt string `json:"prompt"`
}

// AiQueryResult represents the successful result of an AI query
type AiQueryResult struct {
	Answer string `json:"answer"`
}
