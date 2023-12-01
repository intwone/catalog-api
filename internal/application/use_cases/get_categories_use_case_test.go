package use_cases_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/catalog/internal/application/use_cases"
	"github.com/intwone/catalog/internal/tests/factories"
	"github.com/intwone/catalog/internal/tests/mocks"
	"github.com/stretchr/testify/require"
)

func TestGetCategoriesUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should be able to get all categories", func(t *testing.T) {
		// Arange
		categories := factories.CategoryFactory(5, true)
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().Get().Return(categories, nil).AnyTimes().Times(1)
		useCase := uc.NewGetCategoriesUseCase(categoryRepository)

		// Act
		input := uc.GetCategoriesInput{}
		output, err := useCase.Execute(input)

		// Assert
		require.Nil(t, err)
		require.Equal(t, len(*categories), len(*output))
	})
}
