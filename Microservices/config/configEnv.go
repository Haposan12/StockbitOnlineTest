package config

import "github.com/joho/godotenv"

var ENV map[string]string

func LoadEnv() (err error) {
	ENV, err = godotenv.Read(".env")

	if err != nil {
		return
	}

	return
}
