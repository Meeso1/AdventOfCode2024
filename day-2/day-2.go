package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() ([][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lists := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		currentList := []int{}
		for _, s := range split {
			number, err := strconv.ParseInt(s, 0, 0)
			if err != nil {
				return nil, err
			}

			currentList = append(currentList, int(number))
		}

		lists = append(lists, currentList)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lists, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func step(prev int, next int) (bool, bool, bool) {
	change := abs(prev - next)
	gapIsOk := change > 0 && change < 4
	increased := next > prev
	decreased := next < prev
	return gapIsOk, increased, decreased
}

func analyze(report []int) (int, int, int) {
	wrongGaps := 0
	increasing := 0
	decreasing := 0

	for i := 1; i < len(report); i++ {
		gapIsOk, increased, decreased := step(report[i-1], report[i])

		if !gapIsOk {
			wrongGaps++
		}

		if increased {
			increasing++
		}

		if decreased {
			decreasing++
		}
	}

	//fmt.Println(report, ": wrong gaps:", wrongGaps, ", increasing: ", increasing, ", decreasing", decreasing)
	return wrongGaps, increasing, decreasing
}

func checkResult(wrongGaps int, increasing int, decreasing int) bool {
	if wrongGaps > 0 {
		return false
	}

	if increasing > 0 && decreasing > 0 {
		return false
	}

	return true
}

func isSafe(report []int) bool {
	wrongGaps, increasing, decreasing := analyze(report)
	return checkResult(wrongGaps, increasing, decreasing)
}

func removeItemAt(report []int, index int) []int {
	result := make([]int, len(report)-1)
	copy(result[:index], report[:index])
	copy(result[index:], report[index+1:])
	return result
}

func isSafeWithSkip(report []int) bool {
	for i := 0; i < len(report); i++ {
		reduced := removeItemAt(report, i)
		result := isSafe(reduced)

		//fmt.Println("Checking", reduced, "result:", result)

		if result {
			return true
		}
	}

	return false
}

func part1(reports [][]int) {
	numSafe := 0
	for _, report := range reports {
		if isSafe(report) {
			numSafe++
		}
	}

	fmt.Println("Safe reports:", numSafe)
}

func part2(reports [][]int) {
	numSafe := 0
	for _, report := range reports {
		if isSafeWithSkip(report) {
			numSafe += 1
		}
	}

	fmt.Println("Safe reports (2):", numSafe)
}

func main() {
	lists, err := readInput()
	if err != nil {
		panic(err)
	}

	fmt.Println("Read lists:", len(lists))
	part1(lists)
	part2(lists)
}
