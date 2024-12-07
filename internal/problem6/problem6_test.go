package problem6_test

import (
	"testing"

	"github.com/ninyah/advent-of-code-2024/internal/problem6"
)

func BenchmarkSolveProblem6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem6.Solve()
	}
}
