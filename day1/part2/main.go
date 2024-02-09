package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("../input.txt")
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
		res += calibrationSum(lineBytes)
	}

	fmt.Println(res)
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

func calibrationSum(line []byte) int {
	res := 0

	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			res += int(line[i]-'0') * 10
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			res += int(line[i] - '0')
			break
		}
	}

	return res
}
