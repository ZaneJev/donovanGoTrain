package main

import "testing"

// task 1.3
var result string

func BenchmarkEchoStringJoin(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = echoStringJoin([]string{"first", "second"})
	}
	result = s
}

func BenchmarkEchoBuffConcat(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = echoBuffConcat([]string{"first", "second"})
	}
	result = s
}

func BenchmarkEchoStringConcat(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = echoStringConcat([]string{"first", "second"})
	}
	result = s
}
