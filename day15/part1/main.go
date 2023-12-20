package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	initSequence := parseFile("example.txt")
	res := 0
	for _, val := range initSequence {
		res += hashString(val)
	}
	fmt.Println(res)
}

func parseFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}

	res := strings.Split(line, ",")

	return res
}

func hashString(s string) int {
	hash := 0
	for _, c := range s {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}
