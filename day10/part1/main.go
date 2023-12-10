package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coord struct {
	x int
	y int
}

var orientations = map[byte][]Coord{
	'|': {{0, -1}, {0, 1}},
	'-': {{-1, 0}, {1, 0}},
	'L': {{0, -1}, {1, 0}},
	'J': {{-1, 0}, {0, -1}},
	'7': {{-1, 0}, {0, 1}},
	'F': {{0, 1}, {1, 0}},
	'.': {{0, 0}, {0, 0}},
}

func main() {
	lines := getLines("example.txt")
	cycle := getCycle(lines)
	fmt.Println(len(cycle) / 2)
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

func getStartPoint(pipeGrid [][]byte) (Coord, Coord) {
	startPoint := Coord{-1, -1}
	for y, pipes := range pipeGrid {
		for x, pipe := range pipes {
			if pipe == 'S' {
				startPoint = Coord{x, y}
			}
		}
	}
	next := Coord{-1, -1}

	for y := startPoint.y - 1; y < startPoint.y+2; y++ {
		for x := startPoint.x - 1; x < startPoint.x+2; x++ {
			if x < 0 || y < 0 || y >= len(pipeGrid) || x >= len(pipeGrid[y]) {
				continue
			}
			location := pipeGrid[y][x]
			orientation := orientations[location]
			for _, coord := range orientation {
				if x+coord.x == startPoint.x && y+coord.y == startPoint.y {
					next = Coord{x, y}
					break
				}
			}
		}
	}

	return startPoint, next
}

func getCycle(pipeGrid [][]byte) []Coord {
	start, loc := getStartPoint(pipeGrid)
	previous := start
	cycle := []Coord{start, loc}
	i := 0
	for loc != start {
		pipe := pipeGrid[loc.y][loc.x]
		direction := orientations[pipe][0]
		if loc.x+direction.x == previous.x && loc.y+direction.y == previous.y {
			direction = orientations[pipe][1]
		}
		previous = loc
		loc.x += direction.x
		loc.y += direction.y
		cycle = append(cycle, loc)
		i++
	}
	return cycle
}
