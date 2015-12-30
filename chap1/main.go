// print command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	//1
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("Loop time -> %.2fs elapsed\n", time.Since(start).Seconds())

	//2
	s, sep = "", ""

	start = time.Now()

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("Join time -> %.2fs elapsed\n", time.Since(start).Seconds())

	//2.5
	s, sep = "",""
	fmt.Println(strings.Join(os.Args[1:], " "))

	//3
	s, sep = "", ""
	fmt.Println("Program: " + os.Args[0])
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	//4
	s, sep = "",""
	for idx, arg := range os.Args[1:] {
		fmt.Println("" + strconv.Itoa(idx) + ": " + arg)
		sep = " "
	}
}