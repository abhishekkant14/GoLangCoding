package main

import "fmt"

func factroial(s int) int {
	if s == 0 {
		return 1

	}

	return s * factroial(s-1)

}

func main() {

	fmt.Println(factroial(5))

}
