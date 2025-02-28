package config

import (
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

func (a Config) GetAppEnv() string {
	environment := ""

	switch a.AppEnv {
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
	err = godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	cfg = &Config{}

	err = env.Parse(cfg)
	if err != nil {
		return
	}

	log.Println("Success load to config")
	return
}
