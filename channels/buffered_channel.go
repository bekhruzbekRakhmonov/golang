// Simple program to demonstrate use of Buffered Channel

package main

import (
	"fmt"	
	"math/rand"
	"sync"
	"time"
)

var goRoutine sync.WaitGroup

func main(){
	rand.Seed(time.Now().Unix())

	// Create a buffered channel to manage the employee vs project load.
	projects := make(chan string,10)

	// Launch 5 goroutines to handle the projects.
	goRoutine.Add(5)
	for i :=1; i <= 5; i++ {
		go employee(projects, i)
	}

	for j :=1; j <= 10; j++ {
		projects <- fmt.Sprintf("Project :%d", j)
	}

	// Close the channel so the goroutines will quit	
	close(projects)
	goRoutine.Wait()
}

func employee(projects chan string, employee int) {
	defer goRoutine.Done()
	for {
		// Wait for project to be assigned.
		project, result := <-projects

		if result==false {
			// This means the channel is empty and closed.
			fmt.Printf("Employee : %d : Exit\n", employee)
			return
		}

		fmt.Printf("Employee : %d : Started   %s\n", employee, project)

		// Randomly wait to simulate work time.
		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// Display time to wait
		fmt.Println("\nTime to sleep",sleep,"ms\n")

		// Display project completed by employee.
		fmt.Printf("Employee : %d : Completed %s\n", employee, project)
	}

}

/*
Every time you run this program the output for this program will be 
different this is because of the random nature of the program and the Go scheduler.
In above program, a buffered channel of type string is created with a capacity of 10. 
WaitGroup is given the count of 5, one for each goroutine. 
10 strings are sent into the channel to simulate or replicate project for the goroutines. 
Once the last string is sent into the channel, 
the channel is going to be closed and the main function waits for 
all the project to be completed.
*/