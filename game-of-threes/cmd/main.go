package main

import (
	game_of_threes "code-katas/game-of-threes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Input provided is not a number")
		os.Exit(42)
	}
	output := game_of_threes.DivideUntilOne(num)
	for _, n := range output {
		fmt.Println(n)
	}
}
