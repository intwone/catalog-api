package entities_test

import (
	"testing"

	"github.com/intwone/catalog/internal/domain/entities"
	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/stretchr/testify/require"
)

func TestCatalog(t *testing.T) {
	t.Run("should be able to create a catalog", func(t *testing.T) {
		validName := "name test"
		validDescription := "description test"
		catalog, err := entities.NewCatalog(validName, validDescription, true)

		require.NotNil(t, catalog)
		require.Nil(t, err)
		require.NotNil(t, catalog.GetID())
		require.Equal(t, catalog.GetName(), validName)
		require.Equal(t, catalog.GetDescription(), validDescription)
		require.True(t, catalog.GetIsActive())
	})

	t.Run("should be able active a catalog", func(t *testing.T) {
		validName := "name test"
		validDescription := "description test"
		catalog, _ := entities.NewCatalog(validName, validDescription, false)
		result := catalog.Active()

		require.Nil(t, result)
		require.True(t, catalog.GetIsActive())
	})

	t.Run("should be able deactive a catalog", func(t *testing.T) {
		validName := "name test"
		validDescription := "description test"
		catalog, _ := entities.NewCatalog(validName, validDescription, true)
		result := catalog.Deactive()

		require.Nil(t, result)
		require.False(t, catalog.GetIsActive())
	})

	t.Run("should not be able active a catalog that is already active", func(t *testing.T) {
		validName := "name test"
		validDescription := "description test"
		catalog, _ := entities.NewCatalog(validName, validDescription, true)
		result := catalog.Active()

		require.Equal(t, result.Error(), errs.CannotActiveCatalogError.Error())
	})

	t.Run("should not be able deactive a catalog that is already deactive", func(t *testing.T) {
		validName := "name test"
		validDescription := "description test"
		catalog, _ := entities.NewCatalog(validName, validDescription, false)
		result := catalog.Deactive()

		require.Equal(t, result.Error(), errs.CannotDeactiveCatalogError.Error())
	})
}
