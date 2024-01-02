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
	color     string
}

type Tile struct {
	onBorder    bool
	orientation [2]Coord
	color       string
}

var directions = map[string]Coord{"U": {0, -1}, "R": {1, 0}, "D": {0, 1}, "L": {-1, 0}}

func parseFile(filePath string) ([][]Tile, int) {
	var moves []Move

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		ss := strings.Split(line, " ")
		steps, _ := strconv.Atoi(ss[1])
		direction := directions[ss[0]]
		move := Move{direction, steps, ss[2]}
		moves = append(moves, move)
	}

	width, height, start := getDimensions(moves)
	fmt.Println(start)
	tileGrid := make([][]Tile, height)
	for i := range tileGrid {
		tileGrid[i] = make([]Tile, width)
		for j := range tileGrid[i] {
			tileGrid[i][j] = Tile{onBorder: false}
		}
	}

	borderTiles := 0
	dir := Coord{0, 0}
	loc := start
	for _, move := range moves {
		orientation := [2]Coord{{-dir.x, -dir.y}, move.direction}
		dir = move.direction
		tileGrid[loc.y][loc.x] = Tile{true, orientation, move.color}
		borderTiles += 1
		loc.x += dir.x
		loc.y += dir.y
		for i := 1; i < move.steps; i++ {
			orientation = [2]Coord{{-dir.x, -dir.y}, dir}
			tileGrid[loc.y][loc.x] = Tile{true, orientation, move.color}
			borderTiles += 1
			loc.x += dir.x
			loc.y += dir.y
		}
	}
	tileGrid[start.y][start.x].orientation[0] = Coord{-dir.x, -dir.y}
	fmt.Println(tileGrid[start.y][start.x].orientation)

	return tileGrid, borderTiles
}

func getDimensions(moves []Move) (int, int, Coord) {
	maximum := Coord{0, 0}
	minimum := Coord{0, 0}
	current := Coord{0, 0}
	for _, move := range moves {
		for i := 0; i < move.steps; i++ {
			current.x += move.direction.x
			current.y += move.direction.y
			maximum.x = max(maximum.x, current.x)
			maximum.y = max(maximum.y, current.y)
			minimum.x = min(minimum.x, current.x)
			minimum.y = min(minimum.y, current.y)
		}
	}
	width := maximum.x - minimum.x + 1
	height := maximum.y - minimum.y + 1
	start := Coord{-minimum.x, -minimum.y}
	return width, height, start
}

func enclosedTiles(grid [][]Tile) int {
	res := 0
	for _, line := range grid {
		var last Tile
		crosses := 0
		for _, tile := range line {
			if !tile.onBorder {
				if crosses%2 == 1 {
					fmt.Print("#")
					res += 1
					continue
				}
				fmt.Print(".")
				continue
			}
			fmt.Print("#")
			vertical := [2]Coord{{0, -1}, {0, 1}}
			if equivalent(tile, vertical) {
				crosses += 1
				continue
			}
			bottomRight := [2]Coord{{0, -1}, {-1, 0}}
			topRight := [2]Coord{{0, 1}, {-1, 0}}
			topLeft := [2]Coord{{1, 0}, {0, 1}}
			bottomLeft := [2]Coord{{1, 0}, {0, -1}}
			if equivalent(tile, bottomRight) && equivalent(last, topLeft) {
				crosses += 1
				continue
			}
			if equivalent(tile, topRight) && equivalent(last, bottomLeft) {
				crosses += 1
				continue
			}
			if equivalent(tile, topLeft) || equivalent(tile, bottomLeft) {
				last = tile
			}
		}
		fmt.Println()
	}
	return res
}

func equivalent(tile Tile, orientation [2]Coord) bool {
	eq1 := orientation[0] == tile.orientation[0] && orientation[1] == tile.orientation[1]
	eq2 := orientation[0] == tile.orientation[1] && orientation[1] == tile.orientation[0]
	return eq1 || eq2
}

func main() {
	lines, borderTiles := parseFile("example.txt")
	fmt.Println(enclosedTiles(lines) + borderTiles)
}
