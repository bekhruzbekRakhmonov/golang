package main

import "sync"
import "fmt"

func counter(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < i; j++ {
		fmt.Println(j)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go counter(100, &wg)
	go counter(100, &wg)

	// this also does samething like above:
	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)
	// 	go counter(10, &wg)
	// }

	wg.Wait()

	fmt.Println("Finished")
}