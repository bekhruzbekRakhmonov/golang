package main

import "fmt"

func main() {

	// this will stop if i < 50 and exits

	j := 45

	switch {
	case j < 10:
		fmt.Println("j is less than 10")
	case j < 50:
		fmt.Println("j is less than 50")
	case j < 100:
		fmt.Println("j is less than 100")
	}


	fmt.Println("--------------------------")

	// this will look at next case also

	i := 45
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough
	case i < 100:
		fmt.Println("i is less than 100")
	}
}