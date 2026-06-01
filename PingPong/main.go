package main

import (
	"fmt"
	"math/rand"
	"time"
)

func player(name string, table chan string) {
	for {
		ball, ok := <-table
		if !ok {
			fmt.Println(name, "is winner!")
			return
		}
		fmt.Println(name, "hit the ball", ball)
		time.Sleep(time.Millisecond * 200)

		if rand.Intn(10) == 0 {
			fmt.Println(name, "missed the ball", ball)

			close(table)
			return

		}
		table <- ball
	}

}

func main() {

	rand.Seed(time.Now().UnixNano())

	table := make(chan string)

	go player("Player 1", table)
	go player("Player 2", table)

	table <- "ping"

	time.Sleep(time.Second * 10)

}
