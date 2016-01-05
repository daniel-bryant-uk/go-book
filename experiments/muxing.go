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
	r.HandleFunc("/products", productHandler).Methods("GET")
	r.HandleFunc("/header", headerHandler)

	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test")
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Products")
	fmt.Fprintf(w, "HTTP Method %s", r.Method)
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Headers:")
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s : %s\n", k, v)
	}

	//if making request via localhost (and hence relatively), then scheme may not be set
	fmt.Fprintf(w, "IsAbs? %t\n", r.URL.IsAbs())
	fmt.Fprintf(w, "Scheme %s\n", r.URL.Scheme)
}