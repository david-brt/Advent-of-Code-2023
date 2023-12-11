package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

type Coord struct {
  x int
  y int
}

func main() {
  universe := getLines("example.txt")
  universe = expandUniverse(universe)
  galaxies := findGalaxies(universe)
  res := galaxyDistances(galaxies)
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

func expandUniverse(universe [][]byte) [][]byte {
  universe = expandVertically(universe)
  universe = expandHorizontally(universe)
  return universe
}

func expandVertically(universe [][]byte) [][]byte {
  for i:=0; i<len(universe); i++ {
    if slices.Contains(universe[i], '#') {
      continue
    }
    universe = slices.Insert(universe, i, universe[i])
    i++
  }
  return universe
}

func expandHorizontally(universe [][]byte) [][]byte {
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
    for j:=0; j<len(universe); j++ {
      universe[j] = slices.Insert(universe[j], i, '.')
    }
    i++
  }
  return universe
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

func galaxyDistances(galaxies []Coord) int {
  distanceSum := 0
  for i, galaxy := range galaxies {
    for j:=i+1; j<len(galaxies); j++ {
      distanceX := int(math.Abs(float64(galaxy.x - galaxies[j].x)))
      distanceY := int(math.Abs(float64(galaxy.y - galaxies[j].y)))
      distanceSum += distanceX + distanceY
    }
  }
  return distanceSum
}
