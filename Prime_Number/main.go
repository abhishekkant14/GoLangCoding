package main

import "fmt"

func IsPrimeNum(num int) bool {

	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true

}

func main() {

	fmt.Println(IsPrimeNum(7))
	fmt.Println(IsPrimeNum(18))

}
