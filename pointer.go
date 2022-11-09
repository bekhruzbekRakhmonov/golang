package main

import "fmt"

func main(){
	var name string = "Aslbek"
	addr := &name
	*addr = "Asadbek"
	fmt.Println(name)
}
