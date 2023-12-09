package main

import "fmt"

func main() {
  exTime = time
  exMaxDist = maxDist
  y1 :=  exTime - 1
  y2 :=  (exTime - 2) * 2
  a, b := regress(y1, y2)
  lower := searchLower(a, b, exMaxDist)
  upper := searchUpper(a, b, exMaxDist)
  res := upper - lower + 1
  fmt.Println(res)
}

// regresses quadratic formular f(x) = ax^2 + bx + c (c is 0 in our case) based on two given points
// returns parameters a and b
func regress(y1 int, y2 int) (int,int) {
  a := (y2 - 2 * y1) / 2
	b := y1 - a
  return a, b
}

func searchLower(a int, b int, y int) int {
  for x:= 0; x < b; x++ {
    curr := a * x * x + b * x
    if curr >= y {
      return x
    }
  }
  return -1
}

func searchUpper(a int, b int, y int) int {
  for x:= b; x >= 0; x-- {
    curr := a * x * x + b * x
    if curr >= y {
      return x
    }
  }
  return -1
}
