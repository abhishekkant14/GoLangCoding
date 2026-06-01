package main

import "fmt"

func main() {
	a := 0
	b := 2
	n := 20

	for i := 0; i < n; i++ {
		fmt.Println(a, "")
		a, b = b, a+b

	}

}
