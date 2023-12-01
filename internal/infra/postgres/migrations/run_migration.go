package migrations

import "database/sql"

func RunMigration(db *sql.DB) error {
	migrations := []func(*sql.DB) error{
		CreateCategoryTable,
	}
	for _, migration := range migrations {
		if err := migration(db); err != nil {
			return err
		}
	}
	return nil
}
