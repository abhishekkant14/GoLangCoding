package main

import (
	"fmt"
	"strings"
)

func Palindrome(n string) bool {
	palind := strings.ToLower(n)

	for i, j := 0, len(palind)-1; i < j; i, j = i+1, j-1 {
		if palind[i] != palind[j] {
			return false
		}

	}
	return true
}
func main() {
	fmt.Println(Palindrome("madam"))

}
