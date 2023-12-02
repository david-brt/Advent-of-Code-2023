package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// possible colors with their corresponding maximum amount
var rules = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	gameIndex := 1
	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		game, err := parseLine(line)
		if err != nil {
			fmt.Println(err)
			return
		}

		possible := true
		for _, cubes := range game {
			for color, amount := range cubes {
				if amount > rules[color] {
					possible = false
				}
			}
		}
		if possible {
			res += gameIndex
		}

		gameIndex++
	}
	fmt.Println(res)
}

func parseLine(line string) ([]map[string]int, error) {
	_, s, _ := strings.Cut(line, ": ")
	cubeSets := strings.Split(s, "; ")
	var game []map[string]int

	for _, cs := range cubeSets {
		m := make(map[string]int)
		cubes := strings.Split(cs, ", ")
		for _, cube := range cubes {
			count, color, _ := strings.Cut(cube, " ")
			var err error
			m[color], err = strconv.Atoi(count)
			if err != nil {
				return make([]map[string]int, 0), err
			}
		}
		game = append(game, m)
	}
	return game, nil
}
