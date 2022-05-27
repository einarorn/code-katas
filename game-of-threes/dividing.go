package game_of_threes

func DivideUntilOne(number int) []int {
	var output []int
	for {
		output = append(output, number)

		if number == 1 {
			break
		}

		remainder := number % 3

		switch remainder {
		case 0:
			number = number / 3
			break
		case 1:
			number--
			break
		case 2:
			number++
		}
	}
	return output
}
