package types

import "time"

type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password,omitempty"` // Omit in JSON if not needed for API responses
	Companies   []string  `json:"companies"`
	Departments []string  `json:"departments"`
	Roles       []string  `json:"roles"`
	Permissions []string  `json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
	LastLogin   time.Time `json:"last_login,omitempty"`
}

type UserMetadata struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Companies   []string `json:"companies"`
    Departments []string `json:"departments"`
    Roles       []string `json:"roles"`
    Permissions []string `json:"permissions"`
}