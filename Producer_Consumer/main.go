package main

import "fmt"

func Producer(ch chan int) {

	for i := 0; i <= 5; i++ {

		fmt.Println("Produced:", i)
		ch <- i
	}
	close(ch)

}
func Counsumer(ch chan int) {

	for value := range ch {

		fmt.Println("Cousmed", value)
	}

}
func main() {

	ch := make(chan int)

	go Producer(ch)

	Counsumer(ch)

	fmt.Println("Processing completed")

}
