package problem8_test

import (
	"testing"

	"github.com/ninyah/advent-of-code-2024/internal/problem8"
)

func BenchmarkSolveProblem8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem8.Solve()
	}
}
