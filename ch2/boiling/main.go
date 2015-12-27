// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 29.
//!+

// Boiling prints the boiling point of water.
package main 

import "fmt"

const boilingF = 212.0

func boiling_hello() {
	fmt.Println("hello boiling!")
}

func main() {
	/*
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	*/
	//f := boilingF
	//c := (f - 32) * 5 / 9
	//fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	c := (boilingF - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", boilingF, c)
	// Output:
	// boiling point = 212°F or 100°C
}

//!-
