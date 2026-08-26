package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 3, 4, 5, 2, 5, 7, 6, 8, 9}

	seen := make(map[int]bool)

	for _, value := range nums {
		if seen[value] {

			fmt.Println("Number", value)

		} else {

			seen[value] = true

		}

	}

}
