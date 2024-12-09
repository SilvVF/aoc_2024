package days

import (
	"fmt"
	"strconv"
	"strings"
)

type Day7 struct {
	nums     []int
	expected int64
}

func NewDay7() Day[int64] {
	return &Day7{}
}

func (d *Day7) Num() int {
	return 7
}

func (d *Day7) dfs(i int, v int64) bool {
	if i == len(d.nums) || v > d.expected {
		return v == d.expected
	}
	return d.dfs(i+1, v+int64(d.nums[i])) || d.dfs(i+1, max(v, 1)*int64(d.nums[i]))
}

func (d *Day7) dfs2(i int, v int64) bool {
	if i == len(d.nums) || v > d.expected {
		return v == d.expected
	}

	comb, _ := strconv.ParseInt(fmt.Sprintf("%d%d", v, d.nums[i]), 10, 64)

	return d.dfs2(i+1, v+int64(d.nums[i])) || d.dfs2(i+1, max(v, 1)*int64(d.nums[i])) || d.dfs2(i+1, comb)
}

func (d *Day7) Part1(input []string) int64 {
	var total int64
	for _, line := range input {
		split := strings.Split(line, ":")
		d.expected, _ = strconv.ParseInt(split[0], 10, 64)
		rest := strings.Split(split[1], " ")
		d.nums = make([]int, len(rest))
		for i, str := range rest {
			n, _ := strconv.Atoi(str)
			d.nums[i] = n
		}
		if d.dfs(0, 0) {
			total += d.expected
		}
	}
	return total
}

func (d *Day7) Part2(input []string) int64 {
	var total int64
	for _, line := range input {
		split := strings.Split(line, ":")
		d.expected, _ = strconv.ParseInt(split[0], 10, 64)
		rest := strings.Split(split[1], " ")
		d.nums = make([]int, len(rest))
		for i, str := range rest {
			n, _ := strconv.Atoi(str)
			d.nums[i] = n
		}
		if d.dfs2(0, 0) {
			total += d.expected
		}
	}
	return total
}
