package main

import (
	"fmt"
	"os"
	"strconv"

	"code-katas/fibonacci"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num < 0 {
		fmt.Println("Input provided is not a positive number")
		os.Exit(42)
	}

	fmt.Println(fibonacci.Fib(uint(num)))
}
