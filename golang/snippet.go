package main

import "fmt"

// Function to calculate the factorial of a number
func factorial(n int) int {
	/*
		The code below contains incorrect logic on getting the factorial of a number,
		 because it is incrementing the result by i.Also the float64 is not been converted to an int.
		 I would recommend this code:
		 	result := 1
		 for i := 1; i <= n; i++ {
			result *= i
		}
		return result
	*/
	result := 1.00
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

// Function to print the factorial of a number
func printFactorial() {
	num := 5
	fmt.Printf("The factorial of %d is: %d\n", num, factorial(num))
}

func main() {
	//Go is case sensitive, therefore the calling of this function is inappropriate
	/*i would recommend using this:
	printFactorial()
	*/
	printfactorial()
}
