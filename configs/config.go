package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn  string
	Auth AuthConfig
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config { //вернули структуру Config кот содержит Db, Auth
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading env file, using default config")
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}
