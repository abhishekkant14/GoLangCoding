package main

import "fmt"

func main() {
	arr := []int{22, 23, 23, 231, 43, 34, 443, 2, 4, 6}

	max := arr[0]
	Second := arr[0]

	for _, v := range arr {

		if v > max {

			Second = max
			max = v
		} else if v > Second {
			Second = v
		}

	}
	fmt.Println(max)
	fmt.Println(Second)

}
