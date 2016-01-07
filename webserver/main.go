//Simple Product controller
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var products = make([]Product, 10)

//populate test users
func init() {
	createTestUsers()
}

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
	fmt.Fprintf(w, "The product is %v", products)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	fmt.Fprintf(w, "You selected %s", productId)
}

func createTestUsers() {
	products[0] = Product{"1", "Daniel", "daniel.bryant@test.com"}
	products[1] = Product{"2", "Ashley", "ashley@test.com"}
	products[2] = Product{"3", "Rusty", "rusty@test.com"}
}