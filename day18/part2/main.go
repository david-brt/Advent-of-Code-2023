package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Move struct {
	direction Coord
	steps     int
}

var directions = map[byte]Coord{'3': {0, -1}, '0': {1, 0}, '1': {0, 1}, '2': {-1, 0}}

func parseFile(filePath string) []Move {
	var moves []Move

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		color := strings.Trim(strings.Split(line, " ")[2], "()#")
		direction := directions[color[5]]
		steps64, _ := strconv.ParseInt(color[:5], 16, 64)
		steps := int(steps64)
		move := Move{direction, steps}
		moves = append(moves, move)
	}
	return moves
}

func area(moves []Move) int {
	res := 0
	curr := Coord{0, 0}
	for _, move := range moves {
		nextX := curr.x + move.direction.x*move.steps
		nextY := curr.y + move.direction.y*move.steps
		segArea := curr.x*nextY - curr.y*nextX + move.steps
		curr = Coord{nextX, nextY}
		res += segArea
	}
	return res/2 + 1
}

func main() {
	moves := parseFile("example.txt")
	fmt.Println(area(moves))
}
