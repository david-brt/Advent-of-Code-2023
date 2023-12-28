package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

func main() {
	heatGrid := parseFile("../input.txt")
	res := dijkstra(heatGrid)
	fmt.Println(res)
}

type Coord struct {
	x, y int
}

type State struct {
	dir, loc          Coord
	consecutive, cost int
}

type Item struct {
	cost, index int
	state       State
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = 0
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func parseFile(filePath string) [][]int {
	var heatGrid [][]int

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		vals := make([]int, len(line))
		for i, val := range line {
			vals[i] = int(val - '0')
		}
		heatGrid = append(heatGrid, vals)
	}
	return heatGrid
}

func dijkstra(heatGrid [][]int) int {
	sourceRight := State{Coord{1, 0}, Coord{0, 0}, 1, 0}
	sourceDown := State{Coord{0, 1}, Coord{0, 0}, 1, 0}
	target := State{Coord{0, 0}, Coord{len(heatGrid[0]) - 1, len(heatGrid) - 1}, 0, 0}
	pq := PriorityQueue{
		&Item{cost: 0, state: sourceRight, index: 0},
		&Item{cost: 0, state: sourceDown, index: 1},
	}
	heap.Init(&pq)

	cost := make(map[State]int)

	cost[sourceDown] = 0
	cost[sourceRight] = 0

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)
		if cost[current.state] < current.cost {
			continue
		}
		if current.state.loc.x == target.loc.x && current.state.loc.y == target.loc.y {
			return current.cost
		}
		for _, el := range getNeighbors(heatGrid, &current.state) {
			neighbor := el
			newCost := cost[current.state] + neighbor.cost
			if _, exists := cost[neighbor]; exists && newCost >= cost[current.state] {
				continue
			}
			cost[neighbor] = newCost
			heap.Push(&pq, &Item{cost: newCost, state: neighbor})
		}
	}
	return -1
}

func outOfBounds(grid [][]int, x, y int) bool {
	return x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid)
}

func nextIsPrevious(state *State, nextX, nextY int) bool {
	previousX := state.loc.x - state.dir.x
	previousY := state.loc.y - state.dir.y
	return previousX == nextX && previousY == nextY
}

func getNeighbors(grid [][]int, current *State) []State {
	var neighbors []State
	directions := []struct{ x, y int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	for _, dir := range directions {
		nextX, nextY := current.loc.x+dir.x, current.loc.y+dir.y
		if outOfBounds(grid, nextX, nextY) || nextIsPrevious(current, nextX, nextY) {
			continue
		}
		if current.consecutive >= 3 && current.dir == dir {
			continue
		}
		if dir == current.dir {
			neighbors = append(neighbors, State{dir, Coord{nextX, nextY}, current.consecutive + 1, grid[nextY][nextX]})
			continue
		}
		neighbors = append(neighbors, State{dir, Coord{nextX, nextY}, 1, grid[nextY][nextX]})
	}
	return neighbors
}
