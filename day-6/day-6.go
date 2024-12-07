package main

import(
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Grid struct {
	rows 	int
	cols 	int
	guardX	int
	guardY	int
	data 	[]int
}

func (m *Grid) At(i int, j int) int {
	return m.data[i*m.cols+j]
}

func (m *Grid) Row(i int) []int {
	return m.data[i*m.cols : (i+1)*m.cols]
}

func (m *Grid) Column(i int) <-chan int {
	ch := make(chan int)
	go func() {
		for row := 0; row < m.rows; row++ {
			ch <- m.At(row, i)
		}
		close(ch)
	}()
	return ch
}

func readLines() ([]string, error) {
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

func toGrid(lines []string) (*Grid, error) {
	rows := len(lines)
	cols := len(lines[0])
	data := []int{}
	guardX := -1
	guardY := -1
	for rowIndex, line := range lines {
		if len(line) != cols {
			return nil, errors.New("Rows have different lengths")
		}

		rowData := make([]int, len(line))
		for i, c := range line {
			switch c {
			case '.':
				rowData[i] = 0
				break
			case '#':
				rowData[i] = 1
				break
			case '^':
				rowData[i] = 0
				guardX = rowIndex
				guardY = i
				break
			default:
				return nil, errors.New(fmt.Sprintf("Invalid character at %d x %d: %c", rowIndex, i, c))
			}
		}

		data = append(data, rowData...)
	}

	if guardX == -1 || guardY == -1 {
		return nil, errors.New("Grid does not contain guard")
	}

	return &Grid{
		rows: rows,
		cols: cols,
		data: data,
		guardX: guardX,
		guardY: guardY,
	}, nil
}

func part1(grid *Grid) {
	// TODO: Implement
	fmt.Println("Part 1:", grid.rows)
}

func main() {
	lines, err := readLines()
	if err != nil {
		panic(err)
	}

	grid, err := toGrid(lines)
	if err != nil {
		panic(err)
	}

	part1(grid)
}
