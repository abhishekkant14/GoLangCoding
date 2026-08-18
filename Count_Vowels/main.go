package main

import "fmt"

func CountVowels(str string) int {

	count := 0

	for _, ch := range str {

		if ch == 'a' || ch == 'e' || ch == 'i' ||
			ch == 'o' || ch == 'u' ||

			ch == 'A' || ch == 'E' || ch == 'I' ||
			ch == 'O' || ch == 'U' {
			count++
		}

	}
	return count

}
func main() {

	fmt.Println(CountVowels("Abhishek"))
	fmt.Println(CountVowels("Atulesh"))
}
