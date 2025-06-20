package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for index, number := range numbers {
		sums[index] = Sum(number)
	}

	return sums
}
