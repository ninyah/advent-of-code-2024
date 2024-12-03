package internal

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func SolveProblem1() string {
	// TODO
	lists := getListsProblem1()
	listA := lists[0]
	listB := lists[1]
	listBCounts := make(map[int]int)
	for _, number := range listB {
		listBCounts[number] += 1
	}

	sort.Ints(listA)
	sort.Ints(listB)

	totalDistance := 0
	similarityScore := 0

	for index := range listA {
		idA := listA[index]
		idB := listB[index]
		distance := math.Abs(float64(idA - idB))
		totalDistance += int(distance)
		similarityScore += idA * listBCounts[idA]
	}

	return fmt.Sprintf("Total Distance: %v\nSimilarity Score: %v", totalDistance, similarityScore)
}

func getListsProblem1() [][]int {
	var listA []int
	var listB []int
	rawInput := getRawInput(1)
	for _, line := range strings.Split(strings.TrimSuffix(rawInput, "\n"), "\n") {
		ids := strings.Split(line, "   ")
		idA, _ := strconv.Atoi(ids[0])
		idB, _ := strconv.Atoi(ids[1])
		listA = append(listA, idA)
		listB = append(listB, idB)
	}
	return [][]int{listA, listB}
}
