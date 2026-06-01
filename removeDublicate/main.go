package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3}

	m := make(map[int]bool)

	for _, v := range arr {

		if _, ok := m[v]; !ok {
			m[v] = true
			fmt.Println(v)
		}

	}

}
