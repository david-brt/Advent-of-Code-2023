package main

import (
	"bufio"
	"fmt"
	"os"
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
	res := calibrationSum(scanner)

	fmt.Println(res)
}

func calibrationSum(scanner *bufio.Scanner) int {
	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

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
	}

	return res
}
