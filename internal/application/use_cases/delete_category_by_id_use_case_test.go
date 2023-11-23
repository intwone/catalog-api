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

func TestDeleteCategoryByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should be able to delete a category by id", func(t *testing.T) {
		// Arange
		categoryName := "category name"
		categoryDescription := "category description"
		categoryIsActive := true
		category, _ := entities.NewCategory(categoryName, categoryDescription, categoryIsActive)
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().DeleteByID(gomock.Any()).Return(nil).AnyTimes().Times(1)
		useCase := uc.NewDeleteCategoryByIDUseCase(categoryRepository)

		// Act
		input := uc.DeleteCategoryByIDInput{
			ID: category.GetID(),
		}
		_, err := useCase.Execute(input)

		// Assert
		require.Nil(t, err)
	})

	t.Run("should not be able to delete a category by id when category not exists", func(t *testing.T) {
		// Arange
		categoryName := "category name"
		categoryDescription := "category description"
		categoryIsActive := true
		category, _ := entities.NewCategory(categoryName, categoryDescription, categoryIsActive)
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().DeleteByID(gomock.Any()).Return(errs.ResourceNotFound).AnyTimes().Times(1)
		useCase := uc.NewDeleteCategoryByIDUseCase(categoryRepository)

		// Act
		input := uc.DeleteCategoryByIDInput{
			ID: category.GetID(),
		}
		_, err := useCase.Execute(input)

		// Assert
		require.NotNil(t, err)
		require.Equal(t, err.Error(), errs.ResourceNotFound.Error())
	})
}
