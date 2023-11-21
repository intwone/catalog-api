package vos

import "github.com/intwone/catalog/internal/domain/errs"

type Name struct {
	Value string
}

func NewName(value string) (*Name, error) {
	name := &Name{Value: value}
	if err := name.isValid(); err != nil {
		return nil, err
	}
	return name, nil
}

func (d *Name) isValid() error {
	if d.isTooShort() {
		return errs.TooShortNameError
	}
	if d.isTooLong() {
		return errs.TooGreatNameError
	}
	if d.invalidCharacters() {
		return errs.InvalidNameCharactersError
	}
	return nil
}

func (n *Name) isTooShort() bool {
	return len(n.Value) < 3
}

func (n *Name) isTooLong() bool {
	return len(n.Value) > 255
}

func (n *Name) invalidCharacters() bool {
	return false
}
