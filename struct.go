package main

import "fmt"

type User struct {
    name string
    age uint
}
// https://go.dev/tour/moretypes/://go.dev/tour/moretypes/8
func main(){
	user := User{"Bexruz Raxmonov",18}
	fmt.Println(user.name,user.age)
}
