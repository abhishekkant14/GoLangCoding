package main

import "fmt"

func main() {

	num := []int{2, 3, 4, 5, 66, 43, 54, 67, 54}

	LargestNum := num[0]
	SecondLargest := num[0]

	for _, value := range num {

		if value > LargestNum {
			SecondLargest = LargestNum
			LargestNum = value
		} else if value > SecondLargest {
			SecondLargest = value
		}
	}

	fmt.Println("LargestNum", LargestNum)
	fmt.Println("SecondNum", SecondLargest)

}
