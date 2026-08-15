package main

import (
	"fmt"
	"sync"
)

func Printeven(ch chan int) {

	defer close(ch)

	for i := 2; i <= 10; i++ {
		ch <- i
	}

}
func PrintOdd(ch chan int) {
	defer close(ch)

	for i := 1; i <= 9; i++ {
		ch <- i
	}

}
func PrintPrime(ch chan int) {
	defer close(ch)

	Prime := []int{2, 3, 5, 7}

	for _, p := range Prime {
		ch <- p
	}

}

func fanIn(channels ...<-chan int) <-chan int {

	out := make(chan int)

	var wg sync.WaitGroup

	for _, ch := range channels {

		wg.Add(1)

		go func(c <-chan int) {
			defer wg.Done()

			for i := range c {

				out <- i
			}

		}(ch)

	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}
func main() {

	evenCh := make(chan int)
	OddCh := make(chan int)
	PrimeCh := make(chan int)

	go Printeven(evenCh)

	go PrintOdd(OddCh)

	go PrintPrime(PrimeCh)

	marged := fanIn(evenCh, OddCh, PrimeCh)

	for value := range marged {
		fmt.Println(value)
	}

}
