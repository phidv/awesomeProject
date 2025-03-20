package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Env     string
	AppName string

	PostgresURL      string
	PostgresDatabase string

	RestPort string

	Version string

	JwtSecret string
}

var Global = Env{}

func InitApp() {
	LoadEnv()
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		log.Fatal("----------Error loading .env file-----------")
	}

	Global.Env = os.Getenv("ENV")
	Global.AppName = os.Getenv("APP_NAME")

	Global.PostgresURL = os.Getenv("POSTGRES_URL")
	Global.PostgresDatabase = os.Getenv("POSTGRES_DATABASE")

	Global.RestPort = os.Getenv("REST_PORT")

	Global.Version = os.Getenv("API_VERSION")

	Global.JwtSecret = os.Getenv("JWT_SECRET")
}
