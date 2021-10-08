package matrix_test

import (
	"info/go/matrix"
	"testing"
)

var result int

func BenchmarkColumnTraverse(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		// always record the result of ColumnTraverse to prevent
		// the compiler eliminating the function call.
		x = matrix.ColumnTraverse()
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = x
}

func BenchmarkRowTraverse(b *testing.B) {
	var x int
	for i := 0; i < b.N; i++ {
		x = matrix.RowTraverse()
	}
	result = x
}

func BenchmarkColumnTraverseOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// by ignoring the result of ColumnTraverse,
		// this benchmark becomes useless
		matrix.ColumnTraverse()
	}
}

func BenchmarkRowTraverseOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix.RowTraverse()
	}
}
