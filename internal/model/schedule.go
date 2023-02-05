package model

import (
	"time"
)

// Schedule represents an Schedule record.
type Schedule struct {
	Id			uint8				`json:"id"`
	UserId		uint8				`json:"userId"`
	Title     string                 `json:"title"`
	Description  string                 `json:"description"`
	// Monday       map[string]interface{} `json:"monday"`
	// Tuesday      map[string]interface{} `json:"tuesday"`
	// Wednesday      map[string]interface{} `json:"wednesday"`
	// Thursday      map[string]interface{} `json:"thursday"`
	// Friday      map[string]interface{} `json:"friday"`
	// Saturday      map[string]interface{} `json:"saturday"`
	// Sunday      map[string]interface{} `json:"sunday"`

	TimeCreated time.Time `json:"timeCreated"`
	TimeUpdated time.Time `json:"timeUpdated"`
	Color	string		`json:"color"`
}