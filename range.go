package main

import "fmt"

func main(){
	// arrays are immutable
	var a [2]int = [2]int{5,6}
	new_array := a
	new_array[1] = 10
	fmt.Printf("Arrays: %d,%d\n",new_array,a)

	// slices are mutable
	var s []int =  []int{5,6,9,8,45,12,6,45}
	
	// slices are mutable
	var ss []int = []int{5,6}
	new_slices := ss
	new_slices[1] = 10
	fmt.Printf("Slices: %d,%d\n",new_slices,ss)

	for i:=0; i < len(s); i++{
		fmt.Printf("%d\n",s[i])
	}

	for idx,element := range s{
		fmt.Printf("%d: %d\n",idx,element)
	}

	findDuplicatedElements(s)
	findDuplicatedElements(s)
}

func findDuplicatedElements(slices []int){
	for idx,element := range slices{
		for idx2, element2 := range slices{
			if  element == element2 && idx2 > idx{
				fmt.Println(element)
			}
		}
	}
}

func findDuplicatedElements2(slices []int){
	for i, element := range slices{
		for j := i + 1; j < len(slices); j++{
			if element == slices[j]{
				fmt.Printf("Duplicated element is %d.",element)
			}
		}
	}
}
