package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	resp, err := http.Get("http://localhost:3000/products")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := Product{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("Received %v", res)
	}
}
