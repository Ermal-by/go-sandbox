package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"unsafe"
)

func parseArg() (num int) {
	baseName := filepath.Base(os.Args[0])

	if len(os.Args) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "%v: argument <num> is required\n", baseName)
		usage(baseName)

		os.Exit(1)
	}

	parsed, err := strconv.ParseInt(os.Args[1], 0, int(unsafe.Sizeof(num)*8))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v: argument <num> is invalid: %v\n", baseName, err)
		usage(baseName)

		os.Exit(1)
	}

	num = int(parsed)

	return
}
