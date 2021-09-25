package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = a + 200
	fmt.Println(a)
}