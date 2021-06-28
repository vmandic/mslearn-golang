package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	firstName, lastName := "John", "Doe"
	age := 32

	const (
		StatusOK = iota
		StatusConnectionReset
		StatusOtherError
	)

	fmt.Println(firstName, lastName, age)
	fmt.Println(StatusOK, StatusConnectionReset, StatusOtherError)

	// NOTE: working with ints:
	var integer32 int32 = 214748364

	fmt.Println(integer32)
	fmt.Printf("val: %d, type: %T\n", integer32, integer32)

	// NOTE: working with uints:
	var unsignedint32 uint32 = 4_000_000_000
	fmt.Println(unsignedint32)

	fmt.Println("max float32:", math.MaxFloat32)
	fmt.Println("max float64:", math.MaxFloat64)

	// var str string = "test"
	// str2 := "ab "

	// println(str)

	i, _ := strconv.Atoi("-42") // _ is a variable discard, strconv.Atoi returns a pair tuple of values
	s := strconv.Itoa(-42)
	println(i, s)
}
