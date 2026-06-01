package main

import (
	"fmt"
	"strings"
)

func main() {

	S := "Abhishek"

	count := 0

	for _, ch := range strings.ToLower(S) {

		if ch == 'a' || ch == 'i' || ch == 'e' || ch == 'o' || ch == 'u' {
			count++
		}

	}

	fmt.Println(count)

}
