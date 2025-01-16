package types

import "time"

type Resource struct {
	ID         uint                   `json:"id"`
	Type       string                 `json:"type"` // Ex: "document", "photo"
	OwnerID    uint                   `json:"owner_id"`
	ParentID   *uint                  `json:"parent_id,omitempty"` // For hierarchy, nullable
	Roles      []string               `json:"roles"`
	Attributes map[string]interface{} `json:"attributes"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
}

type Document struct {
	Resource  Resource `json:"resource"`
	Data      []byte   `json:"data"`
	Encrypted bool     `json:"encrypted,omitempty"` // For indicating if the document is encrypted
}
