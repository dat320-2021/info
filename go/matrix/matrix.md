# Benchmarking Matrix Traversal

The matrix traversal example is just a dummy traversal of a matrix in two different ways, row-by-row and column-by-column.
The goal of the example is to illustrate the effects that CPU cache lines (64 bytes wide) have on the performance of the traversal.

Dave Cheney's blog on [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
explains a benchmark should be careful to avoid compiler optimizations eliminating the function under test and artificially lowering the run time of the benchmark.

I just got bitten (during lectures) by improved compiler optimizations with the matrix traversal example; previously I had ignored the above advice and used the following benchmark:

```go
func BenchmarkRowTraverseOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
	    matrix.RowTraverse()
	}
}
```

The above benchmark worked fine for many years, but since go1.16, the compiler has become smarter:

```shell
% go1.15 test -bench .
Elements in the matrix 16777216
goos: darwin
goarch: amd64
pkg: info/go/matrix
BenchmarkColumnTraverse-12       	       8	 151121734 ns/op
BenchmarkRowTraverse-12          	     105	  10970847 ns/op
BenchmarkColumnTraverseOld-12    	       8	 152228546 ns/op
BenchmarkRowTraverseOld-12       	     103	  10958098 ns/op
PASS
ok  	info/go/matrix	6.765s
```

```shell
% go1.16 test -bench .
Elements in the matrix 16777216
goos: darwin
goarch: amd64
pkg: info/go/matrix
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkColumnTraverse-12       	       7	 154601012 ns/op
BenchmarkRowTraverse-12          	     100	  11657856 ns/op
BenchmarkColumnTraverseOld-12    	     294	   3991434 ns/op
BenchmarkRowTraverseOld-12       	     296	   4070349 ns/op
PASS
ok  	info/go/matrix	6.671s
```
