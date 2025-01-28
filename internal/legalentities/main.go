package legalentities

import (
	"context"
)

type Service struct {
	repo Repository
}

// NewService creates a new instance of Service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetAllLegalEntities retrieves all legal entities.
func (s *Service) GetAllLegalEntities(ctx context.Context) ([]LegalEntity, error) {
	return s.repo.GetAllLegalEntities(ctx)
}

// CreateLegalEntity creates a new legal entity.
func (s *Service) CreateLegalEntity(ctx context.Context, entity LegalEntity) error {
	return s.repo.CreateLegalEntity(ctx, entity)
}

// UpdateLegalEntity updates an existing legal entity.
func (s *Service) UpdateLegalEntity(ctx context.Context, id string, entity LegalEntity) error {
	return s.repo.UpdateLegalEntity(ctx, id, entity)
}

// DeleteLegalEntity deletes a legal entity by ID.
func (s *Service) DeleteLegalEntity(ctx context.Context, id string) error {
	return s.repo.DeleteLegalEntity(ctx, id)
}
