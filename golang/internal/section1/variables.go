package section1

import "fmt"

var a, b int

func SwapWithoutTemp() {
	fmt.Println("Operation that swaps the values of two variables without using a temporary variable.")
	// Swap values using addition and subtraction
	fmt.Print("Enter Value of First Number (a): ")
	fmt.Scanln(&a)
	fmt.Print("Enter Value of Second Number (b): ")
	fmt.Scanln(&b)

	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)

	// Swap values using addition and subtraction
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("After swap: a = %d, b = %d\n", a, b)
}
