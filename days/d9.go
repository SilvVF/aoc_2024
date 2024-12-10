package days

import (
	"slices"
	"strconv"
	"strings"
)

type Day9 struct{}

func NewDay9() Day[int64] {
	return &Day9{}
}

func (d *Day9) Num() int {
	return 9
}

const (
	FILL = 1
	FREE = -1
)

func (d *Day9) Part1(input []string) int64 {
	block := []int{}

	id := 0
	state := FILL
	for _, c := range strings.Join(input, "") {
		v, _ := strconv.Atoi(string(c))
		switch state {
		case FILL:
			for i := 0; i < v; i++ {
				block = append(block, id)
			}
			id += 1
			state = FREE
		case FREE:
			for i := 0; i < v; i++ {
				block = append(block, FREE)
			}
			state = FILL
		}
	}

	getNext := func() (int, int) {
		free := slices.Index(block, -1)
		part := IndexLastFunc(block, func(e int) bool {
			return e != -1
		})
		return free, part
	}
	check := func(free, last int) bool {
		return last != -1 && free != -1 && free < last
	}
	free, last := getNext()
	for ; check(free, last); free, last = getNext() {
		block[free] = block[last]
		block[last] = -1
	}

	total := int64(0)
	for i, v := range block {
		if v != -1 {
			total += int64(i * v)
		}
	}
	return total
}

func (d *Day9) Part2(input []string) int64 {
	block := []int{}
	space := []byte{}

	id := 0
	state := FILL
	for _, c := range strings.Join(input, "") {
		v, _ := strconv.Atoi(string(c))
		switch state {
		case FILL:
			for i := 0; i < v; i++ {
				block = append(block, id)
				space = append(space, 1)
			}
			id += 1
			state = FREE
		case FREE:
			for i := 0; i < v; i++ {
				block = append(block, FREE)
				space = append(space, 0)
			}
			state = FILL
		}
	}

	offset := len(block) - 1

	getNext := func() (int, int, int, int) {
		fidx := IndexLastFunc(block[0:offset+1], func(e int) bool { return e != -1 })
		file := TakeLastWhile(block[0:fidx+1], func(e int) bool {
			return e == block[fidx]
		})
		offset = fidx - len(file)
		eidx := IndexSlice(space, make([]byte, len(file)))

		return eidx, eidx + len(file), fidx - (len(file) - 1), fidx
	}

	for offset > slices.Index(block, -1) {
		es, ee, fs, fe := getNext()
		if es <= -1 || fs > fe || es > fs {
			continue
		}

		for i := 0; i < ee-es; i++ {
			space[es+i] = 1
			space[fs+i] = 0

			block[es+i] = block[fs+i]
			block[fs+i] = -1
		}
	}
	total := int64(0)
	for i, v := range block {
		if v != -1 {
			total += int64(i * v)
		}
	}
	return total
}
