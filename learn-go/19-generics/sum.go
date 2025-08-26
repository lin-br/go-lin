package generics

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers []int) int {
	addFunc := func(a, b int) int { return a + b }
	return Reduce(numbers, addFunc, 0)
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	sumTail := func(arrayA, arrayB []int) []int {
		if len(arrayB) == 0 {
			return append(arrayA, 0)
		} else {
			tail := arrayB[1:]
			return append(arrayA, Sum(tail))
		}
	}
	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
	var result = initialValue
	for _, item := range collection {
		result = f(result, item)
	}
	return result
}
