package calculator

// In Go, the variables of the containers are public (they can be accesed by code in any folder) if they start with an uppercase and
//  private (they can only be accesed for code in this folder) if they start with an lowercase.
// Example of private variable:
var logMessage = "[LOG]"

// Version of the calculator. Public variable.
var Version = "1.0"

// Substract 1 from the number passed as input. This is a private function.
func internalSum(number int) int {
	return number - 1
}

// Sum two integer numbers. This is a public function.
func Sum(number1, number2 int) int {
	return number1 + number2
}
