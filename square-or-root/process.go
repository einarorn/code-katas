package square_or_root

import "math"

func ProcessNumber(num int) int {
	results := math.Sqrt(float64(num))
	if results != math.Floor(results) {
		results = math.Pow(float64(num), 2)
	}

	return int(results)
}

func ProcessArray(numbers []int) []int {
	for i, n := range numbers {
		numbers[i] = ProcessNumber(n)
	}
	return numbers
}
