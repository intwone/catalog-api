package entities

import (
	"time"

	"github.com/google/uuid"
	seedwork "github.com/intwone/catalog/internal/domain/seed-work"
	"github.com/intwone/catalog/internal/domain/vos"
)

type Category struct {
	seedwork.AggregateRoot
	Name        vos.Name
	Description vos.Description
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
		Name:          *newName,
		Description:   *newDescription,
		IsActive:      isActive,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}

// Methods

func (c *Category) update() {
	c.UpdatedAt = time.Now()
}

// Getters

func (c *Category) GetID() uuid.UUID {
	return c.AggregateRoot.Entity.ID
}

func (c *Category) GetName() string {
	return c.Name.Value
}

func (c *Category) GetDescription() string {
	return c.Description.Value
}

func (c *Category) GetIsActive() bool {
	return c.IsActive
}

func (c *Category) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *Category) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}

// Setters

func (c *Category) SetName(name string) error {
	newName, err := vos.NewName(name)
	if err != nil {
		return err
	}
	c.update()
	c.Name = *newName
	return nil
}

func (c *Category) SetDescription(description string) error {
	newDescription, err := vos.NewDescription(description)
	if err != nil {
		return err
	}
	c.update()
	c.Description = *newDescription
	return nil
}

func (c *Category) SetIsActive(isActive bool) {
	c.update()
	c.IsActive = isActive
}
