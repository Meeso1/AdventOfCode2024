package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readInput() (string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	result := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return result, nil
}

func splitOnMul(input string) []string {
	return strings.Split(input, "mul")[1:]
}

func tryReadNumber(input string) (int, string, bool) {
	numStr := ""
	for _, c := range input {
		if !unicode.IsDigit(c) {
			break
		}

		numStr += string(c)
		if len(numStr) >= 3 {
			break
		}
	}

	rest := input[len(numStr):]
	if len(numStr) == 0 {
		return 0, rest, false
	}

	num, err := strconv.ParseInt(numStr, 10, 0)
	if err != nil {
		return 0, rest, false
	}

	return int(num), rest, true
}

func calculateIfCorrect(input string) (int, bool) {
	if input[0] != '(' {
		return 0, false
	}
	input = input[1:]

	num1, input, ok := tryReadNumber(input)
	if !ok {
		return 0, false
	}

	if input[0] != ',' {
		return 0, false
	}
	input = input[1:]

	num2, input, ok := tryReadNumber(input)
	if !ok {
		return 0, false
	}

	if input[0] != ')' {
		return 0, false
	}

	return num1 * num2, true
}

func part1(input string) int {
	sum := 0
	for _, line := range splitOnMul(input) {
		result, ok := calculateIfCorrect(line)
		if ok {
			sum += result
		}
	}
	return sum
}

func splitOnDo(input string) []string {
	return strings.Split(input, "do()")
}

func partBeforeDont(input string) string {
	return strings.SplitN(input, "don't()", 2)[0]
}

func part2(input string) int {
	sum := 0
	for _, line := range splitOnDo(input) {
		doPart := partBeforeDont(line)
		for _, instruction := range splitOnMul(doPart) {
			result, ok := calculateIfCorrect(instruction)
			if ok {
				sum += result
			}
		}
	}
	return sum
}

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}

	result1 := part1(input)
	fmt.Println("Part 1:", result1)

	result2 := part2(input)
	fmt.Println("Part 2:", result2)
}
