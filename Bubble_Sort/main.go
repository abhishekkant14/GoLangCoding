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

	number := []int{1, 3, 2, 4, 6, 5, 7, 9, 8}

	fmt.Println("Un_Sorted", number)

	Bubble(number)
	fmt.Println("Sorted_Array", number)

}
