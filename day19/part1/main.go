package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Rule struct {
	category, operation string
	value               int
	action              string
}

type Workflow struct {
	rules    []Rule
	fallback string
}

func parseFile(filePath string) (map[string]Workflow, []map[string]int) {
	wfs := make(map[string]Workflow)
	var parts []map[string]int

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	wfSec := true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			wfSec = false
			continue
		}
		if wfSec {
			wf, wfName := parseWorkflow(line)
			wfs[wfName] = wf
			continue
		}
		part := parsePart(line)
		parts = append(parts, part)
	}
	return wfs, parts
}

func parseWorkflow(wf string) (Workflow, string) {
	pattern := regexp.MustCompile(`([a-z]+){.+,([AR]|[a-z]+)}`)
	matches := pattern.FindStringSubmatch(wf)
	wf = matches[0]
	name := matches[1]
	fallback := matches[2]
	pattern = regexp.MustCompile(`([xmas])([<>])(\d+):([AR]|[a-z]+),`)
	ruleStrings := pattern.FindAllStringSubmatch(wf, -1)
	rules := make([]Rule, len(ruleStrings))
	for i, rule := range ruleStrings {
		value, _ := strconv.Atoi(rule[3])
		rules[i] = Rule{rule[1], rule[2], value, rule[4]}
	}
	return Workflow{rules, fallback}, name
}

func parsePart(part string) map[string]int {
	pattern := regexp.MustCompile(`\d+`)
	matches := pattern.FindAllString(part, -1)
	vals := make([]int, 4)
	for i := range vals {
		vals[i], _ = strconv.Atoi(matches[i])
	}
	return map[string]int{"x": vals[0], "m": vals[1], "a": vals[2], "s": vals[3]}
}

func partValid(wfs map[string]Workflow, part map[string]int, curr string) bool {
	wf := wfs[curr]
	for _, rule := range wf.rules {
		valid := false
		if rule.operation == "<" {
			valid = part[rule.category] < rule.value
		}
		if rule.operation == ">" {
			valid = part[rule.category] > rule.value
		}
		if !valid {
			continue
		}
		switch rule.action {
		case "A":
			return true
		case "R":
			return false
		default:
			return partValid(wfs, part, rule.action)
		}
	}
	switch wf.fallback {
	case "A":
		return true
	case "R":
		return false
	default:
		return partValid(wfs, part, wf.fallback)
	}
}

func partSum(wfs map[string]Workflow, part map[string]int) int {
	res := 0
	if partValid(wfs, part, "in") {
		for _, val := range part {
			res += val
		}
	}
	return res
}

func main() {
	wfs, parts := parseFile("example.txt")
	res := 0
	for _, part := range parts {
		res += partSum(wfs, part)
	}
	fmt.Println(res)
}
