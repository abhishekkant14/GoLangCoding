package main

import "fmt"

func producer(ch chan int) {

	for i := 1; i <= 5; i++ {

		ch <- 1
	}
	close(ch)
}

func cunsumer(ch chan int) {

	for value := range ch {
		fmt.Println("consumed", value)

	}

}

func main() {
	ch := make(chan int)

	go producer(ch)

	cunsumer(ch)

}
