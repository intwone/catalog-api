package vos_test

import (
	"strings"
	"testing"

	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/intwone/catalog/internal/domain/vos"
	"github.com/stretchr/testify/require"
)

func TestName(t *testing.T) {
	t.Run("should not be able to create a name", func(t *testing.T) {
		validName := "teste"
		name, err := vos.NewName(validName)

		require.NotNil(t, name)
		require.Nil(t, err)
		require.Equal(t, name.Value, validName)
	})

	t.Run("should not be able to create a name too short", func(t *testing.T) {
		invalidName := ""
		name, err := vos.NewName(invalidName)

		require.Nil(t, name)
		require.NotNil(t, err)
		require.Equal(t, err.Error(), errs.TooShortNameError.Error())
	})

	t.Run("should not be able to create a name too long", func(t *testing.T) {
		invalidName := strings.Repeat("test", 300)
		name, err := vos.NewName(invalidName)

		require.Nil(t, name)
		require.NotNil(t, err)
		require.Equal(t, err.Error(), errs.TooGreatNameError.Error())
	})

	t.Run("should not be able to create a name with numbers", func(t *testing.T) {
		invalidName := "t3ste"
		name, err := vos.NewName(invalidName)

		require.Nil(t, name)
		require.NotNil(t, err)
		require.Equal(t, err.Error(), errs.InvalidNameCharactersError.Error())
	})

	t.Run("should not be able to create a name with special characters", func(t *testing.T) {
		invalidName := "teste#@@!#"
		name, err := vos.NewName(invalidName)

		require.Nil(t, name)
		require.NotNil(t, err)
		require.Equal(t, err.Error(), errs.InvalidNameCharactersError.Error())
	})
}
