package use_cases_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/catalog/internal/application/use_cases"
	"github.com/intwone/catalog/internal/tests/factories"
	"github.com/intwone/catalog/internal/tests/mocks"
	"github.com/stretchr/testify/require"
)

func TestSearchCategoryByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should be able to search a category", func(t *testing.T) {
		// Arange
		categories := factories.CategoryFactory(5, true)
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().Search(gomock.Any(), gomock.Any(), gomock.Any()).Return(categories, nil).AnyTimes().Times(1)
		useCase := uc.NewSearchCategoriesUseCase(categoryRepository)

		// Act
		input := uc.SearchCategoriesInput{
			Offset:  1,
			Limit:   20,
			Keyword: "",
		}
		_, err := useCase.Execute(input)

		// Assert
		require.Nil(t, err)
	})
}
