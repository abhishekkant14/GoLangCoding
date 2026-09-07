package main

import (
	"fmt"
	"time"
)

func PrintNum() {

	for i := 1; i <= 1000; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {

	go PrintNum()

	time.Sleep(time.Second * 5)

	fmt.Println("Main function completed")
}
