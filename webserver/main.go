//Simple Product controller
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"strconv"
)

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var productId = 10
var products = make([]Product, 10)

//populate test users
func init() {
	createTestUsers()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/products", productsHandler)
	r.HandleFunc("/product/{productId}", productHandler).Methods("GET")
	r.HandleFunc("/product", productCreateHandler).Methods("POST")

	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ready\n")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Products\n")
	fmt.Fprintf(w, "The products are %v\n", products)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	fmt.Fprintf(w, "You selected %s\n", productId)
}

func createTestUsers() {
	products[0] = Product{"1", "Daniel", "daniel.bryant@test.com"}
	products[1] = Product{"2", "Ashley", "ashley@test.com"}
	products[2] = Product{"3", "Rusty", "rusty@test.com"}
}

func productCreateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		//Don't write or else status will be set automatically, and we
		//get an error "http: multiple response.WriteHeader calls"
		//fmt.Fprint(w, "Invalid data\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var id int
	name := r.FormValue("name")
	email := r.FormValue("email")

	if (name == "" || email == "") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id = productId
	productId++

	newProduct := Product{strconv.Itoa(id), name, email}
	products[3] = newProduct
	w.WriteHeader(http.StatusCreated)
}