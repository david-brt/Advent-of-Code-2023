package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		lineBytes := parseLine(line)

		res += findFirstDigitLiteral(lineBytes) * 10
		res += findLastDigitLiteral(lineBytes)
	}

	fmt.Println(res)
}

// returns the first digit literal or 0 if none found
func findFirstDigitLiteral(line []byte) int {
	for i := 0; i < len(line); i++ {
		// check if byte is integer between 0 and 9
		if line[i] >= 48 && line[i] < 58 {
			return int(line[i] - '0')
		}
	}
	return 0
}

// returns the last digit literal or 0 if none found
func findLastDigitLiteral(line []byte) int {
	for i := len(line) - 1; i >= 0; i-- {
		// check if byte is integer between 0 and 9
		if line[i] >= 48 && line[i] < 58 {
			return int(line[i] - '0')
		}
	}
	return 0
}

// guarantees that the first and the last number word have their first character converted to a number
func parseLine(line string) []byte {
	lineBytes := []byte(line)
	for _, dm := range digitMappings {
		wordIndex := strings.Index(line, dm.word)
		if wordIndex != -1 {
			lineBytes[wordIndex] = '0' + byte(dm.literal)
		}
	}

	for _, dm := range digitMappings {
		wordIndex := strings.LastIndex(line, dm.word)
		if wordIndex != -1 {
			lineBytes[wordIndex] = '0' + byte(dm.literal)
		}
	}

	return lineBytes
}
