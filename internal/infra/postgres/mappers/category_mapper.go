package mappers

import (
	"database/sql"

	"github.com/intwone/catalog/internal/domain/entities"
)

func InfraToDomain(rows *sql.Rows) (*[]entities.Category, error) {
	var categories []entities.Category
	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(
			&category.ID,
			&category.Name.Value,
			&category.Description.Value,
			&category.IsActive,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return &categories, nil
}
