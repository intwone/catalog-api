package use_cases

import (
	"time"

	"github.com/intwone/catalog/internal/domain/entities"
	"github.com/intwone/catalog/internal/domain/repositories"
)

type CreateCategoryInput struct {
	Name        string
	Description string
	IsActive    bool
}

type CreateCategoryOutput struct {
	ID          string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

type CreateCategoryUseCaseInterface interface {
	Execute(input CreateCategoryInput) (*CreateCategoryOutput, error)
}

type CreateCategoryUseCase struct {
	CategoryRepository repositories.CategoryRepositoryInterface
}

func NewCreateCategoryUseCase(cr repositories.CategoryRepositoryInterface) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		CategoryRepository: cr,
	}
}

func (uc *CreateCategoryUseCase) Execute(input CreateCategoryInput) (*CreateCategoryOutput, error) {
	category, categoryErr := entities.NewCategory(input.Name, input.Description, input.IsActive)
	if categoryErr != nil {
		return nil, categoryErr
	}
	categoryRepositoryErr := uc.CategoryRepository.Save(*category)
	if categoryRepositoryErr != nil {
		return nil, categoryRepositoryErr
	}
	return &CreateCategoryOutput{
		ID:          category.GetID().String(),
		Name:        category.GetName(),
		Description: category.GetDescription(),
		IsActive:    category.GetIsActive(),
		CreatedAt:   category.GetCreatedAt(),
	}, nil
}
