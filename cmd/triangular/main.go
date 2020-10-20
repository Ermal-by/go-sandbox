package main

import (
	"fmt"

	"github.com/Ermal-by/go-sandbox/pkg/increment"
)

func main() {
	num := parseArg()

	if num > 0 {
		fmt.Printf("First %v triangular numbers are:\n", num)
	}

	for i := 1; i <= num; i++ {
		increment.Increment(i)

		fmt.Print(increment.GlobalVar, " ")

		increment.GlobalVar = 0
	}

	fmt.Println()
}
