package repositories

import (
	"github.com/intwone/catalog/internal/domain/entities"
)

type CategoryRepositoryInterface interface {
	Save(category entities.Category) error
	Get() (*[]entities.Category, error)
	GetByID(id string) (*entities.Category, error)
	DeleteByID(id string) error
	Update(category entities.Category) error
	Search(offset int64, limit int64, keyword string) (*[]entities.Category, error)
}
