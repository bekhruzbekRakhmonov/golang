package main

import "fmt"
import "rsc.io/quote"

func main(){
	fmt.Println("hello world")
	fmt.Println(quote.Go())

	/* 
	while loop
	for i:=1;true;i++{
	        fmt.Println(i)
	}
	
	for {
		fmt.Println("For loop")
	}
	*/
	fmt.Println("quote-hello:",quote.Hello())

}
