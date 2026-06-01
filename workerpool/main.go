package main

import "fmt"

func worker(id int, job <-chan int, result chan<- int) {

	for j := range job {
		fmt.Printf("Worker %d finished job %d\n", id, j)
		result <- j * 2
	}
}

func main() {

	jobs := make(chan int, 10)

	result := make(chan int, 10)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	for j := 1; j <= 5; j++ {

		jobs <- j
	}
	close(jobs)

	for i := 1; i <= 5; i++ {
		fmt.Println("Result:", <-result)
	}

}
