package model

import (
	"time"
)

// Exercise represents an Exercise record.
type Exercise struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}