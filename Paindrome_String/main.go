package main

import "fmt"

func Paindrome(a string) bool {

	rev := []rune(a)

	for i, a := 0, len(rev)-1; i < a; i, a = i+1, a-1 {

		if rev[i] != rev[a] {

			return false

		}

	}
	return true

}
func main() {
	fmt.Println(Paindrome("madam"))
	fmt.Println(Paindrome("Hello"))

}
