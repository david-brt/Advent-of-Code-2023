package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pattern struct {
	rows    []string
	columns []string
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
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			pattern := Pattern{lines, transposeStrings(lines)}
			patterns = append(patterns, pattern)
			lines = make([]string, 0)
			continue
		}
		lines = append(lines, line)
	}
	patterns = append(patterns, Pattern{lines, transposeStrings(lines)})
	return patterns
}

func transposeStrings(matrix []string) []string {
	transposeMat := make([][]byte, len(matrix[0]))
	for i := range transposeMat {
		transposeMat[i] = make([]byte, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			transposeMat[j][i] = matrix[i][j]
		}
	}

	res := make([]string, len(transposeMat))
	for i, row := range transposeMat {
		res[i] = string(row)
	}
	return res
}

func checkSymmetries(pattern Pattern) int {
	horizontal := getSymmetry(pattern.rows)
	if horizontal != -1 {
		return (horizontal + 1) * 100
	}
	vertical := getSymmetry(pattern.columns)
	return vertical + 1
}

func getSymmetry(pattern []string) int {
	max := len(pattern) - 1
	for i := range pattern[:max] {
		symmetric := true
		for j := 0; j < min(i+1, len(pattern)-i-1); j++ {
			if pattern[i-j] != pattern[i+j+1] {
				symmetric = false
				break
			}
		}
		if symmetric {
			return i
		}
	}
	return -1
}
