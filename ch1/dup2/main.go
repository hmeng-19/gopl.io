// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	sources := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, sources, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, sources, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)

			for _, item := range sources[line] {
				fmt.Printf("%s\t", item)
			}

			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]int, sources map[string][]string, fname string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		sources[input.Text()] = append(sources[input.Text()], fname)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
