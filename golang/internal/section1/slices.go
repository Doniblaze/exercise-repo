package section1

func SumAllEvenNumber(numbers []int) int {
	sumOfNumbers := 0
	//a for loop to check if a number in the slice has a remainder or not
	for _, number := range numbers {
		if number%2 == 0 {
			sumOfNumbers += number
		}
	}
	return sumOfNumbers
}
