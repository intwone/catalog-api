package use_cases_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/catalog/internal/application/use_cases"
	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/intwone/catalog/internal/tests/mocks"
	"github.com/stretchr/testify/require"
)

func TestCreateCategoryUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should be able to create a category", func(t *testing.T) {
		// Arange
		categoryName := "category name"
		categoryDescription := "category description"
		categoryIsActive := true
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes().Times(1)
		useCase := uc.NewCreateCategoryUseCase(categoryRepository)

		// Act
		input := uc.CreateCategoryInput{
			Name:        categoryName,
			Description: categoryDescription,
			IsActive:    categoryIsActive,
		}
		output, err := useCase.Execute(input)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, output.ID)
		require.Equal(t, output.Name, categoryName)
		require.Equal(t, output.Description, categoryDescription)
		require.True(t, output.IsActive)
		require.NotNil(t, output.CreatedAt)
	})

	t.Run("should not be able to create a category with invalid input", func(t *testing.T) {
		// Arange
		categoryName := ""
		categoryDescription := "category description"
		categoryIsActive := true
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes().Times(0)
		useCase := uc.NewCreateCategoryUseCase(categoryRepository)

		// Act
		input := uc.CreateCategoryInput{
			Name:        categoryName,
			Description: categoryDescription,
			IsActive:    categoryIsActive,
		}
		output, err := useCase.Execute(input)

		// Assert
		require.NotNil(t, err)
		require.Nil(t, output)
		require.Equal(t, err.Error(), errs.TooShortNameError.Error())
	})
}
