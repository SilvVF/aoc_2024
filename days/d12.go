package days

import (
	"slices"
)

type Day12 struct {
	grid   [][]rune
	seen   map[Pos]struct{}
	summed map[Pos]struct{}
	sides  map[Pos][]int
}

func NewDay12() Day[int] {
	return &Day12{
		seen:   map[Pos]struct{}{},
		summed: map[Pos]struct{}{},
	}
}

func (d *Day12) Num() int {
	return 12
}

func (d *Day12) buildGrid(input []string) [][]rune {
	grid := make([][]rune, len(input))
	for i, line := range input {
		for _, c := range line {
			grid[i] = append(grid[i], c)
		}
	}
	return grid
}

func (d *Day12) sumArea(i, j int, v rune) (int, int) {
	inBounds := i < len(d.grid) && i >= 0 && j < len(d.grid[i]) && j >= 0
	if !inBounds {
		return 1, 0
	}
	if _, ok := d.seen[Pos{i, j}]; ok {
		return 0, 0
	}
	if d.grid[i][j] != v {
		return 1, 0
	}

	d.seen[Pos{i, j}] = struct{}{}
	d.summed[Pos{i, j}] = struct{}{}
	dirs := []Pos{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	var p, a int

	for _, dir := range dirs {
		pv, av := d.sumArea(i+dir.x, j+dir.y, v)
		p += pv
		a += av
	}
	return p, 1 + a
}

func (d *Day12) Part1(input []string) int {
	d.grid = d.buildGrid(input)

	t := 0
	for i := range d.grid {
		for j := range d.grid {
			if _, ok := d.summed[Pos{i, j}]; ok {
				continue
			}
			pv, av := d.sumArea(i, j, d.grid[i][j])
			t += pv * av
			for p := range d.seen {
				delete(d.seen, p)
			}
		}
	}

	return t
}

var sides = 0

func (d *Day12) sumAreaSides(i, j int, v rune) (int, int) {

	const (
		LEFT  = 0
		RIGHT = 1
		UP    = 2
		DOWN  = 3
	)

	inBounds := func(x, y int) bool {
		return x < len(d.grid) && x >= 0 && y < len(d.grid[x]) && y >= 0
	}
	if !inBounds(i, j) {
		return 1, 0
	}
	if _, ok := d.seen[Pos{i, j}]; ok {
		return 0, 0
	}
	if d.grid[i][j] != v {
		return 1, 0
	}
	d.seen[Pos{i, j}] = struct{}{}
	d.summed[Pos{i, j}] = struct{}{}
	positions := []Pos{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	a := 0
	invalidate := func(xdir, ydir, dir int) {
		for x, y := i, j; inBounds(x, y); x, y = x+xdir, y+ydir {
			if d.grid[x][y] != v {
				break
			}
			checkx := x + positions[dir].x
			checky := y + positions[dir].y
			if inBounds(checkx, checky) && d.grid[checkx][checky] == v {
				break
			}
			d.sides[Pos{x, y}] = append(d.sides[Pos{x, y}], dir)
		}
	}
	for dir, pos := range positions {
		pv, av := d.sumAreaSides(i+pos.x, j+pos.y, v)
		if pv != 0 && !slices.Contains(d.sides[Pos{i, j}], dir) {
			sides += 1

			switch dir {
			case UP, DOWN:
				invalidate(0, -1, dir)
				invalidate(0, 1, dir)
			case LEFT, RIGHT:
				invalidate(-1, 0, dir)
				invalidate(1, 0, dir)
			}
		}
		a += av
	}
	return 0, 1 + a
}

func (d *Day12) Part2(input []string) int {

	d.grid = d.buildGrid(input)
	d.sides = map[Pos][]int{}
	for i := range d.grid {
		for j := range d.grid[i] {
			d.sides[Pos{i, j}] = []int{}
		}
	}

	for p := range d.summed {
		delete(d.summed, p)
	}
	t := 0
	for i := range d.grid {
		for j := range d.grid[i] {
			if _, ok := d.summed[Pos{i, j}]; ok {
				continue
			}
			for p := range d.seen {
				delete(d.seen, p)
			}
			_, av := d.sumAreaSides(i, j, d.grid[i][j])
			t += sides * av

			sides = 0
		}
	}
	return t
}
