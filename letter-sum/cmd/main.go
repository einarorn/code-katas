package main

import (
	"fmt"
	"os"

	letterSum "code-katas/letter-sum"
)

func main() {
	word := os.Args[1]
	fmt.Println(letterSum.CalculateASCIISum(word))
}
