package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(*os.Stderr)
	fmt.Println(*os.Stdin)
	fmt.Println(*os.Stdout)
}
