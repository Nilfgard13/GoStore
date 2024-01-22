package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nilfgard13/GOSTORE/app/model"
	"github.com/Nilfgard13/GOSTORE/database/seeder"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUSer     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome To " + appConfig.AppName)

	server.InitializeDB(dbConfig)
	server.InitializeRoutes()
	// seeder.DBSeed(server.DB)
}

func (server *Server) InitializeDB(dbConfig DBConfig) {
	var err error
	if dbConfig.DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUSer, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
		server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		// dsn := fmt.Sprintf("user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
		// server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("Failed on connecting to the database server")
	}

}

func (server *Server) dbMigrate() {
	for _, model := range model.RegisterModel() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database Migration Success")
}

func (server *Server) InitCommand(config AppConfig, dbConfig DBConfig) {
	server.InitializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeder.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
