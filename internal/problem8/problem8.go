package problem8

import (
	"fmt"
	"strings"

	"github.com/ninyah/advent-of-code-2024/internal/utils"
)

type Node struct {
	X      int
	Y      int
	Letter string
}

func Solve() string {
	rawInput := utils.GetRawInput(8)
	inputLines := strings.Split(rawInput, "\n")
	inputLines = utils.Map(inputLines, func(line string) string { return strings.TrimSpace(line) })
	inputGrid := utils.Map(inputLines, func(line string) []string { return strings.Split(line, "") })

	gridWidth, gridHeight := len(inputGrid[0]), len(inputGrid)

	var nodes []Node
	for y, line := range inputGrid {
		for x, letter := range line {
			if letter != "." {
				nodes = append(nodes, Node{
					X:      x,
					Y:      y,
					Letter: letter,
				})
			}
		}
	}

	var antiNodes []Node
	for firstIndex, firstNode := range nodes {
		for secondIndex, secondNode := range nodes {
			if secondIndex == firstIndex {
				continue
			}

			deltaX, deltaY := secondNode.X-firstNode.X, secondNode.Y-firstNode.Y
			if firstNode.Letter == secondNode.Letter {
				newPosX, newPosY := secondNode.X, secondNode.Y
				for {
					if newPosX >= 0 && newPosX < gridWidth &&
						newPosY >= 0 && newPosY < gridHeight {
						if !nodeExistsAtPosition(antiNodes, newPosX, newPosY) {
							antiNodes = append(antiNodes, Node{
								X:      newPosX,
								Y:      newPosY,
								Letter: "#",
							})
						}
					} else {
						break
					}
					newPosX, newPosY = newPosX+deltaX, newPosY+deltaY
				}
			}

		}
	}

	return fmt.Sprintf("\nSolution 1: %v", len(antiNodes))
}

func nodeExistsAtPosition(nodes []Node, x int, y int) bool {
	for _, node := range nodes {
		if node.X == x && node.Y == y {
			return true
		}
	}

	return false
}
