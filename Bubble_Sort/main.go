package main

import "fmt"

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

	arr := []int{7, 3, 4, 1, 5, 8, 3}

	fmt.Println("total", arr)
	BubbleSort(arr)

	fmt.Println("After:", arr)

}
