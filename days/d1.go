package days

import (
	"slices"
	"strconv"
	"strings"
)

func NewDay1() Day {
	return &Day1{}
}

type Day1 struct{}

func (d *Day1) Num() int { return 1 }

func (d *Day1) Part1(input []string) int {

	left := []int{}
	right := []int{}

	for _, line := range input {

		split := strings.Split(line, " ")
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[len(split)-1])

		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	t := 0
	for i := 0; i < len(left); i++ {
		dif := left[i] - right[i]

		if dif < 0 {
			t += dif * -1
		} else {
			t += dif
		}
	}
	return t
}

func (d *Day1) Part2(input []string) int {
	sim := map[int]int{}
	lv := map[int]int{}

	for _, line := range input {

		split := strings.Split(line, " ")
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[len(split)-1])

		if v, ok := lv[l]; ok {
			lv[l] = v + 1
		} else {
			lv[l] = 1
		}

		if v, ok := sim[r]; ok {
			sim[r] = v + 1
		} else {
			sim[r] = 1
		}
	}
	t2 := 0
	for k, v := range sim {
		if a, ok := lv[k]; ok {
			t2 += (k * v) * a
		}
	}

	return t2
}
