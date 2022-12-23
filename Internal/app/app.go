package app

import (
	"fmt"
	"log"
	"os"

	psql "github.com/ITOTDEL4U/CRUD/Internal/repository"
	"github.com/ITOTDEL4U/CRUD/pkg/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func Start() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.NewPosgresDBConnector(database.ConnectionInfo{

		Host:     os.Getenv("host"),
		Port:     os.Getenv("port"),
		Username: os.Getenv("user"),
		DBName:   os.Getenv("dbname"),
		SSLMode:  "disable",
		Password: os.Getenv("password"),
	})
	if err != nil {
		log.Fatal(err)
	}

	booksRepo := psql.NewBooks(db)
	fmt.Println(booksRepo)
}
