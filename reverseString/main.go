package main

import "fmt"

func reverse(str string) string {
	reverse := []rune(str)

	for i, a := 0, len(reverse)-1; i < a; i, a = i+1, a-1 {
		reverse[i], reverse[a] = reverse[a], reverse[i]
	}
	return string(reverse)
}
func main() {

	fmt.Println(reverse("Abhishek"))

}
