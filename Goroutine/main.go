package main

import (
	"fmt"
	"sync"
)

var (
	totL int

	MU sync.Mutex
)

func Worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d Processing %d\n", id, job)
		result := job * job
		MU.Lock()
		totL += result
		MU.Unlock()
		results <- result

	}
}

func main() {

	jobs := make(chan int, 20)

	results := make(chan int, 5)

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {

		wg.Add(1)

		go Worker(i, jobs, results, &wg)
	}

	go func() {
		for j := 1; j <= 20; j++ {
			jobs <- j
		}
		close(jobs)

	}()
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Println("Result", result)
	}
	fmt.Println("Total Worker", totL)
}
