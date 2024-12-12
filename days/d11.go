package days

import (
	"fmt"
	"strconv"
	"strings"
)

type Day11 struct {
	blinks int
	cache  map[int]map[int64]int
}

func NewDay11() Day[int] {
	return &Day11{
		blinks: 25,
		cache:  map[int]map[int64]int{},
	}
}

func (d *Day11) Num() int {
	return 11
}

func (d *Day11) parseInput(input []string) []int64 {
	arr := []int64{}
	for _, line := range input {
		stones := strings.Split(line, " ")
		for _, s := range stones {
			n, _ := strconv.ParseInt(s, 10, 64)
			arr = append(arr, n)
		}
	}
	return arr
}

func (d *Day11) getNext(stone int64, iter int) int {

	if iter == 0 {
		return 1
	}

	if m, ok := d.cache[iter]; ok {
		if v, ok := m[stone]; ok {
			return v
		}
	}

	if stone == 0 {
		v := d.getNext(1, iter-1)
		d.cache[iter][stone] = v
		return v
	}
	str := fmt.Sprint(stone)
	if len(str)%2 == 0 {
		l, _ := strconv.ParseInt(str[0:len(str)/2], 10, 64)
		r, _ := strconv.ParseInt(str[len(str)/2:], 10, 64)

		v := d.getNext(l, iter-1) + d.getNext(r, iter-1)
		d.cache[iter][stone] = v
		return v
	}

	v := d.getNext(stone*int64(2024), iter-1)
	d.cache[iter][stone] = v
	return v
}

func (d *Day11) Part1(input []string) int {
	stones := d.parseInput(input)

	total := 0

	for i := range d.blinks + 1 {
		d.cache[i] = map[int64]int{}
	}
	for _, stone := range stones {
		total += d.getNext(stone, d.blinks)
	}
	return total
}

func (d *Day11) Part2(input []string) int {
	d.blinks = 75
	return d.Part1(input)
}
