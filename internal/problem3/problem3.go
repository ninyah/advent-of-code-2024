package problem3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ninyah/advent-of-code-2024/internal/utils"
)

func Solve() string {
	rawInput := utils.GetRawInput(3)
	mulRegex := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	mulExpressions := mulRegex.FindAllString(rawInput, -1)
	mulExpressionIndices := mulRegex.FindAllStringIndex(rawInput, -1)

	total := 0
	enabledTotal := 0

	for expressionIndex, expression := range mulExpressions {
		expressionString, _ := strings.CutPrefix(expression, "mul(")
		expressionString, _ = strings.CutSuffix(expressionString, ")")
		numbersStrings := strings.Split(expressionString, ",")
		number1, _ := strconv.Atoi(numbersStrings[0])
		number2, _ := strconv.Atoi(numbersStrings[1])
		mulExpressionIndex := mulExpressionIndices[expressionIndex][1]

		total += number1 * number2
		if mulIsEnabled(rawInput, mulExpressionIndex) {
			enabledTotal += number1 * number2
		}
	}

	return fmt.Sprintf("\nPart 1: %v\nPart 2: %v", total, enabledTotal)
}

func mulIsEnabled(input string, index int) bool {
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	doIndices := doRegex.FindAllStringIndex(input, -1)
	dontIndices := dontRegex.FindAllStringIndex(input, -1)

	latestDoIndex := 0
	for _, doIndex := range doIndices {
		if doIndex[1] < index {
			latestDoIndex = doIndex[0]
		} else {
			break
		}
	}

	latestDontIndex := 0
	for _, dontIndex := range dontIndices {
		if dontIndex[1] < index {
			latestDontIndex = dontIndex[0]
		} else {
			break
		}
	}

	return latestDoIndex >= latestDontIndex
}
