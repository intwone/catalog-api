package use_cases

import (
	"time"

	"github.com/intwone/catalog/internal/application/errs"
	"github.com/intwone/catalog/internal/domain/repositories"
)

type ActiveCategoryByIDInput struct {
	ID string
}

type ActiveCategoryByIDOutput struct {
	ID          string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

type ActiveCategoryByIDUseCaseInterface interface {
	Execute(input ActiveCategoryByIDInput) (*ActiveCategoryByIDOutput, error)
}

type ActiveCategoryByIDUseCase struct {
	CategoryRepository repositories.CategoryRepositoryInterface
}

func NewActiveCategoryByIDUseCase(cr repositories.CategoryRepositoryInterface) *ActiveCategoryByIDUseCase {
	return &ActiveCategoryByIDUseCase{
		CategoryRepository: cr,
	}
}

func (uc *ActiveCategoryByIDUseCase) Execute(input ActiveCategoryByIDInput) (*ActiveCategoryByIDOutput, error) {
	category, getCategoryRepositoryErr := uc.CategoryRepository.GetByID(input.ID)
	if err := errs.HandleRepositoryError(getCategoryRepositoryErr); err != nil {
		return nil, err
	}
	category.SetIsActive(true)
	updateCategoryRepositoryErr := uc.CategoryRepository.Update(*category)
	if err := errs.HandleRepositoryError(updateCategoryRepositoryErr); err != nil {
		return nil, err
	}
	return &ActiveCategoryByIDOutput{
		ID:          category.GetID().String(),
		Name:        category.GetName(),
		Description: category.GetDescription(),
		IsActive:    category.GetIsActive(),
		CreatedAt:   category.GetCreatedAt(),
	}, nil

}
