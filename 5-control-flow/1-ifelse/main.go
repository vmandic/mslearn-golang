package main

import "fmt"

func main() {
	x := 27

	if num := getNumber(); x%2 == 0 { // num is available only in the if / else statement blocks
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "is odd")

		if num < 0 {
			fmt.Println(num, "is negative")
		}
	}
}

func getNumber() int {
	return -42
}
