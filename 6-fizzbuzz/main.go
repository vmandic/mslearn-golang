package main

import "fmt"

func main() {
	ifElseFizzBuzz()
	fmt.Println("---------------")
	switchFizzBuzz()
}

func ifElseFizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
			if i%5 == 0 {
				fmt.Print("Buzz")
			}
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}

		fmt.Print("\n")
	}
}

func switchFizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
