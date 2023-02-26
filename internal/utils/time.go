package utils

import (
	"time"
)

type RawTime []byte

func (t RawTime) Time() (time.Time, error) {
    return time.Parse("2023-02-11 12:15", string(t))
}