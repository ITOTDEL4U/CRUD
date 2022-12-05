package webserver

import (
	"fmt"
    "github.com/ITOTDEL4U/CRUD/Internal/rates"
    "log"
	"net/http"

	_ "github.com/lib/pq" // import for side effect
)

func Start() {
	fmt.Println("Successfully connected!")

	http.HandleFunc("", handleUsers)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.Method == http.MethodGet:
        rates.MethodGet();
	case r.Method == http.MethodPost:
        rates.MethodPost();
	case r.Method == http.MethodPut:
        rates.MethodPut();
	case r.Method == http.MethodDelete:
        rates.MethodDelete();
	default:
		w.WriteHeader(http.StatusNotImplemented)

	}
	w.Write([]byte("hello world\n"))
}
