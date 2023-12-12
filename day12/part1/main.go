package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getLines("../input.txt")
	res := 0
	for _, line := range lines {
		res += countArrangements(line)
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

func countArrangements(line []byte) int {
	springs, brokenString, _ := strings.Cut(string(line), " ")
	broken := strings.Split(brokenString, ",")
	questionMarks := 0
	for _, spring := range []byte(springs) {
		if spring == '?' {
			questionMarks += 1
		}
	}
	arrangements := 0
	var permutations []string
	permutate(questionMarks, "", &permutations)

	for i := 0; i < len(permutations); i++ {
		tmp := make([]byte, len(springs))
		copy(tmp, springs)
		k := 0
		for j := 0; j < len(springs); j++ {
			if springs[j] == '?' {
				tmp[j] = permutations[i][k]
				k++
			}
		}
		permutations[i] = string(tmp)
	}

	for i := 0; i < len(permutations); i++ {
		possiblyBroken := countBroken([]byte(permutations[i]))
		if stringSliceEqualsIntSlice(broken, possiblyBroken) {
			arrangements += 1
		}
	}
	return arrangements
}

func countBroken(slice []byte) []int {
	broken := 0
	var res []int
	for _, b := range slice {
		if b == '#' {
			broken += 1
			continue
		}
		if broken > 0 {
			res = append(res, broken)
		}
		broken = 0
	}
	if broken > 0 {
		res = append(res, broken)
	}
	return res
}

func stringSliceEqualsIntSlice(strings []string, ints []int) bool {
	if len(strings) != len(ints) {
		return false
	}
	for i, a := range ints {
		b, _ := strconv.Atoi(strings[i])
		if a != b {
			return false
		}
	}
	return true
}

func permutate(n int, prefix string, result *[]string) {
	if n == 0 {
		*result = append(*result, prefix)
	} else {
		permutate(n-1, prefix+"#", result)
		permutate(n-1, prefix+".", result)
	}
}
