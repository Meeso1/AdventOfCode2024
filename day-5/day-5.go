package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	prev int
	next int
}

func isRuleLine(line string) bool {
	return strings.Contains(line, "|")
}

func parseRuleLine(line string) (Rule, error) {
	parts := strings.Split(line, "|")
	if len(parts) != 2 {
		return Rule{}, errors.New(fmt.Sprintf("Incorrect number of parts in rule: %d", len(parts)))
	}

	first, err := strconv.ParseInt(parts[0], 10, 0)
	if err != nil {
		return Rule{}, err
	}

	second, err := strconv.ParseInt(parts[1], 10, 0)
	if err != nil {
		return Rule{}, err
	}

	return Rule{
		prev: int(first),
		next: int(second),
	}, nil
}

func parseEntryLine(line string) ([]int, error) {
	parts := strings.Split(line, ",")

	result := []int{}
	for _, part := range parts {
		num, err := strconv.ParseInt(part, 10, 0)
		if err != nil {
			return nil, err
		}
		result = append(result, int(num))
	}

	return result, nil
}

func readInput() ([]Rule, [][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	rules := []Rule{}
	entries := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if isRuleLine(line) {
			rule, err := parseRuleLine(line)
			if err != nil {
				return nil, nil, err
			}
			rules = append(rules, rule)
		} else {
			entry, err := parseEntryLine(line)
			if err != nil {
				return nil, nil, err
			}
			entries = append(entries, entry)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, entries, nil
}

func getRequirementsFor(element int, after map[int][]int, all map[int]int) []int {
	requirements, ok := after[element]
	if !ok {
		return []int{}
	}

	result := []int{}
	for _, requirement := range requirements {
		if _, ok := all[requirement]; ok {
			result = append(result, requirement)
		}
	}

	return result
}

func checkEntry(after map[int][]int, entry []int) (bool, int, int) {
	all := map[int]int{}
	for _, element := range entry {
		all[element] = 0
	}

	seen := map[int]int{}
	for index, element := range entry {
		requirements := getRequirementsFor(element, after, all)
		for _, required := range requirements {
			if _, ok := seen[required]; !ok {
				return false, required, index
			}
		}

		seen[element] = index
	}

	return true, -1, -1
}

func makeAfterDict(rules []Rule) map[int][]int {
	after := map[int][]int{}
	for _, rule := range rules {
		if _, ok := after[rule.next]; ok {
			after[rule.next] = append(after[rule.next], rule.prev)
		} else {
			after[rule.next] = []int{rule.prev}
		}
	}

	return after
}

func part1(rules []Rule, entries [][]int) {
	after := makeAfterDict(rules)
	sum := 0
	for _, entry := range entries {
		if ok, _, _ := checkEntry(after, entry); ok {
			sum += entry[int((len(entry)-1)/2)]
		}
	}

	fmt.Println("Part 1:", sum)
}

func findIndex(slice []int, element int) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}

	return -1
}

func getFixedEntry(entry []int, after map[int][]int) []int {
	fixed := make([]int, len(entry))
	copy(fixed, entry)

	ok, movedElement, index := checkEntry(after, fixed)
	for ; !ok; ok, movedElement, index = checkEntry(after, fixed) {

		newFixed := []int{movedElement}
		if index > 0 {
			newFixed = []int{}
			newFixed = append(newFixed, fixed[:index]...)
			newFixed = append(newFixed, movedElement)
		}

		newFixed = append(newFixed, fixed[index:findIndex(fixed, movedElement)]...)
		newFixed = append(newFixed, fixed[findIndex(fixed, movedElement)+1:]...)
		fixed = newFixed
	}

	return fixed
}

func part2(rules []Rule, entries [][]int) {
	after := makeAfterDict(rules)
	sum := 0
	for _, entry := range entries {
		if ok, _, _ := checkEntry(after, entry); ok {
			continue
		} else {
			sum += getFixedEntry(entry, after)[int((len(entry)-1)/2)]
		}
	}

	fmt.Println("Part 2:", sum)
}

func main() {
	rules, entries, err := readInput()
	if err != nil {
		panic(err)
	}

	fmt.Println("Rules:", len(rules))
	fmt.Println("Entries:", len(entries))
	part1(rules, entries)
	part2(rules, entries)
}
