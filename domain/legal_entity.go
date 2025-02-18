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

type BankAccount struct {
	ID            uuid.UUID `json:"id"`
	LegalEntityID uuid.UUID `json:"legal_entity_id"`
	BIK           string    `json:"bik"` // цифра
	BankName      string    `json:"bank_name"`
	Address       string    `json:"address"`

	AccountNumber string    `json:"account_number"`
	Currency      string    `json:"currency"`
	Comment       string    `json:"comment"`
	IsPrimary     bool      `json:"is_primary"`
}
