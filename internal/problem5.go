package internal

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolveProblem5() string {
	rawInput := getRawInput(5)
	splitRawInputs := strings.Split(rawInput, "\n\r")

	pageOrderingRulesStr := strings.TrimSpace(splitRawInputs[0])
	pagesToProduceStr := strings.TrimSpace(splitRawInputs[1])

	pageOrderingRules := make(map[string][]string)
	pageOrderingRulesReverse := make(map[string][]string)
	for _, line := range strings.Split(pageOrderingRulesStr, "\n") {
		ruleStr := strings.Split(strings.TrimSpace(line), "|")
		firstPage := ruleStr[0]
		secondPage := ruleStr[1]

		pageOrderingRules[firstPage] = append(pageOrderingRules[firstPage], secondPage)
		pageOrderingRulesReverse[secondPage] = append(pageOrderingRulesReverse[secondPage], firstPage)
	}

	validTotal := 0
	invalidTotal := 0
	for _, line := range strings.Split(pagesToProduceStr, "\n") {
		pagesStr := strings.Split(strings.TrimSpace(line), ",")
		pagesValid := true

		for pageIndex, page := range pagesStr {
			for _, otherPage := range pagesStr[pageIndex+1:] {
				if slices.Contains(pageOrderingRules[otherPage], page) {
					pagesValid = false
				}
			}
		}

		if pagesValid {
			middleIndex := (len(pagesStr) - 1) / 2
			middleNumber, _ := strconv.Atoi(pagesStr[middleIndex])

			validTotal += middleNumber
		} else {
			slices.SortFunc(pagesStr, func(firstPage string, secondPage string) int {
				if slices.Contains(pageOrderingRules[firstPage], secondPage) {
					return -1
				} else if slices.Contains(pageOrderingRulesReverse[secondPage], firstPage) {
					return 1
				} else {
					return 0
				}
			})

			middleIndex := (len(pagesStr) - 1) / 2
			middleNumber, _ := strconv.Atoi(pagesStr[middleIndex])

			invalidTotal += middleNumber
		}
	}

	return fmt.Sprintf("\nSolution 1: %v\nSolution 2: %v", validTotal, invalidTotal)
}
