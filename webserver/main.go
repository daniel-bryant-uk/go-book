package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/products", productsHandler)
	r.HandleFunc("/product/{productId}", productHandler)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ready")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Products")
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	fmt.Fprintf(w, "You selected %s", productId)
}