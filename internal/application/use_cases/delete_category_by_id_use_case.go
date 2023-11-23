package use_cases

import (
	"errors"

	"github.com/google/uuid"
	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/intwone/catalog/internal/domain/repositories"
)

type DeleteCategoryByIDInput struct {
	ID uuid.UUID
}

type DeleteCategoryByIDOutput struct{}

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
	categoryRepositoryErr := uc.CategoryRepository.DeleteByID(input.ID)
	switch {
	case errors.Is(categoryRepositoryErr, errs.ResourceNotFound):
		return nil, errs.ResourceNotFound
	case errors.Is(categoryRepositoryErr, errs.UnexpectedError):
		return nil, errs.UnexpectedError
	default:
		return &DeleteCategoryByIDOutput{}, nil
	}
}
