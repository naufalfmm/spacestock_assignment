package models

import (
	"time"
)

// BaseModel model
type BaseModel struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
