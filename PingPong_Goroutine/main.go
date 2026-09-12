package main

import (
	"fmt"
	"sync"
	"time"
)

func Player1(ball chan int, result chan string, winner chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	score := 0

	for i := 1; i <= 5; i++ {

		n := <-ball

		fmt.Println("Player1 received the ball", n)

		time.Sleep(2 * time.Second)

		score++

		fmt.Println("Player1 hit the ball")

		// Final ball
		if n == 10 {
			winner <- "Player1"
		}

		// Send next ball only if this is not the final ball
		if n < 10 {
			ball <- n + 1
		}
	}

	result <- fmt.Sprintf("Player1 Score: %d", score)
}

func Player2(ball chan int, result chan string, winner chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	score := 0

	for i := 1; i <= 5; i++ {

		n := <-ball

		fmt.Println("Player2 received the ball", n)

		time.Sleep(2 * time.Second)

		score++

		fmt.Println("Player2 hit the ball")

		// Final ball
		if n == 10 {
			winner <- "Player2"
		}

		// Send next ball only if this is not the final ball
		if n < 10 {
			ball <- n + 1
		}
	}

	result <- fmt.Sprintf("Player2 Score: %d", score)
}

func main() {

	var wg sync.WaitGroup

	// Unbuffered channel
	ball := make(chan int)

	// Buffered channels
	result := make(chan string, 2)
	winner := make(chan string, 1)

	wg.Add(2)

	go Player1(ball, result, winner, &wg)
	go Player2(ball, result, winner, &wg)

	// Start the match
	ball <- 1

	// Wait for both players
	wg.Wait()

	// Get winner
	win := <-winner

	// Close result channel
	close(result)

	fmt.Println("\n========= GAME RESULT =========")

	for r := range result {
		fmt.Println(r)
	}

	fmt.Println("Winner:", win)
	fmt.Println("Finished")
}
