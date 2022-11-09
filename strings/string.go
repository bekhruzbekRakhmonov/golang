package main

import "fmt"
import "strconv"
import "reflect"

func main() {
	var num int = 6

	s1 := fmt.Sprintf("%v",num)

	fmt.Println(string(num),s1)

	s2 := strconv.Itoa(num)
	fmt.Println("Strconv",s2)

	switch v := reflect.ValueOf(s1); v.Kind() {
	case reflect.String:
		fmt.Println("Right converted.")
	default:
		fmt.Println("Wrong converted.")
	}

	s1Type := reflect.TypeOf(s1)
	s2Type := reflect.TypeOf(s2)

	fmt.Println("s1 and s2 types are: ",s1Type,s2Type)

}