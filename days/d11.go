package days

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Day11 struct{}

func NewDay11() Day[int] {
	return &Day11{}
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

func (d *Day11) Part1(input []string) int {
	stones := d.parseInput(input)

	blinks := 25
	for i := 0; i < blinks; i++ {
		cpy := []int64{}
		for _, stone := range stones {

			if stone == 0 {
				cpy = append(cpy, 1)
				continue
			}
			str := fmt.Sprint(stone)
			if len(str)%2 == 0 {
				l, _ := strconv.ParseInt(str[0:len(str)/2], 10, 64)
				r, _ := strconv.ParseInt(str[len(str)/2:], 10, 64)

				cpy = append(cpy, l, r)
				continue
			}

			cpy = append(cpy, int64(stone*2024))
		}
		stones = cpy
	}

	return len(stones)
}

func (d *Day11) Part2(input []string) int {
	stones := d.parseInput(input)
	blinks := 75
	total := 0
	wg := sync.WaitGroup{}
	for _, start := range stones {
		wg.Add(1)
		go func() {
			defer wg.Done()
			prev := []int64{start}
			curr := []int64{}

			for i := 0; i < blinks; i++ {
				for _, stone := range prev {
					if stone == 0 {
						curr = append(curr, 1)
						continue
					}
					str := fmt.Sprint(stone)
					if len(str)%2 == 0 {
						l, _ := strconv.ParseInt(str[0:len(str)/2], 10, 64)
						r, _ := strconv.ParseInt(str[len(str)/2:], 10, 64)

						curr = append(curr, l, r)
						continue
					}

					curr = append(curr, int64(stone*2024))
				}
				prev = curr
				curr = []int64{}
			}
			total += len(prev)
		}()
	}
	wg.Wait()
	return total
}
