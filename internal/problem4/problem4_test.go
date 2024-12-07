package problem4_test

import (
	"testing"

	"github.com/ninyah/advent-of-code-2024/internal/problem4"
)

func BenchmarkSolveProblem4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem4.Solve()
	}
}
