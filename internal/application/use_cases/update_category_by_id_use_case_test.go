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

func TestUpdateCategoryByIDUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should be able to update a category by id", func(t *testing.T) {
		// Arange
		categoryName := "category name"
		categoryDescription := "category description"
		categoryIsActive := true
		category, _ := entities.NewCategory(categoryName, categoryDescription, categoryIsActive)
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().GetByID(gomock.Any()).Return(category, nil).AnyTimes().Times(1)
		categoryRepository.EXPECT().Update(gomock.Any()).Return(nil).AnyTimes().Times(1)
		useCase := uc.NewUpdateCategoryByIDUseCase(categoryRepository)

		// Act
		input := uc.UpdateCategoryByIDInput{
			ID:          category.GetID().String(),
			Name:        "other category name",
			Description: "other category description",
		}
		output, err := useCase.Execute(input)

		// Assert
		require.Nil(t, err)
		require.Equal(t, "other category name", output.Name)
		require.Equal(t, "other category description", output.Description)
	})

	t.Run("should not be able to update a category when inputs are invalids", func(t *testing.T) {
		// Arange
		categoryName := "category name"
		categoryDescription := "category description"
		categoryIsActive := true
		category, _ := entities.NewCategory(categoryName, categoryDescription, categoryIsActive)
		categoryRepository := mocks.NewMockCategoryRepositoryInterface(ctrl)
		categoryRepository.EXPECT().GetByID(gomock.Any()).Return(category, nil).AnyTimes().Times(1)
		categoryRepository.EXPECT().Update(gomock.Any()).Return(nil).AnyTimes().Times(0)
		useCase := uc.NewUpdateCategoryByIDUseCase(categoryRepository)

		// Act
		input := uc.UpdateCategoryByIDInput{
			ID:          category.GetID().String(),
			Name:        "",
			Description: "other category description",
		}
		output, err := useCase.Execute(input)

		// Assert
		require.NotNil(t, err)
		require.Nil(t, output)
		require.Equal(t, errs.TooShortNameError.Error(), err.Error())
	})
}
