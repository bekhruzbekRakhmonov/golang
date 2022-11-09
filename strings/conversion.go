package main

import "fmt"
import _ "strings"
import "strconv"

func main() {
	num := "127k"

	intVar, err := strconv.Atoi(num)
	if err != nil {
		fmt.Errorf("Could not convert to integer")
		return
	}

	fmt.Println("Converted int: ",intVar)

}