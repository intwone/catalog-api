package repositories

import "github.com/intwone/catalog/internal/domain/entities"

type CategoryRepositoryInterface interface {
	Save(entities.Category) error
	GetByID(id string) (*entities.Category, error)
}
