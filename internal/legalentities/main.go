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

// BankAccount.
func (s *Service) GetAllBankAccounts(ctx context.Context, legalEntityID uuid.UUID) ([]domain.BankAccount, error) {
	return s.repo.GetAllBankAccounts(ctx, legalEntityID)
}

func (s *Service) CreateBankAccount(ctx context.Context, account domain.BankAccount) (domain.BankAccount, error) {
	err := s.repo.CreateBankAccount(ctx, &account)
	return account, err
}

func (s *Service) UpdateBankAccount(ctx context.Context, account domain.BankAccount) error {
	return s.repo.UpdateBankAccount(ctx, &account)
}

func (s *Service) DeleteBankAccount(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteBankAccount(ctx, id)
}
