package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/intwone/catalog/internal/domain/errs"
	seedwork "github.com/intwone/catalog/internal/domain/seed-work"
	"github.com/intwone/catalog/internal/domain/vos"
)

type Category struct {
	seedwork.AggregateRoot
	name        vos.Name
	description vos.Description
	isActive    bool
	createdAt   time.Time
	updatedAt   time.Time
}

func NewCategory(name string, description string, isActive bool) (*Category, error) {
	newName, nameErr := vos.NewName(name)
	if nameErr != nil {
		return nil, nameErr
	}
	newDescription, descriptionErr := vos.NewDescription(description)
	if descriptionErr != nil {
		return nil, descriptionErr
	}
	return &Category{
		AggregateRoot: seedwork.AggregateRoot{Entity: *seedwork.NewEntity()},
		name:          *newName,
		description:   *newDescription,
		isActive:      isActive,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}, nil
}

// Methods

func (c *Category) Active() error {
	if c.isActive {
		return errs.CannotActiveCategoryError
	}
	c.isActive = true
	return nil
}

func (c *Category) Deactive() error {
	if !c.isActive {
		return errs.CannotDeactiveCategoryError
	}
	c.isActive = false
	return nil
}

func (c *Category) update() {
	c.updatedAt = time.Now()
}

// Getters

func (c *Category) GetID() uuid.UUID {
	return c.AggregateRoot.Entity.ID
}

func (c *Category) GetName() string {
	return c.name.Value
}

func (c *Category) GetDescription() string {
	return c.description.Value
}

func (c *Category) GetIsActive() bool {
	return c.isActive
}

func (c *Category) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *Category) GetUpdatedAt() time.Time {
	return c.updatedAt
}

// Setters

func (c *Category) SetName(name string) error {
	newName, err := vos.NewName(name)
	if err != nil {
		return err
	}
	c.update()
	c.name = *newName
	return nil
}

func (c *Category) SetDescription(description string) error {
	newDescription, err := vos.NewDescription(description)
	if err != nil {
		return err
	}
	c.update()
	c.description = *newDescription
	return nil
}
