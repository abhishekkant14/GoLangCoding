package main

import "fmt"

func BubbleNum(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

func main() {

	num := []int{9, 6, 3, 8, 2, 7, 3, 2, 1, 4, 5}

	fmt.Println("ExistNum", num)

	BubbleNum(num)

	fmt.Println("Sorted arr", num)

}
