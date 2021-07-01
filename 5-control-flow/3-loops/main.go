package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0
	for i := 1; i <= 3; i++ {
		sum += i
	}

	println(sum)

	// example with not pre / post statements:

	var num int32
	rand.Seed(time.Now().Unix())

	startTime := time.Now()
	for num != 5 {
		num = rand.Int31n(10)
		fmt.Println(num)
	}

	fmt.Println("Elapsed:", time.Since(startTime))

	// infinite for loop:

	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Print("Writting inside the loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num)
	}

	sum = 0

	for num = 1; num <= 100; num++ {
		if num%5 != 0 {
			continue
		}

		sum += int(num)
	}

	println("Sum of divisable 5 numbers up to 100:", sum)
}
