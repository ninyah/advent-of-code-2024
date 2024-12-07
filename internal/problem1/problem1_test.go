package problem1_test

import (
	"testing"

	"github.com/ninyah/advent-of-code-2024/internal/problem1"
)

func BenchmarkSolveProblem1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem1.Solve()
	}
}
