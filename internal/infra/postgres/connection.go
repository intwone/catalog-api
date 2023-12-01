package postgres

import (
	"database/sql"
	"fmt"

	"github.com/intwone/catalog/internal/config"
	_ "github.com/lib/pq"
)

type Config struct {
	User    string
	Pass    string
	Host    string
	Port    string
	Name    string
	Dialect string
	SSLMode string
}

func NewConnection(c *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.Pass,
		c.User,
		c.Name,
		c.SSLMode,
	)
	db, err := sql.Open(c.Dialect, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Postgres connect successfully!")
	return db, nil
}

func GetConnection() (*sql.DB, error) {
	env := config.Env()
	config := &Config{
		Host:    env.DATABASE_HOST,
		Port:    env.DATABASE_PORT,
		User:    env.DATABASE_USER,
		Pass:    env.DATABASE_PASSWORD,
		Name:    env.DATABASE_NAME,
		Dialect: env.DATABASE_DIALECT,
		SSLMode: env.DATABASE_SSL_MODE,
	}
	db, err := NewConnection(config)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, nil
}
