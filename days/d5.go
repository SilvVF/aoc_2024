package days

import (
	"slices"
	"strconv"
	"strings"
)

type Day5 struct{}

func NewDay5() Day {
	return &Day5{}
}

func (d *Day5) Num() int {
	return 5
}

func (d *Day5) parseInput(input []string) (map[int][]int, [][]int) {
	rules := map[int][]int{}
	inputs := [][]int{}

	for _, line := range input {
		if line == "" {
			continue
		}
		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")
			num, _ := strconv.Atoi(rule[0])
			before, _ := strconv.Atoi(rule[1])

			rules[num] = append(rules[num], before)
		} else {
			values := strings.Split(line, ",")
			arr := make([]int, len(values))
			for i, str := range values {
				num, _ := strconv.Atoi(str)
				arr[i] = num
			}
			inputs = append(inputs, arr)
		}
	}

	return rules, inputs
}

func (d *Day5) Part1(input []string) int {
	rules, inputs := d.parseInput(input)
	t := 0
	for _, arr := range inputs {
		cpy := make([]int, len(arr))
		copy(cpy, arr)
		for i, n := range cpy {
			rules, ok := rules[n]
			if !ok {
				continue
			}

			idx := slices.IndexFunc(arr, func(e int) bool {
				return slices.Contains(rules, e)
			})
			if idx == -1 || i < idx {
				continue
			}

			tmp := cpy[idx]
			cpy[idx] = n
			cpy[i] = tmp

			i = idx
		}
		if slices.Compare(arr, cpy) == 0 {
			e := cpy[len(cpy)/2]
			t += e
		}
	}

	return t
}

func (d *Day5) Part2(input []string) int {
	rules, inputs := d.parseInput(input)
	t := 0
	for _, arr := range inputs {

		cpy := make([]int, len(arr))
		cpyIdx := 0
		cpyMax := len(arr) - 1
		curr := map[int][]int{}

		valid := true
		for i, x := range arr {
			for _, y := range arr[i:] {
				if slices.Contains(rules[y], x) {
					valid = false
				}
			}
			if !valid {
				break
			}
			cpy[cpyIdx] = x
			cpyIdx += 1
		}
		valid = true
		for i := len(arr) - 1; i >= 0; i-- {
			for _, y := range arr[:i] {
				if slices.Contains(rules[arr[i]], y) && !slices.Contains(rules[y], arr[i]) {
					valid = false
				}
			}
			if !valid {
				break
			}
			cpy[cpyMax] = arr[i]
			cpyMax -= 1
		}
		if slices.Compare(arr, cpy) == 0 {
			continue
		}

		for _, n := range arr[cpyIdx : cpyMax+1] {
			curr[n] = rules[n]
		}

	outer:
		for idx := cpyIdx; idx <= cpyMax; idx++ {
			for x := range curr {
				found := false
				for _, ry := range curr {
					if slices.Contains(ry, x) {
						found = true
						break
					}
				}
				if !found {
					cpy[idx] = x
					delete(curr, x)
					continue outer
				}
			}
		}
		if slices.Compare(arr, cpy) != 0 {
			e := cpy[len(cpy)/2]
			t += e
		}
	}
	return t
}
