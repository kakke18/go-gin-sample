package models

import "time"

type Member struct {
	Key          string    `json:"key"`
	Name         string    `json:"name"`
	EmailAddress string    `json:"email_address"`
	CreatedAt    time.Time `json:"created_at"`
}
