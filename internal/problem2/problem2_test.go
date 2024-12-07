package problem2_test

import (
	"testing"

	"github.com/ninyah/advent-of-code-2024/internal/problem2"
)

func BenchmarkSolveProblem2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem2.Solve()
	}
}
