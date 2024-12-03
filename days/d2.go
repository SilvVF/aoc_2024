package days

import (
	"strconv"
	"strings"
)

type Day2 struct{}

func NewDay2() Day {
	return &Day2{}
}

func (d *Day2) Num() int { return 2 }

func dec(arr []int) bool {
	prev := 0
	for i, num := range arr {
		if i == 0 {
			prev = num
		} else if prev > num {
			if absDiffInt(prev, num) > 3 {
				return false
			}
			prev = num
		} else {
			return false
		}
	}
	return true
}

func inc(arr []int) bool {
	prev := 0
	for i, num := range arr {
		if i == 0 {
			prev = num
		} else if prev < num {
			if absDiffInt(prev, num) > 3 {
				return false
			}
			prev = num
		} else {
			return false
		}
	}
	return true
}

func (d *Day2) Part1(input []string) int {

	safe := make([][]int, 0, len(input))

	for _, line := range input {
		report := []int{}
		nums := strings.Split(line, " ")
		for _, s := range nums {
			if s != " " {
				n, _ := strconv.Atoi(s)
				report = append(report, n)
			}
		}

		if inc(report) || dec(report) {
			safe = append(safe, report)
		}
	}
	return len(safe)
}

func decFC(arr []int, fc int) bool {

	if fc > 1 {
		return false
	}

	prev := 0
	for i, num := range arr {
		if i == 0 {
			prev = num
		} else if prev > num {
			if absDiffInt(prev, num) > 3 {
				return rmvAdj(i, arr, fc+1)
			}
			prev = num
		} else {
			return rmvAdj(i, arr, fc+1)
		}
	}
	return true
}

func rmvAdj(i int, arr []int, fc int) bool {
	arrs := [][]int{}
	if i > 0 {
		arrs = append(arrs, removeCpy(arr, i-1))
	}
	if i < len(arr)-1 {
		arrs = append(arrs, removeCpy(arr, i+1))
	}
	arrs = append(arrs, removeCpy(arr, i))
	for _, report := range arrs {
		if incFC(report, fc) || decFC(report, fc) {
			return true
		}
	}
	return false
}

func incFC(arr []int, fc int) bool {

	if fc > 1 {
		return false
	}

	prev := 0
	for i, num := range arr {
		if i == 0 {
			prev = num
		} else if prev < num {
			if absDiffInt(prev, num) > 3 {
				return rmvAdj(i, arr, fc+1)
			}
			prev = num
		} else {
			return rmvAdj(i, arr, fc+1)
		}
	}
	return true
}

func (d *Day2) Part2(input []string) int {
	safe := make([][]int, 0, len(input))

	for _, line := range input {
		report := []int{}
		nums := strings.Split(line, " ")
		for _, s := range nums {
			if s != " " {
				n, _ := strconv.Atoi(s)
				report = append(report, n)
			}
		}

		if incFC(report, 0) || decFC(report, 0) {
			safe = append(safe, report)
		}
	}
	return len(safe)
}
