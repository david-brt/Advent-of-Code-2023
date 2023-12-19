package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var stationary, stationaryTransposed, rolling, rollingTransposed [][]int

func main() {
	lines := getLines("example.txt")
	load := cycleN(lines, 1_000_000_000)
	fmt.Println(load)
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

func getRocks(matrix [][]byte) ([][]int, [][]int) {
	stationary := make([][]int, len(matrix))
	rolling := make([][]int, len(matrix))
	for i, rocks := range matrix {
		stationary[i] = append(stationary[i], -1)
		for j, rock := range rocks {
			if rock == '#' {
				stationary[i] = append(stationary[i], j)
			}
			if rock == 'O' {
				rolling[i] = append(rolling[i], j)
			}
		}
	}
	return stationary, rolling
}

func cycleN(lines [][]byte, n int) int {
	transposed := transpose(lines)
	stationary, rolling = getRocks(lines)
	stationaryTransposed, next := getRocks(transposed)

	var states [][][]int
	var loads []int
	var load int

	for i := 0; i < n; i++ {
		next = roll(next, stationaryTransposed, len(lines))
		next = roll(next, stationary, len(lines[0]))
		next = rollBackwards(next, stationaryTransposed, len(lines))
		load = getLoad(next)
		next = rollBackwards(next, stationary, len(lines[0]))

		loopStart := contains(loads, load)
		if loopStart != -1 {
			if compare(states[loopStart], next) {
				loopLength := i - loopStart
				i = loopStart + loopLength*((n-loopStart)/loopLength)
			}
		}
		states = append(states, next)
		loads = append(loads, load)
	}
	return load
}

func roll(rolling, stationary [][]int, height int) [][]int {
	next := make([][]int, height)
	for i := 0; i < len(rolling); i++ {
		rocksAtStationary := 0
		lastStationary := len(stationary[i]) - 1
		for j := len(rolling[i]) - 1; j >= 0; j-- {
			for rolling[i][j] < stationary[i][lastStationary] {
				lastStationary -= 1
				rocksAtStationary = 0
			}
			rocksAtStationary += 1
			rockIndex := stationary[i][lastStationary] + rocksAtStationary
			next[rockIndex] = append(next[rockIndex], i)
		}
	}
	return next
}

func rollBackwards(rolling, stationary [][]int, width int) [][]int {
	next := make([][]int, width)
	for i := 0; i < len(rolling); i++ {
		rocksAtStationary := 0
		stationary[i] = append(stationary[i], width)
		lastStationary := 0
		for j := 0; j < len(rolling[i]); j++ {
			for rolling[i][j] > stationary[i][lastStationary] {
				lastStationary += 1
				rocksAtStationary = 0
			}
			rocksAtStationary += 1
			rockIndex := stationary[i][lastStationary] - rocksAtStationary
			next[rockIndex] = append(next[rockIndex], i)
		}
	}
	return next
}

func getLoad(rocks [][]int) int {
	res := 0
	for i, row := range rocks {
		res += len(row) * (len(rocks) - i)
	}
	return res
}

func compare(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func contains(slice []int, num int) int {
	for i, val := range slice {
		if val == num {
			return i
		}
	}
	return -1
}
