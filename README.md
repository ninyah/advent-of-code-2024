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

- Create your `problemX.go` and `problemX_test.go` files
- Add into your `problemX.go`

```
func SolveProblemX() string {
    return "Answer"
}
```

- Add your `SolveProblemX` function call to `main.go` at the end of the problems array
- Optionally add into a `problemX_test.go`

```
func BenchmarkSolveProblemX(b *testing.B) {
    for n := 0; n < b.N; n++ {
		internal.SolveProblemX()
	}
}
```

# Running Benchmarks

To benchmark all functions, cd into the internal folder, and run `go test -bench=.`
Or specify a specific file with `go test -bench=Problem1`

**Important: All function problems must return their answers in a string**
