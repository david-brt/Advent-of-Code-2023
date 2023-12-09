package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	nodes, path := parseFile("../input.txt")
	var stepCounts []int
	for node := range nodes {
		if strings.LastIndex(node, "A") == 2 {
			stepCounts = append(stepCounts, stepCount(nodes, path, node, 0))
		}
	}
	res := stepCounts[0]
	for _, stepCount := range stepCounts {
		res = lcm(res, stepCount)
	}
	fmt.Println(res)
}

func parseFile(filePath string) (map[string][]string, []byte) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var lines []string
	nodes := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	path := []byte(lines[0])

	for i := 2; i < len(lines); i++ {
		node, next, _ := strings.Cut(lines[i], " = ")
		next = strings.Trim(next, "()")
		left, right, _ := strings.Cut(next, ", ")
		var nextArr = []string{left, right}
		nodes[node] = nextArr
	}
	return nodes, path
}

func stepCount(nodes map[string][]string, path []byte, node string, steps int) int {
	direction := path[steps%len(path)]
	var next string
	if strings.LastIndex(node, "Z") == 2 {
		return steps
	}
	if direction == 'L' {
		next = nodes[node][0]
	}
	if direction == 'R' {
		next = nodes[node][1]
	}
	return stepCount(nodes, path, next, steps+1)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	tmp := a
	a = b
	b = tmp % a
	return gcd(a, b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
