package repositories

import (
	"database/sql"

	"github.com/intwone/catalog/internal/domain/entities"
	"github.com/intwone/catalog/internal/domain/errs"
	"github.com/intwone/catalog/internal/infra/postgres/mappers"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) Save(c entities.Category) error {
	query := `insert into "categories" (category_id, name, description, is_active, created_at, updated_at) values ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, c.GetID().String(), c.GetName(), c.GetDescription(), c.GetIsActive(), c.GetCreatedAt(), c.GetUpdatedAt())
	if err != nil {
		return errs.UnexpectedError
	}
	return nil
}

func (r *CategoryRepository) Get() (*[]entities.Category, error) {
	query := `select * from "categories" where is_active = true`
	rows, queryErr := r.db.Query(query)
	if queryErr != nil {
		return nil, errs.UnexpectedError
	}
	defer rows.Close()
	categories, mapperErr := mappers.InfraToDomain(rows)
	if mapperErr != nil {
		return nil, errs.UnexpectedError
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(id string) (*entities.Category, error) {
	query := `select * from "categories" where category_id = $1 and is_active = true`
	rows, queryErr := r.db.Query(query, id)
	if queryErr != nil {
		return nil, errs.UnexpectedError
	}
	var category entities.Category
	if rows.Next() {
		scanErr := rows.Scan(&category.ID, &category.Name.Value, &category.Description.Value, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
		if scanErr != nil {
			return nil, errs.UnexpectedError
		}
	} else {
		return nil, errs.ResourceNotFound
	}
	return &category, nil
}

func (r *CategoryRepository) DeleteByID(id string) error {
	query := `update "categories" set is_active = false where category_id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return errs.UnexpectedError
	}
	rowsEffected, err := result.RowsAffected()
	if rowsEffected == 0 {
		return errs.ResourceNotFound
	}
	return nil
}

func (r *CategoryRepository) Update(category entities.Category) error {
	query := `update "categories" set name = $1, description = $2 where category_id = $3`
	result, err := r.db.Exec(query, category.GetName(), category.GetDescription(), category.GetID().String())
	if err != nil {
		return errs.UnexpectedError
	}
	rowsEffected, err := result.RowsAffected()
	if rowsEffected == 0 {
		return errs.ResourceNotFound
	}
	return nil
}

func (r *CategoryRepository) Search(offset int64, limit int64, keyword string) (*[]entities.Category, error) {
	query := `select * from "categories" where is_active = true and (name ilike '%' || $1 || '%' or description ilike '%' || $1 || '%') limit $2 offset $3`
	rows, queryErr := r.db.Query(query, keyword, limit, offset)
	if queryErr != nil {
		return nil, errs.UnexpectedError
	}
	defer rows.Close()
	categories, mapperErr := mappers.InfraToDomain(rows)
	if mapperErr != nil {
		return nil, errs.UnexpectedError
	}
	return categories, nil
}
