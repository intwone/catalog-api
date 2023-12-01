package use_cases

import (
	"time"

	"github.com/intwone/catalog/internal/application/errs"
	"github.com/intwone/catalog/internal/domain/repositories"
)

type UpdateCategoryByIDInput struct {
	ID          string
	Name        string
	Description string
}

type UpdateCategoryByIDOutput struct {
	ID          string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

type UpdateCategoryByIDUseCaseInterface interface {
	Execute(input UpdateCategoryByIDInput) (*UpdateCategoryByIDOutput, error)
}

type UpdateCategoryByIDUseCase struct {
	CategoryRepository repositories.CategoryRepositoryInterface
}

func NewUpdateCategoryByIDUseCase(cr repositories.CategoryRepositoryInterface) *UpdateCategoryByIDUseCase {
	return &UpdateCategoryByIDUseCase{
		CategoryRepository: cr,
	}
}

func (uc *UpdateCategoryByIDUseCase) Execute(input UpdateCategoryByIDInput) (*UpdateCategoryByIDOutput, error) {
	category, getCategoryRepositoryErr := uc.CategoryRepository.GetByID(input.ID)
	if err := errs.HandleRepositoryError(getCategoryRepositoryErr); err != nil {
		return nil, err
	}
	if err := category.SetName(input.Name); err != nil {
		return nil, err
	}
	if err := category.SetDescription(input.Description); err != nil {
		return nil, err
	}
	updateCategoryRepositoryErr := uc.CategoryRepository.Update(*category)
	if err := errs.HandleRepositoryError(updateCategoryRepositoryErr); err != nil {
		return nil, err
	}
	return &UpdateCategoryByIDOutput{
		ID:          category.GetID().String(),
		Name:        category.GetName(),
		Description: category.GetDescription(),
		IsActive:    category.GetIsActive(),
		CreatedAt:   category.GetCreatedAt(),
	}, nil
}
