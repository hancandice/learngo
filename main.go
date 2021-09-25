package main

import "fmt"

func main() {
	names := []string{"jeeyoung", "kyul", "sehyun"}
	names = append(names, "heeja")
	fmt.Println(names)
}