package main

import "fmt"

func Paindrome(n int) bool {

	orignal := n
	reverse := 0

	for n > 0 {
		digit := n % 10

		reverse = reverse*10 + digit
		n = n / 10
	}
	return orignal == reverse
}
func main() {

	num := 121
	if Paindrome(num) {
		fmt.Println(num, "Paindrome Num")
	} else {
		fmt.Println(num, "Not a PaindromeNum")
	}
}
