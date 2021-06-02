// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercise 1.3
package main

import "testing"

func BenchmarkChapter1example3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Chapter1example3()
	}
}

func BenchmarkChapter1example4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Chapter1example4()
	}
}
