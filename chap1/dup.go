//prints the text of each line that appears more than once

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]map[string]bool)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}

	}
	for line, n := range counts {
		if n > 1 {
			var filenameStr string
			for key, _ := range filenames[line] {
				filenameStr = filenameStr + " " + key
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, filenameStr)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if filenames[input.Text()] == nil {
			filenames[input.Text()] = make(map[string]bool)
		}
		filenames[input.Text()][f.Name()] = true
	}
	//note: ignoring potential errors from input.Err()
}
