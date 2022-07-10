package main

import (
	"fmt"
	"os"
	"strconv"

	cashmachine "code-katas/cash-machine"
)

func main() {
	num, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil || num < 0 {
		fmt.Println("Input provided is not a valid float64 type")
		os.Exit(42)
	}

	cm := cashmachine.New(cashmachine.InitializeGPB())
	value, err := cm.BreakIntoChange(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
