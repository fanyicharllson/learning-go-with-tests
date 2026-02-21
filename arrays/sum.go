package arrays

import "fmt"

func Sum(num []int) int {
	sum := 0
	for _, number := range num {
		sum += number
	}
	return sum
}

func SumMulti(numberToSum ...[]int) []int {
	totalNumber := len(numberToSum)
	sum := make([]int, totalNumber)

	fmt.Printf("Slice param passed to the function %v", numberToSum)

	for i, value := range numberToSum {
		sum[i] = Sum(value)
	}

	return sum

}
func OptionalSumMulti(numberToSum ...[]int) []int {
	var sums []int
	fmt.Printf("Slice param passed to the function %v", numberToSum)

	for _, value := range numberToSum {
		sums = append(sums, Sum(value))
	}

	return sums
}

func SumTotalTails(numberToSum ...[]int) []int {
	var sums []int
	fmt.Printf("Slice param passed to the function %v", numberToSum)
	for _, value := range numberToSum {
		tails := value[1:]
		sums = append(sums, Sum(tails))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int { // [][]int{ { }, {3,4,5} }
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
