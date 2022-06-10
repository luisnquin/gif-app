package models

import "time"

type Profile struct {
	ID             string    `json:"id" db:"id"`
	LastConnection time.Time `json:"lastConnection" db:"last_connection"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}
