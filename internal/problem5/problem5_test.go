package problem5_test

import (
	"testing"

	"github.com/ninyah/advent-of-code-2024/internal/problem5"
)

func BenchmarkSolveProblem5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem5.Solve()
	}
}
