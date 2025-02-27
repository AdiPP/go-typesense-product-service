package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App       App
	Typesense Typesense
}

type App struct {
	Name string `env:"APP_NAME"`
	Env  string `env:"APP_ENV"`
	Host string `env:"APP_HOST"`
	Port int    `env:"APP_PORT"`
}

func (a App) GetEnv() string {
	environment := ""

	switch a.Env {
	case "dev", "development":
		environment = "development"
	case "stg", "staging":
		environment = "staging"
	default:
		environment = "production"
	}

	return environment
}

type Typesense struct {
	Host   string
	Port   string
	APIKey string
}

func LoadConfig() (cfg *Config, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	cfg = &Config{
		App: App{
			Name: os.Getenv("APP_NAME"),
			Env:  os.Getenv("APP_ENV"),
			Port: func() int {
				port, err := strconv.Atoi(os.Getenv("APP_PORT"))
				if err != nil {
					log.Fatalf("Error converting APP_PORT to int: %v", err)
				}
				return port
			}(),
			Host: os.Getenv("APP_HOST"),
		},
		Typesense: Typesense{
			Host:   os.Getenv("TYPESENSE_HOST"),
			Port:   os.Getenv("TYPESENSE_PORT"),
			APIKey: os.Getenv("TYPESENSE_API_KEY"),
		},
	}
	return
}
