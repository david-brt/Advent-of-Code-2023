package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Character struct {
	Value byte
	// 'w' for whitespace (int this case dots), 'd' for digits, 's' for special characters
	Type                  byte
	AdjacentToSpecialChar bool
}

func main() {
	//lines := getLines("../input.txt")
	lines := getLines("test.txt")
	markAdjacent(lines)
	fmt.Println(getPartNumberSum(lines))

}

func getLines(filePath string) [][]Character {
	var lines [][]Character

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		var chars []Character
		for _, byte := range line {
			char := Character{
				byte,
				getCharacterType(byte),
				false,
			}
			chars = append(chars, char)
		}
		lines = append(lines, chars)
	}
	return lines
}

// returns 'w' for whitespace (dots), 'd' for digits, 's' for special characters
func getCharacterType(char byte) byte {
	if int(char) >= 48 && int(char) <= 57 {
		return 'd'
	}
	if int(char) == 46 {
		return 'w'
	}
	return 's'
}

func markAdjacent(lines [][]Character) {
	for i, line := range lines {
		for j, char := range line {
			if char.Type == 's' {
				// mark above
				if i > 0 {
					for k := j - 1; k <= j+1; k++ {
						lines[i-1][k].AdjacentToSpecialChar = true
					}
				}
				// mark left
				lines[i][j-1].AdjacentToSpecialChar = true
				// mark right
				lines[i][j+1].AdjacentToSpecialChar = true
				// mark below
				if i < len(lines)-1 {
					for k := j - 1; k <= j+1; k++ {
						lines[i+1][k].AdjacentToSpecialChar = true
					}
				}
			}
		}
	}
}

func getPartNumberSum(lines [][]Character) int {
	res := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j].Type == 'd' && lines[i][j].AdjacentToSpecialChar {
				num, numLength := getNumber(lines[i], j)
				j = numberStart(lines[i], j) + numLength - 1
				res += num
			}
		}
	}
	return res
}

func numberStart(line []Character, currentIndex int) int {
	if line[currentIndex].Type != 'd' {
		return currentIndex + 1
	}
	if currentIndex == 0 {
		return currentIndex
	}
	return numberStart(line, currentIndex-1)
}

// takes an index of a digit and returns the number that contains that digit as well as its length
func getNumber(line []Character, index int) (int, int) {
	var num []byte
	index = numberStart(line, index)
	//for line[index].Type == 'd' {
	//num = append(num, line[index].Value)
	//index++
	//if index >= len(line) {
	//	break
	//}
	//}
	for i := index; i < len(line); i++ {
		if line[i].Type != 'd' {
			break
		}
		num = append(num, line[i].Value)
	}
	res, err := strconv.Atoi(string(num))
	if err != nil {
		log.Fatalln(err)
	}
	return res, len(num)
}
