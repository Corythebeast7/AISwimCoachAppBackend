package service

import (
	vo2 "AISwimCoachAppBackend/src/vo"
	"encoding/json"
	"net/http"
	"strings"
)

// AiQueryService performs AI-related query operations
// For now, this is a simple stub that echoes the prompt back as an answer.
type AiQueryService struct{}

// NewAiQueryService constructs a new AiQueryService instance
func NewAiQueryService() *AiQueryService {
	return &AiQueryService{}
}

// Query processes the given prompt and returns an answer
func (s *AiQueryService) Query(prompt string) (vo2.AiQueryResult, error) {
	answer := "You asked: " + strings.TrimSpace(prompt)
	return vo2.AiQueryResult{Answer: answer}, nil
}

var aiService = NewAiQueryService()

// AiQueryHandler handles POST /api/ai/query requests
// It expects a JSON body: { "prompt": "..." }
func AiQueryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req vo2.AiQueryRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(vo2.ErrorResponse{Status: "error", Message: "invalid JSON body"})
		return
	}

	if strings.TrimSpace(req.Prompt) == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(vo2.ErrorResponse{Status: "error", Message: "prompt is required"})
		return
	}

	result, err := aiService.Query(req.Prompt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(vo2.ErrorResponse{Status: "error", Message: "failed to process ai query"})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(vo2.Response{Status: "ok", Data: result})
}
