package models

import "time"

type Record struct {
	Key        string
	CreatedAt  time.Time
	TotalCount []int
}
