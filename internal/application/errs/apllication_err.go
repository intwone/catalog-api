package errs

import (
	"errors"

	"github.com/intwone/catalog/internal/domain/errs"
)

func HandleRepositoryError(e error) error {
	switch {
	case errors.Is(e, errs.ResourceNotFound):
		return errs.ResourceNotFound
	case errors.Is(e, errs.UnexpectedError):
		return errs.UnexpectedError
	}
	return nil
}
