package main

import "fmt"

func main() {

	num := []int{1, 2, 33, 56, 76, 35}

	LargestNum := num[0]
	SecondNum := num[0]

	for _, value := range num {

		if value > LargestNum {
			SecondNum = LargestNum
			LargestNum = value
		} else if value > SecondNum && value != LargestNum {
			SecondNum = value
		}
	}
	fmt.Println("Largest", LargestNum)
	fmt.Println("SecondNum", SecondNum)
}
