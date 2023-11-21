package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/intwone/catalog/internal/domain/vos"
)

type Catalog struct {
	id          uuid.UUID
	name        vos.Name
	description vos.Description
	isActive    bool
	createdAt   time.Time
	updatedAt   time.Time
}

func NewCatalog(name string, description string, isActive bool) (*Catalog, error) {
	voName, nameErr := vos.NewName(name)
	if nameErr != nil {
		return nil, nameErr
	}
	voDescription, descriptionErr := vos.NewDescription(description)
	if descriptionErr != nil {
		return nil, descriptionErr
	}
	return &Catalog{
		id:          uuid.New(),
		name:        *voName,
		description: *voDescription,
		isActive:    isActive,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

// Methods

func (c *Catalog) Active() error {
	if c.isActive {
		return errs.CannotActiveCatalogError
	}
	c.isActive = true
	return nil
}

func (c *Catalog) Deactive() error {
	if !c.isActive {
		return errs.CannotDeactiveCatalogError
	}
	c.isActive = false
	return nil
}

// Getters

func (c *Catalog) GetID() uuid.UUID {
	return c.id
}

func (c *Catalog) GetName() string {
	return c.name.Value
}

func (c *Catalog) GetDescription() string {
	return c.description.Value
}

func (c *Catalog) GetIsActive() bool {
	return c.isActive
}

func (c *Catalog) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *Catalog) GetUpdatedAt() time.Time {
	return c.updatedAt
}
