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

func SumAllTails(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {

			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums
}
