package problem7_test

import (
	"testing"

	"github.com/ninyah/advent-of-code-2024/internal/problem7"
)

func BenchmarkSolveProblem7(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem7.Solve()
	}
}
