package main

import (
	"os"
	"strconv"
)

func main() {
	println("Sum 1:", sum(os.Args[1], os.Args[2]))

	sum, mul := calc(os.Args[1], os.Args[2])

	println("Sum 2:", sum)
	println("Mul 2:", mul)

	sum, _ = calc(os.Args[1], os.Args[2])

	println("Sum 3:", sum)

	var someName = "Jack"
	updateName(someName) // NOTE: golang copies by value
	println(someName)

	updateNameByRef(&someName)
	println(someName)
}

func sum(strnum1 string, strnum2 string) (result int) {
	int1, _ := strconv.Atoi(strnum1)
	int2, _ := strconv.Atoi(strnum2)
	result = int1 + int2
	return
}

func calc(strnum1 string, strnum2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(strnum1)
	int2, _ := strconv.Atoi(strnum2)

	sum = int1 + int2
	mul = int1 * int2
	return
}

func updateName(name string) {
	name = "David"
}

func updateNameByRef(name *string) {
	*name = "David"
}
