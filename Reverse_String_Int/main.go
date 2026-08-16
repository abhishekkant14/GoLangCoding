package main

import "fmt"

func reverse(str string) string {

	rev := []rune(str)

	for i, a := 0, len(rev)-1; i < a; i, a = i+1, a-1 {
		rev[i], rev[a] = rev[a], rev[i]
	}
	return string(rev)

}
func main() {

	name := "Abhishek"
	fmt.Println(reverse(name))

	num := 12345678
	rev := 0

	for num > 0 {
		digit := num % 10
		rev = rev*10 + digit
		num = num / 10
	}
	fmt.Println(rev)

}
