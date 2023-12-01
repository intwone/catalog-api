package factories

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/intwone/catalog/internal/domain/entities"
)

func CategoryFactory(count int, isActive bool) *[]entities.Category {
	fake := gofakeit.New(0)
	categories := make([]entities.Category, count)
	for i := 0; i < count; i++ {
		categories[i].SetName(fake.FirstName())
		categories[i].SetDescription(fake.Sentence(3))
		categories[i].SetIsActive(isActive)
	}
	return &categories
}
