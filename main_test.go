package main

import (
	"testing"

	_ "github.com/matheus0214/book/testes/matchers"
	"github.com/matheus0214/book/testes/search"
)

func BenchmarkSearch(t *testing.B) {
	for i := 0; i < t.N; i++ {
		search.Run()
	}
}
