package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Cords struct {
	x int
	y int
}

type Grid struct {
	rows  int
	cols  int
	guard Cords
	dir   Direction
	data  []int
}

func (m *Grid) At(i int, j int) int {
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return -1
	}

	return m.data[i*m.cols+j]
}

func (m *Grid) Advance() (Cords, bool) {
	newCords := m.guard
	direction := m.dir

	switch m.dir {
	case North:
		switch m.At(m.guard.x-1, m.guard.y) {
		case 0:
			newCords = Cords{
				x: m.guard.x - 1,
				y: m.guard.y,
			}
			break
		case 1:
			direction = East
		default:
			return Cords{}, true
		}
		break
	case East:
		switch m.At(m.guard.x, m.guard.y+1) {
		case 0:
			newCords = Cords{
				x: m.guard.x,
				y: m.guard.y + 1,
			}
			break
		case 1:
			direction = South
		default:
			return Cords{}, true
		}
		break
	case South:
		switch m.At(m.guard.x+1, m.guard.y) {
		case 0:
			newCords = Cords{
				x: m.guard.x + 1,
				y: m.guard.y,
			}
			break
		case 1:
			direction = West
		default:
			return Cords{}, true
		}
		break
	case West:
		switch m.At(m.guard.x, m.guard.y-1) {
		case 0:
			newCords = Cords{
				x: m.guard.x,
				y: m.guard.y - 1,
			}
			break
		case 1:
			direction = North
		default:
			return Cords{}, true
		}
		break
	}

	m.guard = newCords
	m.dir = direction
	return newCords, false
}

func (m *Grid) AddWall(i int, j int) bool {
	if m.At(i, j) != 0{
		return false
	}

	m.data[i*m.cols+j] = 1
	return true
}

func (m *Grid) RemoveWall(i int, j int) {
	m.data[i*m.cols+j] = 0
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
		guard: Cords{
			x: guardX,
			y: guardY,
		},
		dir: North,
	}, nil
}

func part1(grid *Grid) {
	visited := map[Cords]int{}
	visited[grid.guard] = 1
	for cords, out := grid.Advance(); !out; cords, out = grid.Advance() {
		visited[cords] = 1
	}

	fmt.Println("Part 1:", len(visited))
}

func part2(grid *Grid) {
	visited := map[Cords]Direction{}
	locations := 0

	visited[grid.guard] = North
	for cords, out := grid.Advance(); !out; cords, out = grid.Advance() {
		if _, ok := visited[cords]; ok {
			locations += 1
		} else {
			visited[cords] = grid.dir
		}
	}

	fmt.Println("Part 2:", locations)
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

	grid, err = toGrid(lines)
	if err != nil {
		panic(err)
	}

	part2(grid)
}
