package use_cases_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uc "github.com/intwone/catalog/internal/application/use_cases"
	"github.com/intwone/catalog/internal/domain/entities"
	"github.com/intwone/catalog/internal/tests/mocks"
	"github.com/stretchr/testify/require"
)

func TestDeactiveCategoryByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should be able to deactive a category by id", func(t *testing.T) {
		// Arange
		categoryName := "category name"
		categoryDescription := "category description"
		categoryIsActive := true
		category, _ := entities.NewCategory(categoryName, categoryDescription, categoryIsActive)
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().GetByID(gomock.Any()).Return(category, nil).AnyTimes().Times(1)
		categoryRepository.EXPECT().Update(gomock.Any()).Return(nil).AnyTimes().Times(1)
		useCase := uc.NewDeactiveCategoryByIDUseCase(categoryRepository)

		// Act
		input := uc.DeactiveCategoryByIDInput{
			ID: category.GetID().String(),
		}
		output, err := useCase.Execute(input)

		// Assert
		require.Nil(t, err)
		require.False(t, output.IsActive)
	})
}
