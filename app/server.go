package app

import (
	"flag"
	"os"

	"log"

	"github.com/Nilfgard13/GOSTORE/app/controller"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = controller.Server{}
	var appConfig = controller.AppConfig{}
	var dbConfig = controller.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error om loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoStore")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUSer = getEnv("DB_USER", "fachrizal")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	dbConfig.DBName = getEnv("DB_NAME", "gostoredb")
	dbConfig.DBPort = getEnv("DB_PORT", "3306")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "mysql")

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		server.InitCommand(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}

}
