package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	var sums []int
	for i := range lengthOfNumbers {
		sums = append(sums, Sum(numbersToSum[i]))
	}
	return sums
}
