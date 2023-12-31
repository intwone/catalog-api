package errs

import "errors"

var (
	TooShortNameError           = errors.New("the name must have 3 or more characters")
	TooGreatNameError           = errors.New("the name should be 255 or fewer characters")
	InvalidNameCharactersError  = errors.New("the name should not be special characters")
	TooGreatDescriptionError    = errors.New("the description should be 10000 or fewer characters")
	CannotActiveCategoryError   = errors.New("cannot activate category that is already active")
	CannotDeactiveCategoryError = errors.New("cannot deactivate category that is already deactive")
	ResourceNotFound            = errors.New("resource not found")
	UnexpectedError             = errors.New("unexpected error")
)
