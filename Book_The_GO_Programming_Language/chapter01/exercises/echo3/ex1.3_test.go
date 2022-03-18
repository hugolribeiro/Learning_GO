package main

import (
	"strings"
	"testing"
)

var args = []string{"/usr/local/bin/go", "foo", "bar", "baz", "zoo", "woo"}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args[1:], " ")
	}
}

func BenchmarkStringsConcatenate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, sep string
		for j := 1; j < len(args); j++ {
			s += sep + args[j]
			sep = " "
		}
	}
}
