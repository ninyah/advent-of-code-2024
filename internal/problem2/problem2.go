package problem2

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ninyah/advent-of-code-2024/internal/utils"
)

func Solve() string {
	levels := getLevelsProblem2()
	safeLevels := 0
	for _, level := range levels {
		unsafeLevels := 0
		levelIncreasing := (level[1] - level[0]) > 0
		lastStep := level[0]
		for stepIndex, step := range level {
			if stepIndex == 0 {
				continue
			}

			stepIncreasing := step-lastStep > 0
			stepDelta := int(math.Abs(float64(step - lastStep)))
			if stepIncreasing != levelIncreasing || stepDelta < 1 || stepDelta > 3 {
				unsafeLevels += 1
			}
			lastStep = step
		}
		if unsafeLevels < 2 {
			safeLevels += 1
		}
	}
	return fmt.Sprintf("%v", safeLevels)
}

func getLevelsProblem2() [][]int {
	rawInput := utils.GetRawInput(2)
	var allLevels [][]int
	for _, line := range strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n") {
		var levels []int
		for _, rawLevel := range strings.Split(line, " ") {
			level, _ := strconv.Atoi(rawLevel)
			levels = append(levels, level)
		}
		allLevels = append(allLevels, levels)
	}
	return allLevels
}
