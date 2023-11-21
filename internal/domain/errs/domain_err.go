package errs

import "errors"

var (
	TooShortNameError          = errors.New("the name must have 3 or more characters")
	TooGreatNameError          = errors.New("the name should be 255 or fewer characters")
	InvalidNameCharactersError = errors.New("the name should not be special characters")
	TooGreatDescriptionError   = errors.New("the description should be 10000 or fewer characters")
	CannotActiveCatalogError   = errors.New("cannot activate catalog that is already active")
	CannotDeactiveCatalogError = errors.New("cannot deactivate catalog that is already deactive")
)
