package repositories

import (
	"github.com/intwone/catalog/internal/domain/entities"
)

type DataToUpdate struct {
	Name        *string
	Description *string
}

type CategoryRepositoryInterface interface {
	Save(category entities.Category) error
	Get() (*[]entities.Category, error)
	GetByID(id string) (*entities.Category, error)
	DeleteByID(id string) error
	Update(category entities.Category) error
}
