package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)

	switch i {
	case 0:
		fmt.Println("zero...")
	case 1:
		fmt.Println("one...")
	case 2:
		fmt.Println("two...")
	default:
		fmt.Println("No match!", i)
	}

	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Time to learn go!")
	default:
		fmt.Println("It's weekend, time to rest!")
	}

	fmt.Println("OK")
}
