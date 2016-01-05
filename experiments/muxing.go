package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/test", testHandler)

	http.ListenAndServe(":8080", r)
}

func homeHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func testHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test")
}