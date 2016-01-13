package main

import (
	"os"
	"fmt"
	"encoding/json"
)

type Config struct {
	BooleanOption bool
	StringOption string
	ArrayOption []string
}

func main() {
	file, err := os.Open("config.json")
	if err!= nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Config: %v\n", configuration)

}
