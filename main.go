package main

import (
	"flag"
	"fmt"

	"github.com/ninyah/advent-of-code-2024/internal/problem1"
	"github.com/ninyah/advent-of-code-2024/internal/problem2"
	"github.com/ninyah/advent-of-code-2024/internal/problem3"
	"github.com/ninyah/advent-of-code-2024/internal/problem4"
	"github.com/ninyah/advent-of-code-2024/internal/problem5"
	"github.com/ninyah/advent-of-code-2024/internal/problem6"
	"github.com/ninyah/advent-of-code-2024/internal/problem7"
	"github.com/ninyah/advent-of-code-2024/internal/problem8"
)

func main() {
	// Just add new problem functions to the array, as you complete them
	// Because you can't store functions in arrays, we had to use the any type
	problems := []any{
		nil,
		problem1.Solve,
		problem2.Solve,
		problem3.Solve,
		problem4.Solve,
		problem5.Solve,
		problem6.Solve,
		problem7.Solve,
		problem8.Solve,
	}
	problemCount := len(problems) - 1

	// Generate an optional command line flag to run a specific problem, defaulting to the latest
	problemNumber := flag.Int("p", problemCount, "The problem number to be run")
	flag.Parse()

	// Grab the function from the array and run it, if possible
	if problems[*problemNumber] != nil {
		// Convert our any type back to a function and call it
		answer := problems[*problemNumber].(func() string)()

		fmt.Printf("Problem %v's answer is: %s\n", *problemNumber, answer)
	} else {
		fmt.Printf("Invalid problem number was specified. Enter a valid problem number: <1-%v>.\n", problemCount)
		flag.Usage()
	}
}
