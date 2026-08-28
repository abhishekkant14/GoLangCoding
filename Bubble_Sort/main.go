package main

import "fmt"

func Bubble_Sort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {

				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {

	arr := []int{4, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Un_Sorted", arr)

	Bubble_Sort(arr)

	fmt.Println("Sorted_arr", arr)

}
