// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"gopl.io/ch2/tempconv"
)

func main() {
	cmdArgs := os.Args[1:]
	if len(cmdArgs) != 0 {
		for _, arg := range cmdArgs {
			conv(arg)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			conv(input.Text())
		}
	}
}

func conv(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

//!-
