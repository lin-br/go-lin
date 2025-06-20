package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int

	for _, numbersOfSlice := range numbers {
		sums = append(sums, Sum(numbersOfSlice))
	}

	return sums
}
