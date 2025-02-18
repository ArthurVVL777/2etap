package dto

import "github.com/google/uuid"

type BankAccountDTO struct {
	ID            uuid.UUID `json:"id"`
	LegalEntityID uuid.UUID `json:"legalEntityId"`
	BIK           string    `json:"bik"`
	BankName      string    `json:"bankName,omitempty"`
	Address       string    `json:"address,omitempty"`
	CorrAccount   string    `json:"corrAccount,omitempty"`
	AccountNumber string    `json:"accountNumber"`
	Currency      string    `json:"currency,omitempty"`
	Comment       string    `json:"comment,omitempty"`
}
