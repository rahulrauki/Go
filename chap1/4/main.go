// This program has the solution for Ex 1.4
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 { // meaning have to read the text instead of a file
		countLines(os.Stdin, counts, "os.Stdin") // os.Stdin has whatever the user has typed as text
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error while opening file : %v\n", err)
			} else {
				countLines(f, counts, args)
			}
			f.Close()
		}
	}
	for key, value := range counts {
		if value > 1 {
			splice := strings.Split(key, "#")
			fmt.Printf("Duplicate : %s , occurence : %d , File : %s\n", splice[0], value, splice[1])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()+"#"+fileName]++ // map :: defaultdict
	}
}
