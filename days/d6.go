package days

import (
	"fmt"
	"slices"
)

type Day6 struct{}

func NewDay6() *Day6 {
	return &Day6{}
}

func (d *Day6) Num() int {
	return 6
}

const (
	PASSED = 1
	EMPTY  = 0
	OCC    = -1
	OCCHIT = -2
	BLOCK  = -3
)

func (d *Day6) buildGrid(input []string) (row, col int, dir rune, grid [][]int) {
	grid = make([][]int, len(input))

	for i, line := range input {
		grid[i] = make([]int, len(line))
		for j, c := range line {
			var v int
			switch c {
			case '.':
				v = EMPTY
			case '#':
				v = OCC
			case '^', '<', '>', 'v':
				row, col = i, j
				dir = c
				v = 1
			}
			grid[i][j] = v
		}
	}
	return row, col, dir, grid
}

type Pos struct {
	x int
	y int
}

func (d *Day6) getNextDir(dir rune) rune {
	switch dir {
	case '^':
		return '>'
	case '>':
		return 'v'
	case '<':
		return '^'
	case 'v':
		return '<'
	}
	panic("invalid dir")
}

func (d *Day6) Part2(input []string) int {
	row, col, dir, grid := d.buildGrid(input)
	hit := map[Pos][]rune{}
	start := Pos{row, col}
	checkBounds := func(row int, col int) bool {
		return row < len(grid) && row >= 0 && col < len(grid[row]) && col >= 0
	}
	raycastNextDirForHit := func(row, col int, dir rune) (i, j int, ok bool) {
		switch dir {
		case '^':
			idx := slices.Index(grid[row][col+1:], OCC)
			return row, idx + 1 + col, idx != -1

		case '>':
			for i := row; i < len(grid); i++ {
				if grid[i][col] == OCC {
					return i, col, true
				}
			}
			return -1, -1, false
		case '<':
			for i := row; i >= 0; i-- {
				if grid[i][col] == OCC {
					return i, col, true
				}
			}
			return -1, -1, false
		case 'v':
			idx := IndexLast(grid[row][:col], OCC)
			return row, idx, idx != -1
		}
		return -1, -1, false
	}
	checkLoop := func(row, col int, block Pos, dir rune) bool {
		if block == start {
			return false
		}
		seen := map[rune][]Pos{}
		seen[dir] = append(seen[dir], block)

		for i, j, ok := raycastNextDirForHit(row, col, dir); ok; i, j, ok = raycastNextDirForHit(i, j, dir) {
			pos := Pos{x: i, y: j}
			dir = d.getNextDir(dir)
			dirs := hit[pos]
			if slices.Contains(seen[dir], pos) || slices.Contains(dirs, dir) {
				return true
			}
			seen[dir] = append(seen[dir], pos)

			switch dir {
			case '^':
				i += 1
			case '>':
				j -= 1
			case '<':
				j += 1
			case 'v':
				i -= 1
			}
		}
		return false
	}
	handleNext := func(nr, nc int) (v int, cont bool) {
		if !checkBounds(nr, nc) {
			return 0, false
		}
		cont = true
		v = 0

		next := grid[nr][nc]
		if next == OCC {
			pos := Pos{nr, nc}
			hit[pos] = append(hit[pos], dir)
			dir = d.getNextDir(dir)
			return
		} else if next == BLOCK || next == PASSED {
			row = nr
			col = nc
			return
		} else {
			grid[nr][nc] = OCC
			if checkLoop(row, col, Pos{nr, nc}, dir) {
				grid[nr][nc] = BLOCK
				v += 1
			} else {
				grid[nr][nc] = next
			}
			row = nr
			col = nc
			return
		}
	}

	blocks := 0
	for {
		grid[row][col] = PASSED
		var nr, nc int
		switch dir {
		case '^':
			nr = row - 1
			nc = col
		case '<':
			nr = row
			nc = col - 1
		case '>':
			nr = row
			nc = col + 1
		case 'v':
			nr = row + 1
			nc = col
		}
		if v, cont := handleNext(nr, nc); !cont {
			break
		} else {
			fmt.Println("after: ", row, col, grid[row][col])
			blocks += v
		}
	}
	return blocks
}

func (d *Day6) Part1(input []string) int {
	row, col, dir, grid := d.buildGrid(input)
	checkBounds := func(row int, col int) bool {
		return row < len(grid) && row >= 0 && col < len(grid[row]) && col >= 0
	}

outer:
	for checkBounds(row, col) {
		switch dir {
		case '^':
			if !checkBounds(row-1, col) {
				break outer
			}

			next := grid[row-1][col]
			if next == OCC {
				dir = '>'
				continue
			}
			grid[row-1][col] += 1
			row -= 1
		case '<':
			if !checkBounds(row, col-1) {
				break outer
			}
			next := grid[row][col-1]
			if next == OCC {
				dir = '^'
				continue
			}
			grid[row][col-1] += 1
			col -= 1
		case '>':
			if !checkBounds(row, col+1) {
				break outer
			}
			next := grid[row][col+1]
			if next == OCC {
				dir = 'v'
				continue
			}
			grid[row][col+1] += 1
			col += 1
		case 'v':
			if !checkBounds(row+1, col) {
				break outer
			}
			next := grid[row+1][col]
			if next == OCC {
				dir = '<'
				continue
			}
			grid[row+1][col] += 1
			row += 1
		}
	}
	total := 0
	for i := range grid {
		for j := range grid {
			if grid[i][j] > 0 {
				total += 1
			}
		}
	}
	return total
}
