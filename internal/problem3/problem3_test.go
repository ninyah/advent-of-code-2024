package problem3_test

import (
	"testing"

	"github.com/ninyah/advent-of-code-2024/internal/problem3"
)

func BenchmarkSolveProblem3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem3.Solve()
	}
}
