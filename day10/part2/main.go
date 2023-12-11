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

var GRIDSIZE int

var orientations = map[byte][]Coord{
  '.': {{0, 0}, {0, 0}},
  '-': {{-1, 0}, {1, 0}},
  '|': {{0, -1}, {0, 1}},
  'L': {{1, 0}, {0, -1}},
  'J': {{-1, 0}, {0, -1}},
  'S': {{-1, 0}, {0, -1}},
  '7': {{-1, 0}, {0, 1}},
  'F': {{1, 0}, {0, 1}},
}

func main() {
  lines := getLines("example.txt")
  GRIDSIZE = len(lines)
  cycle := getCycle(lines)
  res := enclosedTiles(lines, cycle)
  fmt.Println(res)
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

  var candidates = []Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
  for _, candidate := range candidates {
    x := startPoint.x + candidate.x
    y := startPoint.y + candidate.y
      if outOfBounds(x,y) {
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

func contains(coords []Coord, location Coord) bool {
  for _, coord := range coords {
    if coord == location {
      return true
    }
  }
  return false
}

func outOfBounds(x, y int) bool {
  return x<0 || y<0 || x>=GRIDSIZE || y>=GRIDSIZE
}

func enclosedTiles(pipeGrid [][]byte, cycle []Coord) int {
  res := 0
  for y, line := range pipeGrid {
    var last Coord
    crosses := 0
    for x, pipe := range line {
      if !contains(cycle, Coord{x, y}) {
        if crosses % 2 == 1 {
          res += 1
        }
        continue
      }
      if pipe == '|' {
        crosses += 1
        continue
      }
      if pipe == 'J' || pipe == '7' {
        if last.y - orientations[pipe][1].y != 0 || pipe == '|' {
          crosses += 1
          continue
        }
      }
      if pipe == 'L' || pipe == 'F' {
        last = orientations[pipe][1]
      }
    }
  }
  return res
}
