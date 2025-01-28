package legalentities

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllLegalEntities(ctx context.Context) ([]LegalEntity, error)
	CreateLegalEntity(ctx context.Context, entity LegalEntity) error
	UpdateLegalEntity(ctx context.Context, id string, entity LegalEntity) error
	DeleteLegalEntity(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllLegalEntities(ctx context.Context) ([]LegalEntity, error) {
	var entities []LegalEntity
	err := r.db.WithContext(ctx).Find(&entities).Error
	return entities, err
}

func (r *repository) CreateLegalEntity(ctx context.Context, entity LegalEntity) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *repository) UpdateLegalEntity(ctx context.Context, id string, entity LegalEntity) error {
	return r.db.WithContext(ctx).Model(&LegalEntity{}).Where("id = ?", id).Updates(entity).Error
}

func (r *repository) DeleteLegalEntity(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&LegalEntity{}, "id = ?", id).Error
}
