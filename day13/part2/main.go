package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pattern struct {
	rows    [][]byte
	columns [][]byte
}

func main() {
	patterns := parseFile("example.txt")
	res := 0
	for _, pattern := range patterns {
		res += checkSymmetries(pattern)
	}
	fmt.Println(res)
}

func parseFile(filePath string) []Pattern {
	var patterns []Pattern
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var lines [][]byte

	for scanner.Scan() {
		line := []byte(scanner.Text())
		if len(line) == 0 {
			pattern := Pattern{lines, transposeStrings(lines)}
			patterns = append(patterns, pattern)
			lines = make([][]byte, 0)
			continue
		}
		lines = append(lines, line)
	}
	patterns = append(patterns, Pattern{lines, transposeStrings(lines)})
	return patterns
}

func transposeStrings(matrix [][]byte) [][]byte {
	transposeMat := make([][]byte, len(matrix[0]))
	for i := range transposeMat {
		transposeMat[i] = make([]byte, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			transposeMat[j][i] = matrix[i][j]
		}
	}
	return transposeMat
}

func checkSymmetries(pattern Pattern) int {
	horizontal := getSymmetry(pattern.rows)
	if horizontal != -1 {
		return (horizontal + 1) * 100
	}
	vertical := getSymmetry(pattern.columns)
	return vertical + 1
}

func getSymmetry(pattern [][]byte) int {
	max := len(pattern) - 1
	for i := range pattern[:max] {
		different := 0
		for j := 0; j < min(i+1, len(pattern)-i-1); j++ {
			different += compare(pattern[i-j], pattern[i+j+1])
			if different > 1 {
				break
			}
		}
		if different == 1 {
			return i
		}
	}
	return -1
}

// returns how many elements of the given arrays are different
// returns -1 if they are of different lengths
func compare(a []byte, b []byte) int {
	if len(a) != len(b) {
		return -1
	}
	different := 0
	for i := range a {
		if a[i] != b[i] {
			different += 1
		}
	}
	return different
}
