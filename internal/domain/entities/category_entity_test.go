package entities_test

import (
	"strings"
	"testing"

	"github.com/intwone/catalog/internal/domain/entities"
	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/stretchr/testify/require"
)

func TestCategory(t *testing.T) {
	t.Run("should be able to create a category", func(t *testing.T) {
		validName := "name test"
		validDescription := "description test"
		category, err := entities.NewCategory(validName, validDescription, true)

		require.NotNil(t, category)
		require.Nil(t, err)
		require.NotNil(t, category.GetID())
		require.Equal(t, category.GetName(), validName)
		require.Equal(t, category.GetDescription(), validDescription)
		require.True(t, category.GetIsActive())
	})

	t.Run("should not be able update category name with invalid name", func(t *testing.T) {
		validName := "name test"
		validDescription := "description test"
		invalidName := "n4m3 t3st3!#"
		category, _ := entities.NewCategory(validName, validDescription, true)
		result := category.SetName(invalidName)

		require.Equal(t, result.Error(), errs.InvalidNameCharactersError.Error())
		require.Equal(t, category.GetName(), validName)
	})

	t.Run("should not be able update category description with invalid description", func(t *testing.T) {
		validName := "name test"
		validDescription := "description test"
		invalidDescription := strings.Repeat("test", 10001)
		category, _ := entities.NewCategory(validName, validDescription, true)
		result := category.SetDescription(invalidDescription)

		require.Equal(t, result.Error(), errs.TooGreatDescriptionError.Error())
		require.Equal(t, category.GetDescription(), validDescription)
	})
}
