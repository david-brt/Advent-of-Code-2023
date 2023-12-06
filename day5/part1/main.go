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
	lines := getLines("test.txt")
	lines, seeds := parseSeeds(lines)
	mappings := parseMappings(lines)
	res := 0
	for _, seed := range seeds {
		location := traverseMappings(mappings, seed)
		if location < res || res == 0 {
			res = location
		}
	}
	fmt.Println(res)
}

// extracts the seeds and returns lines with the first two elements removed
func parseSeeds(lines []string) ([]string, []int) {
	seedsString, lines := lines[0], lines[2:]
	seedsString, _ = strings.CutPrefix(seedsString, "seeds: ")
	seeds := strings.Split(seedsString, " ")
	return lines, stringSliceToIntSlice(seeds)
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
		currentSource = findDestination(mapping, currentSource)
	}
	return currentSource
}

func findDestination(mapping map[string][]int, source int) int {
	sources := mapping["sources"]
	ranges := mapping["ranges"]
	destinations := mapping["destinations"]
	for i := 0; i < len(sources); i++ {
		if source < sources[i] || source > sources[i]+ranges[i] {
			continue
		}
		diff := destinations[i] - sources[i]
		return source + diff
	}
	return source
}
