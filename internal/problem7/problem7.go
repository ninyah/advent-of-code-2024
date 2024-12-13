package problem7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ninyah/advent-of-code-2024/internal/utils"
)

func Solve() string {
	rawInput := utils.GetRawInput(7)
	lines := utils.Map(strings.Split(rawInput, "\n"), func(line string) string { return strings.TrimSpace(line) })

	solution1, solution2 := 0, 0

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		firstNumber, _ := strconv.Atoi(strings.TrimSuffix(numbers[0], ":"))
		restOfNumbers := utils.Map(numbers[1:], func(number string) int {
			num, _ := strconv.Atoi(number)
			return num
		})

		listOfSumsAndMults := getListOfSumsAndMults(restOfNumbers)
		listOfSumsAndMultsAndConcats := getListOfSumsAndMultsAndConcats(restOfNumbers)
		if slices.Contains(listOfSumsAndMults, firstNumber) {
			solution1 += firstNumber
		}
		if slices.Contains(listOfSumsAndMultsAndConcats, firstNumber) {
			solution2 += firstNumber
		}
	}

	return fmt.Sprintf("\nSolution 1: %v\nSolution 2: %v", solution1, solution2)
}

func getListOfSumsAndMults(nums []int) []int {
	var resultList []int

	if len(nums) == 2 {
		resultList = append(resultList, nums[0]+nums[1])
		resultList = append(resultList, nums[0]*nums[1])
	} else if len(nums) > 2 {
		lastIndex := len(nums) - 1
		subTotals := getListOfSumsAndMults(nums[:lastIndex])

		for _, subTotal := range subTotals {
			resultList = append(resultList, nums[lastIndex]+subTotal)
			resultList = append(resultList, nums[lastIndex]*subTotal)
		}
	}

	return resultList
}

func getListOfSumsAndMultsAndConcats(nums []int) []int {
	var resultList []int

	if len(nums) == 2 {
		resultList = append(resultList, nums[0]+nums[1])
		resultList = append(resultList, nums[0]*nums[1])
		resultList = append(resultList, concatNums(nums[0], nums[1]))
	} else if len(nums) > 2 {
		lastIndex := len(nums) - 1
		subTotals := getListOfSumsAndMultsAndConcats(nums[:lastIndex])

		for _, subTotal := range subTotals {
			resultList = append(resultList, nums[lastIndex]+subTotal)
			resultList = append(resultList, nums[lastIndex]*subTotal)
			resultList = append(resultList, concatNums(subTotal, nums[lastIndex]))
		}
	}

	return resultList
}

func concatNums(a int, b int) int {
	strA, strB := strconv.Itoa(a), strconv.Itoa(b)
	strResult := strA + strB
	result, _ := strconv.Atoi(strResult)
	return result
}
