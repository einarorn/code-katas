package letter_sum

const ASCIIOffset = 96

func CalculateASCIISum(word string) int {
	sum := 0
	for _, c := range word {
		sum += int(c) - ASCIIOffset
	}
	return sum
}
