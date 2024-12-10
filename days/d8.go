package days

type Day8 struct {
	rows int
	cols int
}

const (
	UP = iota + 1
	DOWN
	LEFT
	RIGHT
	DIAG_LU
	DIAG_LD
	DIAG_RU
	DIAG_RD
)

func NewDay8() Day[int] {
	return &Day8{}
}

func (d *Day8) Num() int {
	return 8
}

type Node struct {
	anti map[rune]struct{}
	sig  rune
}

func (d *Day8) nextPos(i, j, dir, radius int) (int, int, bool) {
	var ni, nj int
	switch dir {
	case UP:
		ni, nj = i-radius, j
	case DOWN:
		ni, nj = i+radius, j
	case LEFT:
		ni, nj = i, j-radius
	case RIGHT:
		ni, nj = i, j+radius
	case DIAG_LD:
		ni, nj = i+radius, j-radius
	case DIAG_RD:
		ni, nj = i+radius, j+radius
	case DIAG_RU:
		ni, nj = i-radius, j+radius
	case DIAG_LU:
		ni, nj = i-radius, j-radius
	default:
		panic("bad dir")
	}

	if ni < 0 || ni >= d.rows || nj < 0 || nj >= d.cols {
		return ni, nj, false // Out of bounds
	}
	return ni, nj, true // Valid position
}

func (d *Day8) Part1(input []string) int {
	grid := make([][]*Node, len(input))
	for i, line := range input {
		for _, c := range line {
			grid[i] = append(grid[i], &Node{sig: c, anti: map[rune]struct{}{}})
		}
	}
	d.rows = len(grid)
	d.cols = len(grid[0])

	dirs := []int{UP, DOWN, LEFT, RIGHT, DIAG_LD, DIAG_LU, DIAG_RD, DIAG_RU}

	checkDirs := func(i, j int, sig rune) {
		for _, dir := range dirs {
			radius := 1
			r, c, ok := d.nextPos(i, j, dir, radius)
			for ok {
				if r < 0 || r >= d.rows || c < 0 || c >= d.cols || grid[r][c].sig != sig {
					break
				}

				nr, nc, nextOk := d.nextPos(r, c, dir, radius) // Check 2 * radius
				if nextOk && grid[nr][nc].sig == sig {
					// Mark antinode
					grid[nr][nc].anti[sig] = struct{}{}
				} else if !nextOk {
					continue
				}

				// Increment radius and continue
				radius += 1
				r, c, ok = d.nextPos(i, j, dir, radius)
			}
		}
	}

	for i, line := range input {
		for j, c := range line {
			if c != '.' {
				checkDirs(i, j, c)
			}
		}
	}

	total := 0
	for i := range grid {
		for j := range grid[i] {
			total += len(grid[i][j].anti)
		}
	}
	return total
}

func (d *Day8) Part2(input []string) int {
	return -1
}
