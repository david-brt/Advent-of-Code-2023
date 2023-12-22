package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type State struct {
	x  int
	y  int
	dx int
	dy int
}

func main() {
	layout := getLines("example.txt")
	// this took about 1min on my machine
	fmt.Println(maxHeat(layout))
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

var energized []State

func shootBeam(layout [][]byte, state State) []State {
	if outOfBounds(layout, state) || contains(energized, state) {
		return energized
	}
	energized = append(energized, state)
	switch layout[state.y][state.x] {
	case '.':
		state = newState(state)
		return shootBeam(layout, state)
	case '\\':
		dx := state.dy
		dy := state.dx
		state = newState(State{state.x, state.y, dx, dy})
		return shootBeam(layout, state)
	case '/':
		dx := -state.dy
		dy := -state.dx
		state = newState(State{state.x, state.y, dx, dy})
		return shootBeam(layout, state)
	case '-':
		if state.dy == 0 {
			state = newState(state)
			return shootBeam(layout, state)
		}
		left := newState(State{state.x, state.y, -1, 0})
		right := newState(State{state.x, state.y, 1, 0})
		return append(shootBeam(layout, left), shootBeam(layout, right)...)
	case '|':
		if state.dx == 0 {
			state = newState(state)
			return shootBeam(layout, state)
		}
		top := newState(State{state.x, state.y, 0, -1})
		bottom := newState(State{state.x, state.y, 0, 1})
		return append(shootBeam(layout, top), shootBeam(layout, bottom)...)
	}
	return energized
}

func outOfBounds(layout [][]byte, state State) bool {
	if state.x < 0 || state.y < 0 {
		return true
	}
	if state.x >= len(layout[0]) || state.y >= len(layout) {
		return true
	}
	return false
}

func newState(s State) State {
	return State{s.x + s.dx, s.y + s.dy, s.dx, s.dy}
}

func contains(arr []State, s State) bool {
	for _, val := range arr {
		if val == s {
			return true
		}
	}
	return false
}

func heat(states []State) int {
	visitedMap := make(map[string]bool)
	for _, s := range states {
		coordString := fmt.Sprint(s.x, s.y)
		if _, visited := visitedMap[coordString]; visited {
			continue
		}
		visitedMap[coordString] = true
	}
	return len(visitedMap)
}

func maxHeat(layout [][]byte) int {
	res := 0

	for x := 0; x < len(layout[0]); x++ {
		energized = make([]State, 0)
		fromTop := heat(shootBeam(layout, State{x, 0, 0, 1}))
		energized = make([]State, 0)
		fromBottom := heat(shootBeam(layout, State{x, len(layout) - 1, 0, -1}))
		res = max(res, fromTop, fromBottom)
	}
	for y := 0; y < len(layout); y++ {
		energized = make([]State, 0)
		fromLeft := heat(shootBeam(layout, State{0, y, 1, 0}))
		energized = make([]State, 0)
		fromRight := heat(shootBeam(layout, State{len(layout[y]) - 1, y, -1, 0}))
		res = max(res, fromLeft, fromRight)
	}
	return res
}
