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
	resp, err := http.Get("http://localhost:3000/product/1")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res := Product{}
	json.Unmarshal(body, &res)
	fmt.Printf("Received %v", res)
}
