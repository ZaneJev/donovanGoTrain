// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// !+
func main() {
	// task 1.1
	fmt.Println(strings.Join(os.Args[:], " "))

	// task 1.2
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}

}

func echoStringJoin(args []string) string {
	return strings.Join(args[:], " ")
}

func echoBuffConcat(args []string) string {
	var buffer bytes.Buffer
	for _, s := range args {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func echoStringConcat(args []string) string {
	var str string
	for _, s := range args {
		str += s
	}
	return str
}

//!-
