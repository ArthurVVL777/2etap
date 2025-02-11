package legalentities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/domain"

	"github.com/krisch/crm-backend/internal/catalogs"
	"github.com/krisch/crm-backend/internal/dictionary"
)

type Service struct {
	repo     RepositoryInterface
	dict     *dictionary.Service
	catalogs *catalogs.Service
}

func NewService(repo RepositoryInterface, dict *dictionary.Service, cs *catalogs.Service) *Service {
	return &Service{
		repo:     repo,
		dict:     dict,
		catalogs: cs,
	}
}

func (s *Service) GetAllLegalEntities(ctx context.Context) (entities []domain.LegalEntity, err error) {
	entities, err = s.repo.GetAllLegalEntities(ctx)
	return entities, err
}

func (s *Service) CreateLegalEntity(ctx context.Context, name string) (entity domain.LegalEntity, err error) {
	entity = domain.LegalEntity{
		UUID:      uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.CreateLegalEntity(ctx, &entity)
	return entity, err
}

func (s *Service) UpdateLegalEntity(ctx context.Context, uid uuid.UUID, name string) (err error) {
	err = s.repo.UpdateLegalEntity(ctx, uid, name)
	return err
}

func (s *Service) DeleteLegalEntity(ctx context.Context, uid uuid.UUID) (err error) {
	err = s.repo.DeleteLegalEntity(ctx, uid)
	return err
}
