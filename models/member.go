package models

import "time"

type Member struct {
	Key          string
	Name         string
	EmailAddress string
	CreatedAt    time.Time
}
