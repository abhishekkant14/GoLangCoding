package main

import (
	"fmt"
	"sync"
	"time"
)

func Production(jobs chan<- int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 10; i++ {

		fmt.Println("Production:_", i)

		jobs <- i

		time.Sleep(time.Second * 2)

	}
	close(jobs)

}

func Consumed(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Consumed :__", job)

		time.Sleep(time.Second * 2)

	}

}

func main() {

	jobs := make(chan int, 5)

	var wg sync.WaitGroup

	wg.Add(1)

	go Production(jobs, &wg)

	wg.Add(1)
	go Consumed(jobs, &wg)

	wg.Wait()

	fmt.Println("All jobs Copleted")
}
