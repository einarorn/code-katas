package rotate

func ShiftLeft(input []int, shifting int) []int {
	k := len(input)

	for i:=0;i<shifting;i++ {
		tmp := input[0]
		for i:=0;i<k-1;i++ {
			input[i] = input[i+1]
		}
		input[k-1] = tmp
	}

	return input
}
