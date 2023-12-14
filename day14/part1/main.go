package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := getLines("example.txt")
	transposed := transpose(lines)
	res := 0
	for _, line := range transposed {
		res += rollRocks(line)
	}
	fmt.Println(res)
}

func getLines(filePath string) [][]byte {
	var lines [][]byte

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		lines = append(lines, line)
	}
	return lines
}

func transpose(matrix [][]byte) [][]byte {
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

func rollRocks(line []byte) int {
	stationary := -1
	rocksAtStationary := 0
	var rolled []int
	for i, rock := range line {
		if rock == '#' {
			stationary = i
			rocksAtStationary = 0
		}
		if rock == 'O' {
			rocksAtStationary += 1
			rolled = append(rolled, stationary+rocksAtStationary)
		}
	}
	res := 0
	for _, rock := range rolled {
		res += len(line) - rock
	}
	return res
}
