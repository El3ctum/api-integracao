package models

type Document struct {
	Resource  Resource `json:"resource"`
	Data      []byte   `json:"data"`
	Encrypted bool     `json:"encrypted,omitempty"` // For indicating if the document is encrypted
}
