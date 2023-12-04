package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := getLines("../input.txt")
	fmt.Println(powerSum(lines))

}

func getLines(filePath string) [][]byte {
	var lines [][]byte

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		var bytes []byte
		for _, byte := range line {
			bytes = append(bytes, byte)
		}
		lines = append(lines, bytes)
	}
	return lines
}

func powerSum(lines [][]byte) int {
	powerSum := 0
	for i, line := range lines {
		for j, byte := range line {
			if byte != '*' {
				continue
			}
			var adjacentNumbers []int
			// check above
			if i > 0 {
				numbersAbove := checkAdjacent(lines[i-1], j)
				adjacentNumbers = append(adjacentNumbers, numbersAbove...)
			}

			numbersRightLeft := checkAdjacent(lines[i], j)
			adjacentNumbers = append(adjacentNumbers, numbersRightLeft...)

			// check below
			if i < len(lines)-1 {
				numbersBelow := checkAdjacent(lines[i+1], j)
				adjacentNumbers = append(adjacentNumbers, numbersBelow...)
			}
			if len(adjacentNumbers) != 2 {
				continue
			}
			powerSum += adjacentNumbers[0] * adjacentNumbers[1]
		}
	}
	return powerSum
}

func numberStart(line []byte, currentIndex int) int {
	if !isDigit(line[currentIndex]) {
		return currentIndex + 1
	}
	if currentIndex == 0 {
		return currentIndex
	}
	return numberStart(line, currentIndex-1)
}

// takes an index of a digit and returns the number that contains that digit as well as its length
func getNumber(line []byte, index int) (int, int) {
	var num []byte
	index = numberStart(line, index)
	for i := index; i < len(line); i++ {
		if !isDigit(line[i]) {
			break
		}
		num = append(num, line[i])
	}
	res, err := strconv.Atoi(string(num))
	if err != nil {
		log.Fatalln(err)
	}
	return res, len(num)
}

// returns the power of numbers from line[i-1] to line[i+1]
func checkAdjacent(line []byte, i int) []int {
	var res []int
	for k := i - 1; k <= i+1; k++ {
		if k > 0 && k < len(line) && isDigit(line[k]) {
			num, numLength := getNumber(line, k)
			res = append(res, num)
			// lookahead so no number is added twice
			if k < len(line)-1 && isDigit(line[k+1]) {
				k = i + numLength - 1
			}
		}
	}
	return res
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
