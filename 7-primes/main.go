package main

import "fmt"

func main() {
	fmt.Println("Prime numbers less than 20")

	for number := 1; number <= 20; number++ {
		if isPrimeNumber(number) {
			fmt.Printf("%v ", number)
		}
	}
}

func isPrimeNumber(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}
