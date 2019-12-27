package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Specification struct {
		DBHost           string
		DBName           string
		DBSellerUsersCol string
		APIPort          string
	}
)

func Spec() *Specification {
	godotenv.Load("dev.env")

	s := Specification{
		DBHost:           os.Getenv("DB_HOST"),
		DBName:           os.Getenv("DB_NAME"),
		DBSellerUsersCol: os.Getenv("DB_SELLERUSERS_COL"),
		APIPort:          os.Getenv("API_PORT"),
	}
	return &s
}
