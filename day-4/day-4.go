package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Matrix struct {
	rows int
	cols int
	data []int
}

func (m *Matrix) At(i int, j int) int {
	return m.data[i*m.cols+j]
}

func (m *Matrix) Row(i int) []int {
	return m.data[i*m.cols : (i+1)*m.cols]
}

func (m *Matrix) Column(i int) <-chan int {
	ch := make(chan int)
	go func() {
		for row := 0; row < m.rows; row++ {
			ch <- m.At(row, i)
		}
		close(ch)
	}()
	return ch
}

func (m *Matrix) DiagonalL(i int) <-chan int {
	ch := make(chan int)
	go func() {
		for x := 0; x < m.rows; x++ {
			if x+i < 0 || x+i >= m.cols {
				continue
			}

			ch <- m.At(x, x+i)
		}
		close(ch)
	}()
	return ch
}

func (m *Matrix) DiagonalR(i int) <-chan int {
	ch := make(chan int)
	go func() {
		for x := 0; x < m.rows; x++ {
			col := m.cols - x - 1 + i
			if col < 0 || col >= m.cols {
				continue
			}

			ch <- m.At(x, col)
		}
		close(ch)
	}()
	return ch
}

func (m *Matrix) Rows() <-chan []int {
	ch := make(chan []int)
	go func() {
		for i := 0; i < m.rows; i++ {
			ch <- m.Row(i)
		}
		close(ch)
	}()
	return ch
}

func (m *Matrix) Cols() <-chan <-chan int {
	ch := make(chan (<-chan int))
	go func() {
		for i := 0; i < m.cols; i++ {
			ch <- m.Column(i)
		}
		close(ch)
	}()
	return ch
}

func (m *Matrix) Diagonals() <-chan <-chan int {
	ch := make(chan (<-chan int))
	go func() {
		for i := -m.cols + 1; i < m.cols-1; i++ {
			ch <- m.DiagonalL(i)
			ch <- m.DiagonalR(i)
		}
		close(ch)
	}()
	return ch
}

func readInput() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func runeToInt(c rune) int {
	switch c {
	case 'X':
		return 1
	case 'M':
		return 2
	case 'A':
		return 3
	case 'S':
		return 4
	default:
		return 0
	}
}

func toMatrix(lines []string) (*Matrix, error) {
	rows := len(lines)
	cols := len(lines[0])
	data := []int{}
	for _, line := range lines {
		if len(line) != cols {
			return nil, errors.New("Rows have different lengths")
		}

		rowData := make([]int, len(line))
		for i, c := range line {
			rowData[i] = runeToInt(c)
		}

		data = append(data, rowData...)
	}

	return &Matrix{
		rows: rows,
		cols: cols,
		data: data,
	}, nil
}

func isMatch(a []int) bool {
	return (a[0] == 1 && a[1] == 2 && a[2] == 3 && a[3] == 4) ||
		(a[0] == 4 && a[1] == 3 && a[2] == 2 && a[3] == 1)
}

func findInArray(a []int) int {
	found := 0
	for i := 0; i < len(a)-3; i++ {
		if isMatch(a[i : i+4]) {
			found++
		}
	}

	return found
}

func findInChannel(channel <-chan int) int {
	array := []int{}
	for num := range channel {
		array = append(array, num)
	}

	return findInArray(array)
}

func part1(matrix *Matrix) {
	found := 0
	for row := range matrix.Rows() {
		found += findInArray(row)
	}

	for column := range matrix.Cols() {
		found += findInChannel(column)
	}

	for diag := range matrix.Diagonals() {
		found += findInChannel(diag)
	}

	fmt.Println("Part 1:", found)
}

func isMas(matrix *Matrix, i int, j int) bool {
	toSign := func(val int) int {
		switch val {
		case runeToInt('M'):
			return -1
		case runeToInt('S'):
			return 1
		default:
			return 0
		}
	}

	if matrix.At(i+1, j+1) != runeToInt('A') {
		return false
	}

	diagL := toSign(matrix.At(i, j)) * toSign(matrix.At(i+2, j+2))
	diagR := toSign(matrix.At(i+2, j)) * toSign(matrix.At(i, j+2))
	return diagL == -1 && diagR == -1
}

func part2(m *Matrix) {
	found := 0
	for i := 0; i < m.rows-2; i++ {
		for j := 0; j < m.cols-2; j++ {
			if isMas(m, i, j) {
				found++
			}
		}
	}

	fmt.Println("Part 2:", found)
}

func main() {
	lines, err := readInput()
	if err != nil {
		panic(err)
	}

	matrix, err := toMatrix(lines)
	if err != nil {
		panic(err)
	}

	fmt.Println("Matrix size:", matrix.rows, "x", matrix.cols)
	part1(matrix)
	part2(matrix)
}
