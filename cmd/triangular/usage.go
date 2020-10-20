package main

import "fmt"

func usage(baseName string) {
	fmt.Printf(
		`
Example usage:
        '%v <num>' to print first <num> triangular numbers
`,
		baseName)
}
