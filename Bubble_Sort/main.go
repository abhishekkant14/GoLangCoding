package main

import "fmt"

func Bubble(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}

	}
}

func main() {

	arr := []int{1, 3, 5, 7, 8, 2, 3, 4, 5, 7, 8}

	fmt.Println("Number Without_Sorted", arr)

	Bubble(arr)

	fmt.Println("sorted_Num ", arr)

}
