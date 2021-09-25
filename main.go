package main

import "fmt"

type person struct {
	name string
	age int
	favFoods []string
}

func main() {
	favFoods := []string{"ramen", "sandwich", "tomato"}
	jeeyoung := person{
		name: "jeeyoung", 
		age: 28, 
		favFoods: favFoods,
	}
	fmt.Println(jeeyoung)
}