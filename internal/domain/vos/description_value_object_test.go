package vos_test

import (
	"strings"
	"testing"

	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/intwone/catalog/internal/domain/vos"
	"github.com/stretchr/testify/require"
)

func TestDescription(t *testing.T) {
	t.Run("should not be able to create a description too long", func(t *testing.T) {
		invalidDescription := strings.Repeat("test", 10001)
		description, err := vos.NewDescription(invalidDescription)

		require.Nil(t, description)
		require.NotNil(t, err)
		require.Equal(t, err.Error(), errs.TooGreatDescriptionError.Error())
	})
}
