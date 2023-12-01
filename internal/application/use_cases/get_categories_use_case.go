package use_cases

import (
	"time"

	"github.com/intwone/catalog/internal/application/errs"
	"github.com/intwone/catalog/internal/domain/repositories"
)

type GetCategoriesInput struct {
}

type GetCategoriesOutput struct {
	ID          string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

type GetCategoriesUseCaseInterface interface {
	Execute(input GetCategoriesInput) (*GetCategoriesOutput, error)
}

type GetCategoriesUseCase struct {
	CategoriesRepository repositories.CategoryRepositoryInterface
}

func NewGetCategoriesUseCase(cr repositories.CategoryRepositoryInterface) *GetCategoriesUseCase {
	return &GetCategoriesUseCase{
		CategoriesRepository: cr,
	}
}

func (uc *GetCategoriesUseCase) Execute(input GetCategoriesInput) (*[]GetCategoriesOutput, error) {
	categories, getCategoriesRepositoryErr := uc.CategoriesRepository.Get()
	if err := errs.HandleRepositoryError(getCategoriesRepositoryErr); err != nil {
		return nil, err
	}
	var categoriesOutput []GetCategoriesOutput
	for _, category := range *categories {
		categoryOutput := GetCategoriesOutput{
			ID:          category.GetID().String(),
			Name:        category.GetName(),
			Description: category.GetDescription(),
			IsActive:    category.GetIsActive(),
			CreatedAt:   category.GetCreatedAt(),
		}
		categoriesOutput = append(categoriesOutput, categoryOutput)
	}
	return &categoriesOutput, nil
}
