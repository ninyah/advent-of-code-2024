package problem6

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ninyah/advent-of-code-2024/internal/utils"
)

type Direction int

const (
	None Direction = iota
	Both
	Up
	Right
	Down
	Left
)

func Solve() string {
	rawInput := utils.GetRawInput(6)

	inputLines := strings.Split(rawInput, "\n")
	inputLines = utils.Map(inputLines, func(line string) string { return strings.TrimSpace(line) })
	inputGrid := utils.Map(inputLines, func(line string) []string { return strings.Split(line, "") })

	gridCopy := utils.Copy2dArray(inputGrid)
	err := trySolve(gridCopy)
	if err != nil {
		fmt.Printf("err: %x", err)
	}

	part1Solution := 0
	for _, line := range gridCopy {
		part1Solution += len(utils.Filter(line, func(str string) bool { return str != "." && str != "#" && str != "O" }))
	}

	part2Solution := 0

	for potentialObstacleY, inputLine := range inputGrid {
		for potentialObstacleX, inputChar := range inputLine {
			if inputChar != "." {
				continue
			}

			gridCopy := utils.Copy2dArray(inputGrid)
			gridCopy[potentialObstacleY][potentialObstacleX] = "O"

			err := trySolve(gridCopy)
			if err != nil {
				part2Solution += 1
			}
		}
	}

	return fmt.Sprintf("\nSolution 1: %v\nSolution 2: %v", part1Solution, part2Solution)
}

func trySolve(grid [][]string) error {
	gridWidth, gridHeight := len(grid[0]), len(grid)
	direction := Up
	positionX, positionY := 0, 0

	for potentialY, line := range grid {
		potentialX := slices.Index(line, "^")
		if potentialX > -1 {
			positionX, positionY = potentialX, potentialY
		}
	}

	for {
		deltaX, deltaY := deltaPosition(direction)
		nextPositionX, nextPositionY := positionX+deltaX, positionY+deltaY

		if nextPositionX < 0 || nextPositionX >= gridWidth || nextPositionY < 0 || nextPositionY >= gridHeight {
			if direction == Right {
				grid[positionY][positionX] = "➡️"
			} else if direction == Left {
				grid[positionY][positionX] = "⬅️"
			} else if direction == Up {
				grid[positionY][positionX] = "⬆️"
			} else if direction == Down {
				grid[positionY][positionX] = "⬇️"
			}
			break
		}
		if grid[nextPositionY][nextPositionX] == "#" || grid[nextPositionY][nextPositionX] == "O" {
			direction = rotateRight(direction)
			continue
		}

		if direction == Right {
			if grid[positionY][positionX] == "➡️" {
				return fmt.Errorf("loop detected")
			}

			grid[positionY][positionX] = "➡️"
		} else if direction == Left {
			if grid[positionY][positionX] == "⬅️" {
				return fmt.Errorf("loop detected")
			}

			grid[positionY][positionX] = "⬅️"
		} else if direction == Up {
			if grid[positionY][positionX] == "⬆️" {
				return fmt.Errorf("loop detected")
			}

			grid[positionY][positionX] = "⬆️"
		} else if direction == Down {
			if grid[positionY][positionX] == "⬇️" {
				return fmt.Errorf("loop detected")
			}

			grid[positionY][positionX] = "⬇️"
		}

		positionX, positionY = nextPositionX, nextPositionY
	}

	return nil
}

func rotateRight(direction Direction) Direction {
	switch direction {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		return None
	}
}

func deltaPosition(direction Direction) (int, int) {
	switch direction {
	case Up:
		return 0, -1
	case Right:
		return 1, 0
	case Down:
		return 0, 1
	case Left:
		return -1, 0
	default:
		return 0, 0
	}
}
