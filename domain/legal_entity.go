package domain

import (
	"time"

	"github.com/google/uuid"
)

type LegalEntity struct {
	UUID      uuid.UUID  `json:"uuid"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
