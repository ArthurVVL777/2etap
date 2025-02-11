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
