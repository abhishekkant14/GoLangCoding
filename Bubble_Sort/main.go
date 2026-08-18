package main

import (
	"fmt"
)

func BubbleSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

func main() {

	arr := []int{6, 5, 3, 2, 3, 1, 4, 7}

	fmt.Println("Orignal Arr", arr)

	BubbleSort(arr)

	fmt.Println("Sorted", arr)

}
