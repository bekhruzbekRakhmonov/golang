package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	// var input string

	// fmt.Print("Enter an input as string: ")
	// fmt.Scanln(&input)

	// fmt.Println("Output:",input)

	// var newInput int

	// fmt.Print("Enter an input as a number: ")
	// fmt.Scanf("%i",&newInput)
	// fmt.Println("Output:",newInput)


	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	outputReader := bufio.NewReader(os.Stdout)
	text, _ = outputReader.ReadString('\n')
	fmt.Println(text)
}