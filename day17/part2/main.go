package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

func main() {
	heatGrid := parseFile("example.txt")
	res := dijkstra(heatGrid)
	fmt.Println(res)
}

type Coord struct {
	x, y int
}

type State struct {
	dir, loc    Coord
	consecutive int
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
	sourceRight := State{Coord{1, 0}, Coord{0, 0}, 4}
	sourceDown := State{Coord{0, 1}, Coord{0, 0}, 4}
	target := State{Coord{0, 0}, Coord{len(heatGrid[0]) - 1, len(heatGrid) - 1}, 0}
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

		if current.state.loc.x == target.loc.x && current.state.loc.y == target.loc.y {
			if current.state.consecutive < 4 {
				continue
			}
			fmt.Println(current)
			return current.cost
		}
		for _, el := range getNeighbors(heatGrid, &current.state) {
			neighbor := el
			newCost := current.cost + heatGrid[neighbor.loc.y][neighbor.loc.x]
			if neighborCost, exists := cost[neighbor]; exists && neighborCost <= newCost {
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
	directions := []Coord{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	for _, dir := range directions {
		next := Coord{current.loc.x + dir.x, current.loc.y + dir.y}
		if outOfBounds(grid, next.x, next.y) || nextIsPrevious(current, next.x, next.y) {
			continue
		}
		if dir == current.dir {
			if current.consecutive < 10 {
				neighbors = append(neighbors, State{dir, next, current.consecutive + 1})
			}
		} else {
			if current.consecutive >= 4 {
				neighbors = append(neighbors, State{dir, next, 1})
			}
		}
	}
	return neighbors
}
