package main

import "fmt"

func main(){
	// slices are mutable
	s := []int{4,5,5,55,8,99,845,4}
	fmt.Printf("%v Length is %v\n",s,len(s))

	ss := s

	ss[0] = 78

	fmt.Printf("%v Length is %v\n",s,len(s))
}