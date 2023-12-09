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
	data := parseFile("example.txt")
	res := 0
	for _, data := range data {
		res += predict(data)
	}
	fmt.Println(res)
}

func parseFile(filePath string) [][]int {
	var data [][]int

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		history := make([]int, len(fields))
		for i, s := range fields {
			history[i], _ = strconv.Atoi(s)
		}
		data = append(data, history)
	}
	return data
}

func predict(history []int) int {
	var series = [][]int{history}
	allZeros := false

	i := 0
	for !allZeros {
		var differences []int
		differences, allZeros = getDifferences(series[i])
		series = append(series, differences)
		i++
	}
	prediction := 0
	for i := len(series) - 1; i > 0; i-- {
		prediction += series[i-1][len(series[i-1])-1]
		series[i-1] = append(series[i-1], prediction)
	}
	return prediction
}

func getDifferences(series []int) ([]int, bool) {
	res := make([]int, len(series)-1)
	zeros := 0
	for i := 0; i < len(series)-1; i++ {
		diff := series[i+1] - series[i]
		if diff == 0 {
			zeros += 1
		}
		res[i] = diff
	}

	allZeros := false
	if len(res) == zeros {
		allZeros = true
	}

	return res, allZeros
}
