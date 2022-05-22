package models

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID        string      `json:"id" db:"id"`
	Username  string      `json:"username" db:"username"`
	Firstname string      `json:"firstname" db:"firstname"`
	Lastname  string      `json:"lastname" db:"lastname"`
	Email     string      `json:"email" db:"email"`
	Password  string      `json:"password,omitempty" db:"password"`
	Role      string      `json:"role" db:"role"`
	Birthday  pq.NullTime `json:"birthday" db:"birthday"`
	CreatedAt time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time   `json:"updatedAt" db:"updated_at"`
}
