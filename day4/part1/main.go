package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getLines("../input.txt")
	pointSum := 0
	for _, line := range lines {
		winning, guesses := parseLine(line)
		correctGuesses := findMatches(winning, guesses)
		if len(correctGuesses) == 0 {
			continue
		}
		pointSum += int(math.Pow(2, float64(len(correctGuesses)-1)))
	}
	fmt.Println(pointSum)
}

func getLines(filePath string) []string {
	var lines []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

// returns one array that contains the winning numbers and one that contains the guesses
func parseLine(line string) ([]int, []int) {
	_, numbers, _ := strings.Cut(line, ": ")
	winningString, guessesString, _ := strings.Cut(numbers, " | ")

	winningSplit := strings.Fields(winningString)
	winning := stringSliceToIntSlice(winningSplit)

	guessesSplit := strings.Fields(guessesString)
	guesses := stringSliceToIntSlice(guessesSplit)

	return winning, guesses
}

// takes an array of numeric strings and returns an array of ints
func stringSliceToIntSlice(arr []string) []int {
	res := make([]int, len(arr))
	for i, s := range arr {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}
		res[i] = num
	}
	return res
}

// returns a slice of all ints that are in both input slices
func findMatches(slice1 []int, slice2 []int) []int {
	var res []int
	for _, v1 := range slice1 {
		for _, v2 := range slice2 {
			if v1 == v2 {
				res = append(res, v1)
			}
		}
	}
	return res
}
