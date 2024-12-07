package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Reference: https://stackoverflow.com/questions/33726731/short-way-to-apply-a-function-to-all-elements-in-a-list-in-golang
func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

// Reference: https://stackoverflow.com/questions/37562873/most-idiomatic-way-to-select-elements-from-an-array-in-golang
func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Copy2dArray[T any](ts [][]T) [][]T {
	tsCopy := make([][]T, len(ts))
	for i := range ts {
		tsCopy[i] = make([]T, len(ts[i]))
		copy(tsCopy[i], ts[i])
	}

	return tsCopy
}

func GetRawInput(problemNumber int) string {
	relativeInput, _ := filepath.Glob("input.txt")
	directoryInput, _ := filepath.Glob(fmt.Sprintf("internal/problem%v/input.txt", problemNumber))
	files := append(relativeInput, directoryInput...)

	file, err := os.Open(files[0])
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return ""
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Printf("File error: %v\n", err)
		}
	}()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		return ""
	}

	return string(bytes)
}
