package internal_test

import (
	"testing"

	"github.com/ninyah/advent-of-code-2024/internal"
)

func BenchmarkSolveProblem1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		internal.SolveProblem1()
	}
}