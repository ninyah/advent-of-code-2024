package internal

import (
	"fmt"
	"io"
	"os"
)

func getRawInput(problemNumber int) string {
	fileName := fmt.Sprintf("internal/problem%vInput.txt", problemNumber)
	file, err := os.Open(fileName)
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
