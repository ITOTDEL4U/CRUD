package app

import (
	"fmt"
	"log"
	"os"

	//	"path"
	psql "github.com/ITOTDEL4U/CRUD/Internal/repository"
	"github.com/ITOTDEL4U/CRUD/Internal/service"
	"github.com/ITOTDEL4U/CRUD/pkg/database"

	//	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
) /*
func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func FindDir(filename string) string {
	dir, _ := os.Getwd()
	for !IsFile(path.Join(dir, filename)) {
		pdir := path.Dir(dir)
		if pdir == dir {
			return ""
		}
		dir = pdir
	}
	return dir
}
*/
func initializatin() {
	fmt.Println("sdf")
	/*
		var BaseDir = FindDir("go.mod")

		if BaseDir == "" {
			log.Fatalln("Please use `go mod init` to create project.")
			return
		}
		envFile := path.Join(BaseDir, ".env")
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Failed to load %s: %+v\n", envFile, err)
			return
		}
		// ....
	*/
}

func Start() {
	initializatin()

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

	Repo := psql.New(db)
	MyService := service.NewBooks(Repo)

	fmt.Println(MyService)
	fmt.Println(Repo)
	/*
		handler := rest.NewHandler(booksService)

		// init & run server
		srv := &http.Server{
			Addr:    ":8080",
			Handler: handler.InitRouter(),
		}

		log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	*/
}
