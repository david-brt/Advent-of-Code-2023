package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  file, err := os.Open("test.txt")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  res := 0

  for scanner.Scan() {
    line := scanner.Text()
    if err := scanner.Err(); err != nil {
      fmt.Println(err)
    }

    for i:=0; i<len(line); i++ {
      // check if byte is integer between 0 and 9
      if line[i] >= 48 && line[i] < 58 {
        res += int(line[i] - '0') * 10
        break
      }
    }

    for i:=len(line)-1; i>= 0; i-- {
      // check if byte is integer between 0 and 9
      if line[i] >= 48 && line[i] < 58 {
        res += int(line[i] - '0')
        break
      }
    }
  }

  fmt.Println(res)
}
