package use_cases_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/catalog/internal/application/use_cases"
	"github.com/intwone/catalog/internal/domain/entities"
	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/intwone/catalog/internal/tests/mocks"
	"github.com/stretchr/testify/require"
)

func TestGetCategoryByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should be able to get a category by id", func(t *testing.T) {
		// Arange
		categoryName := "category name"
		categoryDescription := "category description"
		categoryIsActive := true
		category, _ := entities.NewCategory(categoryName, categoryDescription, categoryIsActive)
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().GetByID(gomock.Any()).Return(category, nil).AnyTimes().Times(1)
		useCase := uc.NewGetCategoryByIDUseCase(categoryRepository)

		// Act
		input := uc.GetCategoryByIDInput{
			ID: category.GetID().String(),
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

	t.Run("should not be able to get a category when id not exists", func(t *testing.T) {
		// Arange
		categoryName := "category name"
		categoryDescription := "category description"
		categoryIsActive := true
		category, _ := entities.NewCategory(categoryName, categoryDescription, categoryIsActive)
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().GetByID(gomock.Any()).Return(nil, errs.ResourceNotFound).AnyTimes().Times(1)
		useCase := uc.NewGetCategoryByIDUseCase(categoryRepository)

		// Act
		input := uc.GetCategoryByIDInput{
			ID: category.GetID().String(),
		}
		output, err := useCase.Execute(input)

		// Assert
		require.Nil(t, output)
		require.NotNil(t, err)
		require.Equal(t, err.Error(), errs.ResourceNotFound.Error())
	})
}
