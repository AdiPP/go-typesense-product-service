package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName string `env:"APP_NAME"`
	AppEnv  string `env:"APP_ENV" envDefault:"production"`
	AppHost string `env:"APP_HOST"`
	AppPort int    `env:"APP_PORT"`

	PgsqlDatabaseHost     string `env:"PGSQL_DB_HOST"`
	PgsqlDatabasePort     string `env:"PGSQL_DB_PORT"`
	PgsqlDatabaseDatabase string `env:"PGSQL_DB_DATABASE"`
	PgsqlDatabaseUsername string `env:"PGSQL_DB_USERNAME"`
	PgsqlDatabasePassword string `env:"PGSQL_DB_PASSWORD"`
	PgsqlDatabaseSchema   string `env:"PGSQL_DB_SCHEMA"`

	TypesenseDatabaseHost   string `env:"TYPESENSE_DB_HOST"`
	TypesenseDatabasePort   string `env:"TYPESENSE_DB_PORT"`
	TypesenseDatabaseAPIKey string `env:"TYPESENSE_DB_API_KEY"`
}

func (c Config) GetPgsqlConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		c.PgsqlDatabaseUsername,
		c.PgsqlDatabasePassword,
		c.PgsqlDatabaseHost,
		c.PgsqlDatabasePort,
		c.PgsqlDatabaseDatabase,
		c.PgsqlDatabaseSchema,
	)
}

func (c Config) GetPgsqlConnStringCdc() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?replication=database",
		c.PgsqlDatabaseUsername,
		c.PgsqlDatabasePassword,
		c.PgsqlDatabaseHost,
		c.PgsqlDatabasePort,
		c.PgsqlDatabaseDatabase,
	)
}

func (c Config) GetAppEnv() string {
	environment := ""

	switch c.AppEnv {
	case "dev", "development":
		environment = "development"
	case "stg", "staging":
		environment = "staging"
	default:
		environment = "production"
	}

	return environment
}

func LoadConfig() (cfg *Config, err error) {
	if err = godotenv.Load(); err != nil {
		log.Fatal("Warning: No .env file found, using system environment variables")
	}

	cfg = &Config{}
	if err = env.Parse(cfg); err != nil {
		log.Fatalf("Failed to parse config: %s", err)
	}

	log.Print("Success load config")
	return
}
