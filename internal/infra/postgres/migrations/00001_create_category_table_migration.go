package migrations

import (
	"database/sql"
)

func CreateCategoryTable(db *sql.DB) error {
	query := `
		CREATE TABLE "categories" (
			"category_id" UUID NOT NULL,
			"name" TEXT NOT NULL,
			"description" TEXT NOT NULL,
			"is_active" BOOLEAN NOT NULL,
			"created_at" TIMESTAMPTZ(6) NOT NULL,
			"updated_at" TIMESTAMPTZ(6) NOT NULL,

			CONSTRAINT "category_pkey" PRIMARY KEY ("category_id")
		);
	`
	_, err := db.Exec(query)
	return err
}
