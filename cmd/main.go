package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloWorld)
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8000",
	}
	log.Fatal(srv.ListenAndServe())
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}
