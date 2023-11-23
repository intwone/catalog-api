package use_cases

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/intwone/catalog/internal/domain/repositories"
)

type GetCategoryByIDInput struct {
	ID uuid.UUID
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
	switch {
	case errors.Is(categoryRepositoryErr, errs.ResourceNotFound):
		return nil, errs.ResourceNotFound
	case errors.Is(categoryRepositoryErr, errs.UnexpectedError):
		return nil, errs.UnexpectedError
	default:
		return &GetCategoryByIDOutput{
			ID:          category.GetID().String(),
			Name:        category.GetName(),
			Description: category.GetDescription(),
			IsActive:    category.GetIsActive(),
			CreatedAt:   category.GetCreatedAt(),
		}, nil
	}
}
