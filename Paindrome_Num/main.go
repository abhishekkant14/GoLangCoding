package main

import "fmt"

func Paindrom(a int) bool {

	oiginal := a

	rev := 0

	for a > 0 {
		digit := a % 10
		rev = rev*10 + digit
		a = a / 10
	}
	return oiginal == rev
}
func main() {

	fmt.Println(Paindrom(121))
	fmt.Println(Paindrom(123))

}
