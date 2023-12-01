package use_cases

import (
	"time"

	"github.com/intwone/catalog/internal/application/errs"
	"github.com/intwone/catalog/internal/domain/repositories"
)

type SearchCategoriesInput struct {
	Offset  int64
	Limit   int64
	Keyword string
}

type SearchCategoriesOutput struct {
	ID          string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
}

type SearchCategoriesUseCaseInterface interface {
	Execute(input SearchCategoriesInput) (*SearchCategoriesOutput, error)
}

type SearchCategoriesUseCase struct {
	CategoriesRepository repositories.CategoryRepositoryInterface
}

func NewSearchCategoriesUseCase(cr repositories.CategoryRepositoryInterface) *SearchCategoriesUseCase {
	return &SearchCategoriesUseCase{
		CategoriesRepository: cr,
	}
}

func (uc *SearchCategoriesUseCase) Execute(input SearchCategoriesInput) (*[]SearchCategoriesOutput, error) {
	categories, searchCategoriesRepositoryErr := uc.CategoriesRepository.Search(input.Offset, input.Limit, input.Keyword)
	if err := errs.HandleRepositoryError(searchCategoriesRepositoryErr); err != nil {
		return nil, err
	}
	var categoriesOutput []SearchCategoriesOutput
	for _, category := range *categories {
		categoryOutput := SearchCategoriesOutput{
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
