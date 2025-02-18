package legalentities

import (
	"time"

	"github.com/google/uuid"
)

// LegalEntity представляет структуру юридического лица.
type LegalEntity struct {
	UUID      uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();not null;unique"`
	Name      string     `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:now();not null"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;default:now();not null"`
	DeletedAt *time.Time `gorm:"type:timestamptz;default:NULL;"`
}

type BankAccount struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	LegalEntityID uuid.UUID `gorm:"type:uuid;not null;index"`
	BIK           string    `gorm:"type:varchar(9);not null"`
	BankName      string    `gorm:"type:varchar(255)"`
	Address       string    `gorm:"type:varchar(255)"`
	CorrAccount   string    `gorm:"type:varchar(20)"`
	AccountNumber string    `gorm:"type:varchar(20);not null"`
	Currency      string    `gorm:"type:varchar(10)"`
	Comment       string    `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (BankAccount) TableName() string {
	return "bank_accounts"
}
