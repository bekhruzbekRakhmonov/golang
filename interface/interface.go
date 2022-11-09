package main

import "fmt"

// https://zetcode.com/golang/interface/#:~:text=The%20primary%20task%20of%20an,as%20exposed%20APIs%20or%20contracts.

func main(){
	var val interface {} = "Hello, World"
	v,ok := val.(string)
	v2,ok2 := val.(int)

	fmt.Println(v,ok)
	fmt.Println(v2,ok2)
}