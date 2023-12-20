package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Step struct {
	label       string
	operation   byte
	focalLength int
}

type Lens struct {
	label       string
	focalLength int
}

func main() {
	initSequence := parseFile("example.txt")
	lenses := fillBoxes(initSequence)
	power := focusingPower(lenses)

	fmt.Println(power)
}

func parseFile(filePath string) []Step {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}

	steps := strings.Split(line, ",")
	res := make([]Step, len(steps))
	for i, step := range steps {
		for _, c := range step {
			if c == '-' {
				label, _ := strings.CutSuffix(step, "-")
				res[i] = Step{label, '-', 0}
			}
			if c == '=' {
				label, focalString, _ := strings.Cut(step, "=")
				focalLength, _ := strconv.Atoi(focalString)
				res[i] = Step{label, '=', focalLength}
			}
		}
	}

	return res
}

func hashString(s string) int {
	hash := 0
	for _, c := range s {
		hash += int(c)
		hash *= 17
		hash %= 256
	}
	return hash
}

func fillBoxes(steps []Step) [][]Lens {
	boxes := make([][]Lens, 256)
	for _, step := range steps {
		boxIndex := hashString(step.label)
		lensIndex := indexOf(boxes[boxIndex], step.label)
		if step.operation == '-' {
			if lensIndex != -1 {
				boxes[boxIndex] = remove(boxes[boxIndex], lensIndex)
			}
			continue
		}
		if lensIndex == -1 {
			boxes[boxIndex] = append(boxes[boxIndex], Lens{step.label, step.focalLength})
			continue
		}
		boxes[boxIndex][lensIndex] = Lens{step.label, step.focalLength}
	}
	return boxes
}

func focusingPower(boxes [][]Lens) int {
	res := 0
	for i, box := range boxes {
		if len(box) == 0 {
			continue
		}
		for j, lens := range box {
			power := 1
			power *= i + 1
			power *= j + 1
			power *= lens.focalLength
			res += power
		}
	}
	return res
}

func indexOf(lenses []Lens, label string) int {
	for i, lens := range lenses {
		if lens.label == label {
			return i
		}
	}
	return -1
}

func remove(slice []Lens, s int) []Lens {
	return append(slice[:s], slice[s+1:]...)
}
