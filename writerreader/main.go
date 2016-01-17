package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"io/ioutil"
	"time"
	"sync"
)

var counter = 1
var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go startClient()
	go startServer()

	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Completed!")
}

func startServer() {
	for {
		r := mux.NewRouter()
		r.HandleFunc("/", homeHandler)
		fmt.Printf("Starting server...\n")
		log.Fatal(http.ListenAndServe(":3000", r))
	}
	wg.Done()
}

func startClient() {
	for {
		resp, err := http.Get("http://localhost:3000/")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s\n", body)
		time.Sleep(time.Second)
	}
	wg.Done()
}

///----

func homeHandler(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, "%v", counter)
}