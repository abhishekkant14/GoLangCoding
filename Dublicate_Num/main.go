package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 9, 4}

	seen := make(map[int]bool)

	for _, value := range nums {

		if seen[value] {
			fmt.Println("Nums", value)

		} else {
			seen[value] = true
		}

	}
}
