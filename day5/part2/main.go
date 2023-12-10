package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getLines("example.txt")
	lines, seedRanges := parseSeeds(lines)
	mappings := parseMappings(lines)
	location := 0

	for _, seedRange := range seedRanges {
		start := seedRange["start"]
		end := start + seedRange["range"] - 1
		current := recurse(mappings, 0, start, end)
		if current < location || location == 0 {
			location = current
		}
	}

	fmt.Println(location)
}

// solves the problem blazing fast by skipping ranges in the mapping. super proud of this one.
func recurse(mappings []map[string][]int, mapIndex int, start int, end int) int {
	destination, rangeEnd := findDestination(mappings[mapIndex], start)
	if mapIndex == len(mappings)-1 {
		if rangeEnd == -1 || rangeEnd >= end {
			return destination
		}
		return min(destination, recurse(mappings, mapIndex, rangeEnd+1, end))
	}
	if rangeEnd == -1 {
		return recurse(mappings, mapIndex+1, start, end)
	}
	if rangeEnd >= end {
		newEnd, _ := findDestination(mappings[mapIndex], end)
		return recurse(mappings, mapIndex+1, destination, newEnd)
	}
	newEnd, _ := findDestination(mappings[mapIndex], end)
	return min(recurse(mappings, mapIndex+1, destination, newEnd), recurse(mappings, mapIndex, rangeEnd+1, end))
}

// extracts the seeds and returns lines with the first two elements removed
func parseSeeds(lines []string) ([]string, []map[string]int) {
	seedsString, lines := lines[0], lines[2:]
	seedsString, _ = strings.CutPrefix(seedsString, "seeds: ")
	vals := stringSliceToIntSlice(strings.Split(seedsString, " "))
	res := make([]map[string]int, len(vals)/2)
	for i := 0; i < len(vals); i += 2 {
		seedRange := make(map[string]int)
		seedRange["start"] = vals[i]
		seedRange["range"] = vals[i+1]
		res[i/2] = seedRange
	}
	return lines, res
}

func getLines(filePath string) []string {
	var lines []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func parseMappings(lines []string) []map[string][]int {
	var mappings = make([]map[string][]int, 7)
	mapIndex := -1
	for _, line := range lines {
		if isEmpty(line) {
			continue
		}
		if strings.HasSuffix(line, "map:") {
			mapIndex += 1
			mappings[mapIndex] = map[string][]int{
				"sources":      {},
				"destinations": {},
				"ranges":       {},
			}
			continue
		}
		split := strings.Split(line, " ")
		vals := stringSliceToIntSlice(split)

		mappings[mapIndex]["destinations"] = append(mappings[mapIndex]["destinations"], vals[0])
		mappings[mapIndex]["sources"] = append(mappings[mapIndex]["sources"], vals[1])
		mappings[mapIndex]["ranges"] = append(mappings[mapIndex]["ranges"], vals[2])
	}
	return mappings
}

func isEmpty(s string) bool {
	return len(strings.Trim(s, " ")) == 0
}

func stringSliceToIntSlice(a []string) []int {
	res := make([]int, len(a))
	for i, s := range a {
		var err error
		res[i], err = strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return res
}

func traverseMappings(mappings []map[string][]int, seed int) int {
	currentSource := seed
	for _, mapping := range mappings {
		currentSource, _ = findDestination(mapping, currentSource)
	}
	return currentSource
}

// returns the destination for a given source as well as the end of the range
func findDestination(mapping map[string][]int, source int) (int, int) {
	sources := mapping["sources"]
	ranges := mapping["ranges"]
	destinations := mapping["destinations"]
	for i := 0; i < len(sources); i++ {
		if source < sources[i] || source >= sources[i]+ranges[i] {
			continue
		}
		diff := destinations[i] - sources[i]
		return source + diff, sources[i] + ranges[i] - 1
	}
	return source, rangeLimit(mapping, source)
}

// checks all seeds for a range and returns the location with the lowest number
func lowestInRange(seedRange map[string]int, mappings []map[string][]int) int {
	res := 0
	rangeStart := seedRange["start"]
	rangeEnd := rangeStart + seedRange["range"]

	for i := rangeStart; i < rangeEnd; i++ {
		location := traverseMappings(mappings, i)
		if location < res || res == 0 {
			res = location
		}
	}

	return res
}

func rangeLimit(mapping map[string][]int, node int) int {
	rangeStart := -1

	for _, v := range mapping["source"] {
		if (v < rangeStart || rangeStart == -1) && v >= node {
			return v - 1
		}
	}

	return rangeStart
}
