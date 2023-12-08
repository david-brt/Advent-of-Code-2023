package main

import "fmt"

type Race struct {
  time int
  maxDistance int
}

func main() {
  res := 1
  for _, race := range exampleInput {
    res *= getWinningPossibilities(race)
  }
  fmt.Println(res)
}

func getWinningPossibilities(race Race) int {
  possibilities := 0
  for i:=0; i<race.maxDistance; i++ {
    timeLeft := race.time - i
    distance := timeLeft * i
    if distance > race.maxDistance {
      possibilities += 1
    }
  }
  return possibilities
}
