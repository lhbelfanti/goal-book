package popcount

import (
	"testing"
)

// -- Benchmarks --

func BenchmarkPopCountSingleExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountSingleExpression(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

// goos: darwin
// goarch: amd64
// pkg: gitlab.com/lhbelfanti/goal-book/chapter2/exercises/exercise2.5
// BenchmarkPopCountSingleExpression
// BenchmarkPopCountSingleExpression-4   	1000000000	         0.317 ns/op

// goos: darwin
// goarch: amd64
// pkg: gitlab.com/lhbelfanti/goal-book/chapter2/exercises/exercise2.5
// BenchmarkPopCountLoop
// BenchmarkPopCountLoop-4   	68945328	        16.4 ns/op

// goos: darwin
// goarch: amd64
// pkg: gitlab.com/lhbelfanti/goal-book/chapter2/exercises/exercise2.5
// BenchmarkPopCountByShifting
// BenchmarkPopCountByShifting-4   	24480542	        48.3 ns/op

// goos: darwin
// goarch: amd64
// pkg: gitlab.com/lhbelfanti/goal-book/chapter2/exercises/exercise2.5
// BenchmarkPopCountByClearing
// BenchmarkPopCountByClearing-4   	73308847	        16.3 ns/op
