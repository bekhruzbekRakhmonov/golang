package main

import "fmt"

func main(){

	// creating slices with new() is not good
	var members *[]string = new([]string)

	fmt.Println(members)

	fmt.Println(*members)


	// creating slices with make() is better
	var new_slices []string = make([]string,2) // -> make(type,size,capacity)

	fmt.Println(new_slices)

	new_slices[0] = "Hello"
	new_slices[1] = "World"

	fmt.Println(new_slices)

	// slices := *members

	// slices[0] = "hello"

	// fmt.Println(slices)
	// fmt.Println(members)

}