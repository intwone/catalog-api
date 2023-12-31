package use_cases

import (
	"time"

	"github.com/intwone/catalog/internal/application/errs"
	"github.com/intwone/catalog/internal/domain/repositories"
)

type DeleteCategoryByIDInput struct {
	ID string
}

type DeleteCategoryByIDOutput struct {
	ID          string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

type DeleteCategoryByIDUseCaseInterface interface {
	Execute(input DeleteCategoryByIDInput) (*DeleteCategoryByIDOutput, error)
}

type DeleteCategoryByIDUseCase struct {
	CategoryRepository repositories.CategoryRepositoryInterface
}

func NewDeleteCategoryByIDUseCase(cr repositories.CategoryRepositoryInterface) *DeleteCategoryByIDUseCase {
	return &DeleteCategoryByIDUseCase{
		CategoryRepository: cr,
	}
}

func (uc *DeleteCategoryByIDUseCase) Execute(input DeleteCategoryByIDInput) (*DeleteCategoryByIDOutput, error) {
	category, getCategoryRepositoryErr := uc.CategoryRepository.GetByID(input.ID)
	if err := errs.HandleRepositoryError(getCategoryRepositoryErr); err != nil {
		return nil, err
	}
	deleteCategoryRepositoryErr := uc.CategoryRepository.DeleteByID(input.ID)
	if err := errs.HandleRepositoryError(deleteCategoryRepositoryErr); err != nil {
		return nil, err
	}
	return &DeleteCategoryByIDOutput{
		ID:          category.GetID().String(),
		Name:        category.GetName(),
		Description: category.GetDescription(),
		IsActive:    category.GetIsActive(),
		CreatedAt:   category.GetCreatedAt(),
	}, nil
}
