package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // import for side effect
)

type ConnectionInfo struct {
    Host     string
    Port     string
    Username string
    DBName   string
    SSLMode  string
    Password string
}
/*
func initTables(db *sql.DB) {
	command := "CREATE TABLE IF NOT EXISTS Exchange_rates (Period timestamp ,Name VARCHAR(255),ShortName VARCHAR(3),Value numeric(10,6));"
	_, err := db.Exec(command)
	CheckError(err)
}
*/


func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func NewPosgresDBConnector(info ConnectionInfo) (*sql.DB,error) {

    /*
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")

*/


    db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
        info.Host, info.Port, info.Username, info.DBName, info.SSLMode, info.Password))
    if err != nil {
        return nil, err
    }

    CheckError(err)

    return db, nil
    /*
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=%s",
        info.Host, info.Port, info.Username, info.Password, info.DBName,info.SSLMode)

	db, err := sql.Open("postgres", connectionString)




	CheckError(err)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	return db
*/


}
/*
func AddToDB(db *sql.DB) {
	command := "INSERT INTO Exchange_rates "
	_, err := db.Exec(command)
	CheckError(err)
}*/
