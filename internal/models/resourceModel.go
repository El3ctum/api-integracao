package models

import "time"

type Resource struct {
	ID         uint                   `json:"id"`
	Type       string                 `json:"type"` // Ex: "document", "photo"
	OwnerID    uint                   `json:"owner_id"`
	ParentID   uint                   `json:"parent_id,omitempty"` // For hierarchy, nullable
	Roles      []string               `json:"roles"`
	Attributes map[string]interface{} `json:"attributes"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
}
