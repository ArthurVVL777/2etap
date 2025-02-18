package legalentities

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/domain"
	"github.com/krisch/crm-backend/pkg/postgres"
	"github.com/krisch/crm-backend/pkg/redis"
	"github.com/sirupsen/logrus"
)

type RepositoryInterface interface {
	GetAllLegalEntities(ctx context.Context) ([]domain.LegalEntity, error)
	CreateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error
	UpdateLegalEntity(ctx context.Context, uid uuid.UUID, name string) error
	DeleteLegalEntity(ctx context.Context, uid uuid.UUID) error

	GetAllBankAccounts(ctx context.Context, legalEntityID uuid.UUID) ([]domain.BankAccount, error)
	CreateBankAccount(ctx context.Context, account *domain.BankAccount) error
	UpdateBankAccount(ctx context.Context, account *domain.BankAccount) error
	DeleteBankAccount(ctx context.Context, id uuid.UUID) error
}

type Repository struct {
	gorm        *postgres.GDB
	rds         *redis.RDS
	middlewares []func(ctx context.Context, name string) error
}

func NewRepository(db *postgres.GDB, rds *redis.RDS) *Repository {
	return &Repository{
		gorm: db,
		rds:  rds,
	}
}

func (r *Repository) Use(fn func(ctx context.Context, name string) error) {
	r.middlewares = append(r.middlewares, fn)
}

func (r *Repository) apply(ctx context.Context, name string) func() {
	c := ctx
	return func() {
		for _, fn := range r.middlewares {
			if err := fn(c, name); err != nil {
				logrus.Error(err)
			}
		}
	}
}

func (r *Repository) PubUpdate() {
	if err := r.rds.Publish(context.Background(), "update", "legal_entities"); err != nil {
		logrus.Error(err)
	}
}

func (r *Repository) CreateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error {
	defer r.apply(ctx, "CreateLegalEntity")() // Используем переданный контекст

	if entity == nil {
		return errors.New("передано пустое юридическое лицо")
	}

	var existing domain.LegalEntity
	if res := r.gorm.DB.WithContext(ctx).Model(&existing). // Добавляем ctx для запроса
								Where("name = ?", entity.Name).
								Where("deleted_at IS NULL").
								First(&existing); res.RowsAffected > 0 {
		return errors.New("юридическое лицо с таким именем уже существует")
	}

	if err := r.gorm.DB.WithContext(ctx).Create(entity).Error; err != nil { // Добавляем ctx
		return err
	}

	r.PubUpdate()
	return nil
}

func (r *Repository) UpdateLegalEntity(ctx context.Context, uid uuid.UUID, name string) error {
	defer r.apply(ctx, "UpdateLegalEntity")() // Передаем контекст

	res := r.gorm.DB.WithContext(ctx).
		Model(&domain.LegalEntity{}).
		Where("uuid = ?", uid).
		Update("name", name).
		Update("updated_at", "now()")

	if res.RowsAffected == 0 {
		return errors.New("юридическое лицо не найдено")
	}

	r.PubUpdate()
	return res.Error
}

func (r *Repository) DeleteLegalEntity(ctx context.Context, uid uuid.UUID) error {
	defer r.apply(ctx, "DeleteLegalEntity")() // Передаем контекст

	res := r.gorm.DB.WithContext(ctx).
		Model(&domain.LegalEntity{}).
		Where("uuid = ?", uid).
		Where("deleted_at IS NULL").
		Update("deleted_at", "now()")

	if res.RowsAffected == 0 {
		return errors.New("юридическое лицо не найдено")
	}

	r.PubUpdate()
	return res.Error
}

func (r *Repository) GetAllLegalEntities(ctx context.Context) ([]domain.LegalEntity, error) {
	var entities []domain.LegalEntity
	if err := r.gorm.DB.WithContext(ctx).Where("deleted_at IS NULL").Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// BankAccount

func (r *Repository) GetAllBankAccounts(ctx context.Context, legalEntityID uuid.UUID) ([]domain.BankAccount, error) {
	var accounts []domain.BankAccount
	if err := r.gorm.DB.WithContext(ctx).
		Where("legal_entity_id = ?", legalEntityID).
		Where("deleted_at IS NULL").
		Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *Repository) CreateBankAccount(ctx context.Context, account *domain.BankAccount) error {
	defer r.apply(ctx, "CreateBankAccount")()

	if account == nil {
		return errors.New("передан пустой банковский счет")
	}

	if err := r.gorm.DB.WithContext(ctx).Create(account).Error; err != nil {
		return err
	}

	r.PubUpdate()
	return nil
}

func (r *Repository) UpdateBankAccount(ctx context.Context, account *domain.BankAccount) error {
	defer r.apply(ctx, "UpdateBankAccount")()

	res := r.gorm.DB.WithContext(ctx).
		Model(&domain.BankAccount{}).
		Where("id = ?", account.ID).
		Updates(map[string]interface{}{
			"bik":            account.BIK,
			"bank_name":      account.BankName,
			"address":        account.Address,
			"corr_account":   account.CorrAccount,
			"account_number": account.AccountNumber,
			"currency":       account.Currency,
			"comment":        account.Comment,
			"updated_at":     "now()",
		})

	if res.RowsAffected == 0 {
		return errors.New("банковский счет не найден")
	}

	r.PubUpdate()
	return res.Error
}

func (r *Repository) DeleteBankAccount(ctx context.Context, id uuid.UUID) error {
	defer r.apply(ctx, "DeleteBankAccount")()

	res := r.gorm.DB.WithContext(ctx).
		Model(&domain.BankAccount{}).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Update("deleted_at", "now()")

	if res.RowsAffected == 0 {
		return errors.New("банковский счет не найден")
	}

	r.PubUpdate()
	return res.Error
}
