// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 1.

// Helloworld is our first Go program.
//!+
package main

import (
	"fmt"
	"unicode/utf8"
)

//import "os" //imported and not used: "os"

//if you put the left parenthess at the start of a new line, you would get the following error:
// syntax error: unexpected semicolon or newline before {
func main() {
	//var v string //v declared and not used
	fmt.Println("Hello, Haiyan Welcome to the 世界!")
	s := "世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}

//!-
