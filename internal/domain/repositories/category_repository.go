package repositories

import (
	"github.com/google/uuid"
	"github.com/intwone/catalog/internal/domain/entities"
)

type CategoryRepositoryInterface interface {
	Save(entities.Category) error
	GetByID(id uuid.UUID) (*entities.Category, error)
	DeleteByID(id uuid.UUID) error
}
