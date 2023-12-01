package use_cases

import (
	"time"

	"github.com/intwone/catalog/internal/application/errs"
	"github.com/intwone/catalog/internal/domain/repositories"
)

type DeactiveCategoryByIDInput struct {
	ID string
}

type DeactiveCategoryByIDOutput struct {
	ID          string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

type DeactiveCategoryByIDUseCaseInterface interface {
	Execute(input DeactiveCategoryByIDInput) (*DeactiveCategoryByIDOutput, error)
}

type DeactiveCategoryByIDUseCase struct {
	CategoryRepository repositories.CategoryRepositoryInterface
}

func NewDeactiveCategoryByIDUseCase(cr repositories.CategoryRepositoryInterface) *DeactiveCategoryByIDUseCase {
	return &DeactiveCategoryByIDUseCase{
		CategoryRepository: cr,
	}
}

func (uc *DeactiveCategoryByIDUseCase) Execute(input DeactiveCategoryByIDInput) (*DeactiveCategoryByIDOutput, error) {
	category, getCategoryRepositoryErr := uc.CategoryRepository.GetByID(input.ID)
	if err := errs.HandleRepositoryError(getCategoryRepositoryErr); err != nil {
		return nil, err
	}
	category.SetIsActive(false)
	updateCategoryRepositoryErr := uc.CategoryRepository.Update(*category)
	if err := errs.HandleRepositoryError(updateCategoryRepositoryErr); err != nil {
		return nil, err
	}
	return &DeactiveCategoryByIDOutput{
		ID:          category.GetID().String(),
		Name:        category.GetName(),
		Description: category.GetDescription(),
		IsActive:    category.GetIsActive(),
		CreatedAt:   category.GetCreatedAt(),
	}, nil
}
