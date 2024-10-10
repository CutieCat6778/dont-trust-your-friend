package lib

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	JWT_SECRET  string
	POSTGES_URI string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	JWT_SECRET = GetEnv("JWT_SECRET")
	POSTGES_URI = GetEnv("POSTGES_URI")
}

func GetEnv(name string) string {
	data := os.Getenv(name)
	if data == "" {
		panic("Environment variable not found: " + name)
	}
	return data
}
