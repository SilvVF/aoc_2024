package days

import (
	"strconv"
)

type Day10 struct {
	seen map[Pos]struct{}
	grid [][]int
}

func NewDay10() Day[int] {
	return &Day10{}
}

func (d *Day10) Num() int {
	return 10
}

func (d *Day10) scoreTrail(i, j, prev int, distinct bool) int {
	boundsCheck := func(i, j int) bool {
		return i >= 0 && i < len(d.grid) && j >= 0 && j < len(d.grid[i])
	}

	if !boundsCheck(i, j) || d.grid[i][j] != prev+1 {
		return 0
	}

	if distinct {
		_, ok := d.seen[Pos{i, j}]
		if ok {
			return 0
		}
		d.seen[Pos{i, j}] = struct{}{}
	}
	if d.grid[i][j] == 9 {
		return 1
	}
	dirs := []Pos{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}

	sum := 0
	for _, dir := range dirs {
		sum += d.scoreTrail(i+dir.x, j+dir.y, d.grid[i][j], distinct)
	}
	return sum
}

func (d *Day10) Part1(input []string) int {
	d.grid = make([][]int, len(input))
	d.seen = map[Pos]struct{}{}
	for i := range input {
		for j := range input {
			n, _ := strconv.Atoi(string(input[i][j]))
			d.grid[i] = append(d.grid[i], n)
		}
	}
	score := 0
	for i := range d.grid {
		for j := range d.grid[i] {
			if d.grid[i][j] == 0 {
				for p := range d.seen {
					delete(d.seen, p)
				}
				score += d.scoreTrail(i, j, -1, true)
			}
		}
	}

	return score
}

func (d *Day10) Part2(input []string) int {
	d.grid = make([][]int, len(input))
	for i := range input {
		for j := range input {
			n, _ := strconv.Atoi(string(input[i][j]))
			d.grid[i] = append(d.grid[i], n)
		}
	}
	score := 0
	for i := range d.grid {
		for j := range d.grid[i] {
			if d.grid[i][j] == 0 {
				score += d.scoreTrail(i, j, -1, false)
			}
		}
	}

	return score
}
