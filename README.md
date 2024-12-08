# Advent of Code

Advent of Code is an Advent calendar of small programming puzzles for a variety of skill levels that can be solved in any programming language you like. People use them as interview prep, company training, university coursework, practice problems, a speed contest, or to challenge each other.

**Further Information: https://adventofcode.com/2024/about**

# Usage

```
Usage of advent-of-code-2024.exe:
  -p int
        The problem number to be run <from 1-31> (defaults to the latest problem)
```

So simply run `go run .` to run the latest problem, or specify with an optional `go run . -p 1`

# How to Add a Problem

- Create your `problemX` directory
- Create your `problemX/problemX.go` and optional `problemX/problemX_test.go` files
- Add into your `problemX/problemX.go`

```
func Solve() string {
    return "Answer"
}
```

- Add your `problemX.Solve` function call to `main.go` at the end of the problems array
- Optionally add into a `problemX/problemX_test.go`

```
func BenchmarkSolveProblemX(b *testing.B) {
  for n := 0; n < b.N; n++ {
		problemX.Solve()
	}
}
```

# Running Benchmarks

To benchmark a function, cd into the `internal/problemX` folder, and run `go test -bench=.`

**Important: All function problems must return their answers in a string**
