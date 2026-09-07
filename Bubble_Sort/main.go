package main

import "fmt"

func Bubble(arr []int) {

	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}
	}

}
func main() {

	arr := []int{2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Un_Sorted", arr)
	Bubble(arr)

	fmt.Println("Sorted_Arr", arr)

}
