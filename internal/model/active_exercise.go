package model

import (
	"time"
)

// Exercise represents an Exercise record.
type Active_Exercise struct {
	Schedule_ID      int    `json:"schedule_id"`
	Generic_ID      int    `json:"generic_id"`
	Reps 			string 	`json:"reps"`
	Sets			int 	`json:"sets"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}