package posgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // import for side effect
)

func initTables(db *sql.DB) {
	command := "CREATE TABLE IF NOT EXISTS Exchange_rates (Period timestamp ,Name VARCHAR(255),ShortName VARCHAR(3),Value numeric(10,6));"
	_, err := db.Exec(command)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewPosgresDBConnector() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	CheckError(err)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	return db

}

func AddToDB(db *sql.DB) {
	command := "INSERT INTO Exchange_rates "
	_, err := db.Exec(command)
	CheckError(err)
}
