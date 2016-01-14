package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"io/ioutil"
	"time"
)

var counter = 1

func main() {
	go startClient()
	go startServer()

	fmt.Println("Done")
	//TODO: Extremely hacky - replace this with WaitGroups
	fmt.Scanln()
}

func startClient() {
	for {
		r := mux.NewRouter()
		r.HandleFunc("/", homeHandler)
		fmt.Printf("Starting server...\n")
		log.Fatal(http.ListenAndServe(":3000", r))
	}
}

func startServer() {
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
}

///----

func homeHandler(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, "%v", counter)
}