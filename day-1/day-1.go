package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() ([]int, []int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	firstList := []int{}
	secondList := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		firstNumber, err := strconv.ParseInt(parts[0], 0, 0)
		if err != nil {
			return nil, nil, err
		}

		secondNumber, err := strconv.ParseInt(parts[len(parts)-1], 0, 0)
		if err != nil {
			return nil, nil, err
		}

		firstList = append(firstList, int(firstNumber))
		secondList = append(secondList, int(secondNumber))
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return firstList, secondList, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(firstList []int, secondList []int) error {
	difference := 0
	for i := 0; i < len(firstList); i++ {
		difference += abs(secondList[i] - firstList[i])
	}

	_, err := fmt.Println("Part 1 answer:", difference)
	return err
}

func part2(firstList []int, secondList []int) error {
	frequencies := map[int]int{}
	for i := 0; i < len(secondList); i++ {
		if _, ok := frequencies[secondList[i]]; ok {
			frequencies[secondList[i]]++
		} else {
			frequencies[secondList[i]] = 1
		}
	}

	sum := 0
	for i := 0; i < len(firstList); i++ {
		if count, ok := frequencies[firstList[i]]; ok {
			sum += firstList[i] * count
		}
	}

	_, err := fmt.Println("Part 2 answer:", sum)
	return err
}

func main() {
	firstList, secondList, err := readInput()
	if err != nil {
		panic(err)
	}

	sort.SliceStable(firstList, func(i, j int) bool {
		return firstList[i] < firstList[j]
	})
	sort.SliceStable(secondList, func(i, j int) bool {
		return secondList[i] < secondList[j]
	})

	err = part1(firstList, secondList)
	if err != nil {
		panic(err)
	}

	err = part2(firstList, secondList)
	if err != nil {
		panic(err)
	}
}
