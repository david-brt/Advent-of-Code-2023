package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	nodes, path := parseFile("example.txt")
	fmt.Println(stepCount(nodes, path, "AAA", 0))
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
	if node == "ZZZ" {
		return steps
	}
	if direction == 'L' {
		return stepCount(nodes, path, nodes[node][0], steps+1)
	}
	if direction == 'R' {
		return stepCount(nodes, path, nodes[node][1], steps+1)
	}
	return -1
}
