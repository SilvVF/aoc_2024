package days

import "strings"

func NewDay4() Day[int] {
	return &Day4{}
}

type Day4 struct{}

func (d *Day4) Num() int { return 4 }

func (d *Day4) Part1(input []string) int {

	letters := [][]string{}
	for _, line := range input {
		letters = append(letters, strings.Split(line, ""))
	}

	t := 0

	for i, line := range letters {
		for j, c := range line {
			if c == "X" {
				t += d.check(i, j, letters)
			}
		}
	}
	return t
}

func (d *Day4) Part2(input []string) int {
	letters := [][]string{}
	for _, line := range input {
		letters = append(letters, strings.Split(line, ""))
	}

	t := 0

	for i, line := range letters {
		for j, c := range line {
			if c == "A" {
				if d.checkA(i, j, letters) {
					t++
				}
			}
		}
	}
	return t
}

func (d *Day4) checkA(i, j int, letters [][]string) bool {
	if i == 0 || j == 0 || i >= len(letters)-1 || j >= len(letters[i])-1 {
		return false
	}

	tl := letters[i-1][j-1]
	tr := letters[i-1][j+1]
	bl := letters[i+1][j-1]
	br := letters[i+1][j+1]

	if tl == "M" && tr == "S" {
		return bl == "M" && br == "S"
	} else if tl == "M" && tr == "M" {
		return bl == "S" && br == "S"
	} else if tl == "S" && tr == "S" {
		return bl == "M" && br == "M"
	} else if tl == "S" && tr == "M" {
		return bl == "S" && br == "M"
	}
	return false
}

func (d *Day4) check(i, j int, letters [][]string) int {
	out := 0
	// UP
	if i-3 >= 0 {
		if (letters[i][j] + letters[i-1][j] + letters[i-2][j] + letters[i-3][j]) == "XMAS" {
			out += 1
		}
	}
	// DOWN
	if i+3 < len(letters) {
		if (letters[i][j] + letters[i+1][j] + letters[i+2][j] + letters[i+3][j]) == "XMAS" {
			out += 1
		}
	}
	// LEFT
	if j-3 >= 0 {
		if (letters[i][j] + letters[i][j-1] + letters[i][j-2] + letters[i][j-3]) == "XMAS" {
			out += 1
		}
	}
	// RIGHT
	if j+3 < len(letters[i]) {
		if (letters[i][j] + letters[i][j+1] + letters[i][j+2] + letters[i][j+3]) == "XMAS" {
			out += 1
		}
	}

	// DIAG DOWN RIGHT
	if j+3 < len(letters[i]) && i+3 < len(letters) {
		if (letters[i][j] + letters[i+1][j+1] + letters[i+2][j+2] + letters[i+3][j+3]) == "XMAS" {
			out += 1
		}
	}
	// DIAG UP LEFT
	if j-3 >= 0 && i-3 >= 0 {
		if (letters[i][j] + letters[i-1][j-1] + letters[i-2][j-2] + letters[i-3][j-3]) == "XMAS" {
			out += 1
		}
	}
	// DIAG UP RIGHT
	if j+3 < len(letters[i]) && i-3 >= 0 {
		if (letters[i][j] + letters[i-1][j+1] + letters[i-2][j+2] + letters[i-3][j+3]) == "XMAS" {
			out += 1
		}
	}
	// DIAG DOWN LEFT
	if j-3 >= 0 && i+3 < len(letters) {
		if (letters[i][j] + letters[i+1][j-1] + letters[i+2][j-2] + letters[i+3][j-3]) == "XMAS" {
			out += 1
		}
	}
	return out
}
