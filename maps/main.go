package main

import "fmt"

func main(){
	// making map -> map[type]type

	// first way
	var mp map[string]int = map[string]int{"Bexruz": 18,"Feruz": 15,"Shoxruz": 12}

	fmt.Println("Orginal:",mp)

	// accessing data from map

	fmt.Println(mp["Feruz"])

	// delete data from map

	delete(mp,"Shoxruz")

	fmt.Println("Item deleted version:",mp)

	// second way
	mp2 := map[string]int{"Bexruz": 18,"Feruz": 15,"Shoxruz": 12}

	fmt.Println(mp2)

	// third way -> this will initialize empty map
	mp3 := make(map[string]int)

	// giving value to map

	mp3["Bexruz"] = 18

	fmt.Println(mp3)
}