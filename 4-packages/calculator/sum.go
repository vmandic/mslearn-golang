package calculator

// If you want something to be private, start its name with a lowercase letter.
// If you want something to be public, start its name with an uppercase letter.

var logMessage = "[LOG]"

var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

func Sum(number1, number2 int) int {
	return number1 + number2
}
