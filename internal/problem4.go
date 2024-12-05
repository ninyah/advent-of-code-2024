package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveProblem4() string {
	rawInput := getRawInput(4)
	inputLines := strings.Split(rawInput, "\n")
	inputGrid := Map(inputLines, func(line string) []string { return strings.Split(line, "") })

	solution1 := solveProblem4Part1(inputGrid, 140, 140)
	solution2 := solveProblem4Part2(inputGrid, 140, 140)
	return fmt.Sprintf("\nSolution 1: %s\nSolution 2: %s", solution1, solution2)
}

func solveProblem4Part1(inputGrid [][]string, gridWidth int, gridHeight int) string {
	totalMatches := 0

	for y := range gridHeight {
		for x := range gridWidth {
			tryUpRight := y > 2 && x < gridWidth-3
			if tryUpRight {
				upRightLine := inputGrid[y][x] + inputGrid[y-1][x+1] + inputGrid[y-2][x+2] + inputGrid[y-3][x+3]
				if upRightLine == "XMAS" || upRightLine == "SAMX" {
					totalMatches += 1
				}
			}

			tryRight := x < gridWidth-3
			if tryRight {
				rightLine := inputGrid[y][x] + inputGrid[y][x+1] + inputGrid[y][x+2] + inputGrid[y][x+3]
				if rightLine == "XMAS" || rightLine == "SAMX" {
					totalMatches += 1
				}
			}

			tryDownRight := y < gridHeight-3 && x < gridWidth-3
			if tryDownRight {
				downRightLine := inputGrid[y][x] + inputGrid[y+1][x+1] + inputGrid[y+2][x+2] + inputGrid[y+3][x+3]
				if downRightLine == "XMAS" || downRightLine == "SAMX" {
					totalMatches += 1
				}
			}

			tryDown := y < gridHeight-3
			if tryDown {
				downLine := inputGrid[y][x] + inputGrid[y+1][x] + inputGrid[y+2][x] + inputGrid[y+3][x]
				if downLine == "XMAS" || downLine == "SAMX" {
					totalMatches += 1
				}
			}
		}
	}

	return strconv.Itoa(totalMatches)
}

func solveProblem4Part2(inputGrid [][]string, gridWidth int, gridHeight int) string {
	totalMatches := 0

	for y := range gridHeight {
		for x := range gridWidth {
			tryX := y > 0 && x > 0 && y < gridHeight-1 && x < gridWidth-1
			if tryX {
				firstDiagonal := inputGrid[y-1][x-1] + inputGrid[y][x] + inputGrid[y+1][x+1]
				secondDiagonal := inputGrid[y-1][x+1] + inputGrid[y][x] + inputGrid[y+1][x-1]
				if (firstDiagonal == "MAS" || firstDiagonal == "SAM") && (secondDiagonal == "MAS" || secondDiagonal == "SAM") {
					totalMatches += 1
				}
			}
		}
	}

	return strconv.Itoa(totalMatches)
}
