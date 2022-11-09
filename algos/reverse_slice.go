package main

import "fmt"

func reverse[T int | string](slice []T) []T {
	var reversed []T  
	for i := len(slice) - 1; i > -1; i-- {
		reversed = append(reversed,slice[i])
	}
	return reversed
} 

func reverse_inplace[T int | string](slice []T) []T {
	// var reversed []T = new([]T,len(slice))

	length := len(slice)

	for i := 0; i < length / 2; i++ {
		j := length - i - 1
		tmp := slice[i]
		slice[i],slice[j] = slice[j], tmp
	}

	return slice
}

func main() {
	var slice []int = []int{1,2,3,4,5,6}

	reversed := reverse[int](slice)
	fmt.Println(reversed)

	var slice2 []int = []int{1,2,3,4,5,6,7,8}

	reversed_inplace := reverse_inplace[int](slice2)
	fmt.Println(reversed_inplace)
}