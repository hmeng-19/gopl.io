package main

import (
	"fmt"
	"time"
)

func main() {
	p := f()
	fmt.Println(p)
	fmt.Println(*p)
	hmeng := f()
	fmt.Println(hmeng)
	fmt.Println(*hmeng)
	fmt.Println(p)
	fmt.Println(*p)
	time.Sleep(200 * time.Second)
}

func f() *int {
	a := 2
	fmt.Println(&a)
	v := 1
	fmt.Println(&v)
	c := 3
	fmt.Println(&c)
	return &v
}
