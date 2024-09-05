package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

const (
	errUndefinedEnvParam = "parameter is undefined"
)

type Config struct {
	Port       string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	port := os.Getenv("PORT")

	if port == "" {
		return nil, errors.New("PORT " + errUndefinedEnvParam)
	}

	dbUser := os.Getenv("DB_USER")

	if dbUser == "" {
		return nil, errors.New("DB_USER " + errUndefinedEnvParam)
	}

	dbPassword := os.Getenv("DB_PASSWORD")

	if dbPassword == "" {
		return nil, errors.New("DB_PASSWORD " + errUndefinedEnvParam)
	}

	dbHost := os.Getenv("DB_HOST")

	if dbHost == "" {
		return nil, errors.New("DB_HOST " + errUndefinedEnvParam)
	}

	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		return nil, errors.New("DB_PORT " + errUndefinedEnvParam)
	}

	dbName := os.Getenv("DB_NAME")

	if dbName == "" {
		return nil, errors.New("DB_NAME " + errUndefinedEnvParam)
	}

	return &Config{
		Port:       port,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
	}, nil
}
