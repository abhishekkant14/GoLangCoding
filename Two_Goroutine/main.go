package main

import (
	"fmt"
	"sync"
)

func even(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 2; i <= 10; i++ {
		fmt.Println("EvenNum", i)
	}

}

func Odd(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 9; i++ {
		fmt.Println("OddNum", i)
	}

}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go even(&wg)
	go Odd(&wg)

	wg.Wait()
	fmt.Println("main completed")
}
