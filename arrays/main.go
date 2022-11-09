package main

import "fmt"

func main(){
	var scores [5]uint16 = [5]uint16{86,88,87,95,90}

	fmt.Println(scores)

	// creating empty array

	var empty_array []uint16 = make([]uint16,5)

	// equivalent

	// empty_array := make([]uint16,5)

	fmt.Println(empty_array)


	// creating n + 1 dimensional arrays

	var arr [3][2]uint8 = [3][2]uint8{{5,8},{9,8},{14,25}}

	fmt.Println(arr)

	// creating empty n + 1 dimensional arrays

	var empty_arr [][]uint8 = make([][]uint8, 3, 100) // -> make(type,size,capacity)

	// equivalent

	// empty_arr := make([][]uint8,3,100)

	fmt.Println(empty_arr)

	// length is not limited array
	s := [...]int{4,5,5,55,8,99,845,4}
	fmt.Printf("%v Length is %v\n",s,len(s))

}