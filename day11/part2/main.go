package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Coord struct {
  x int
  y int
}

func main() {
  universe := getLines("example.txt")
  galaxies := findGalaxies(universe)
  res := galaxyDistances(galaxies, emptyRows(universe), emptyColumns(universe))
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

func emptyRows(universe [][]byte) []int {
  var emptyRows []int
  for i:=0; i<len(universe); i++ {
    if slices.Contains(universe[i], '#') {
      continue
    }
    emptyRows = append(emptyRows, i)
  }
  return emptyRows
}

func emptyColumns(universe [][]byte) []int {
  var emptyColumns []int
  for i:=0; i<len(universe[0]); i++ {
    empty := true
    for j:=0; j<len(universe); j++ {
      if universe[j][i] == '#'{
        empty = false
      }
    }
    if !empty {
      continue
    }
    emptyColumns = append(emptyColumns, i)
  }
  return emptyColumns
}

func findGalaxies(universe [][]byte) []Coord {
  var galaxies []Coord
  for y, l := range universe {
    for x, v := range l {
      if v == '#' {
        galaxies = append(galaxies, Coord{x, y})
      }
    }
  }
  return galaxies
}

func galaxyDistances(galaxies []Coord, emptyRows, emptyColumns []int) int {
  distanceSum := 0
  for i, galaxy := range galaxies {
    for _, nextGalaxy := range galaxies[i+1:] {
        x1, x2 := min(galaxy.x, nextGalaxy.x), max(galaxy.x, nextGalaxy.x)
        y1, y2 := min(galaxy.y, nextGalaxy.y), max(galaxy.y, nextGalaxy.y)
        distanceX, distanceY := x2-x1, y2-y1

        for _, columnIndex := range emptyColumns {
            if x1 < columnIndex && x2 > columnIndex {
                distanceX += 1000000 - 1
            }
        }
        for _, rowIndex := range emptyRows {
            if y1 < rowIndex && y2 > rowIndex {
                distanceY += 1000000 - 1
            }
        }
        distanceSum += distanceX + distanceY
    }
}

  return distanceSum
}
