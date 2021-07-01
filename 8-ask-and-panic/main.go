package main

import "fmt"

func main() {
	val := 0

	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("You entered a negative number!")
		case val == 0:
			fmt.Println("You entered zero.")
		default:
			fmt.Println("You entered:", val)
		}
	}
}
