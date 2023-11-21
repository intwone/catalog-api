package vos

import "github.com/intwone/catalog/internal/domain/errs"

type Description struct {
	Value string
}

func NewDescription(value string) (*Description, error) {
	description := &Description{Value: value}
	if err := description.isValid(); err != nil {
		return nil, err
	}
	return description, nil
}

func (d *Description) isValid() error {
	if d.isTooLong() {
		return errs.TooGreatDescriptionError
	}
	return nil
}

func (d *Description) isTooLong() bool {
	return len(d.Value) > 10000
}
