package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		km := metresToKilometres(t)
		m := kilometresToMetres(t)

		fmt.Println(km, m)
	}
}

// metresToKilometres - convert metres to kilometres
func metresToKilometres(metres int64) int64 {
	return metres / 1000
}

// kilometresToMetres - convert kilometres to metres
func kilometresToMetres(kilometres int64) int64 {
	return kilometres * 1000
}

//!-
