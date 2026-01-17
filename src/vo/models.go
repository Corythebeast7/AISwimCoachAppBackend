package vo

import (
	"time"
)

// BaseModel contains common fields for all models
type BaseModel struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User represents a user in the system
type User struct {
	BaseModel
	Name  string `json:"name"`
	Email string `json:"email"`
}

// SwimData represents swimming data for a user
type SwimData struct {
	BaseModel
	UserID      string  `json:"user_id"`
	Date        string  `json:"date"`
	Distance    float64 `json:"distance"`
	Duration    float64 `json:"duration"`
	StrokeType  string  `json:"stroke_type"`
	Notes       string  `json:"notes,omitempty"`
}