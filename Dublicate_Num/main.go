package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3, 4, 7, 3, 6, 8, 9, 1, 2, 4}

	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			fmt.Println("Dublicate", num)
		} else {
			seen[num] = true
		}
	}

}
