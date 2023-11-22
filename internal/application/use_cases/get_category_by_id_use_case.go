package use_cases

import (
	"time"

	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/intwone/catalog/internal/domain/repositories"
)

type GetCategoryByIDInput struct {
	ID string
}

type GetCategoryByIDOutput struct {
	ID          string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

type GetCategoryByIDUseCaseInterface interface {
	Execute(input GetCategoryByIDInput) (*GetCategoryByIDOutput, error)
}

type GetCategoryByIDUseCase struct {
	CategoryRepository repositories.CategoryRepositoryInterface
}

func NewGetCategoryByIDUseCase(cr repositories.CategoryRepositoryInterface) *GetCategoryByIDUseCase {
	return &GetCategoryByIDUseCase{
		CategoryRepository: cr,
	}
}

func (uc *GetCategoryByIDUseCase) Execute(input GetCategoryByIDInput) (*GetCategoryByIDOutput, error) {
	category, categoryRepositoryErr := uc.CategoryRepository.GetByID(input.ID)
	if categoryRepositoryErr.Error() == errs.ResourceNotFound.Error() {
		return nil, categoryRepositoryErr
	}
	if categoryRepositoryErr != nil {
		return nil, categoryRepositoryErr
	}
	return &GetCategoryByIDOutput{
		ID:          category.GetID().String(),
		Name:        category.GetName(),
		Description: category.GetDescription(),
		IsActive:    category.GetIsActive(),
		CreatedAt:   category.GetCreatedAt(),
	}, nil
}
