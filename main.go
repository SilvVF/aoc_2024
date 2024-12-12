package main

import (
	"bufio"
	"fmt"
	"main/days"
	"os"
	"path/filepath"
)

func main() {

	daysArr := []any{
		days.NewDay1(),
		days.NewDay2(),
		days.NewDay3(),
		days.NewDay4(),
		days.NewDay5(),
		days.NewDay6(),
		days.NewDay7(),
		days.NewDay8(),
		days.NewDay9(),
		days.NewDay10(),
		days.NewDay11(),
		days.NewDay12(),
	}

	run := func(d int, test bool) {

		day := daysArr[d-1]

		var input []string

		fillInput := func(num int) {
			var path string
			if test {
				path = fmt.Sprintf("d%d%s.txt", num, "test")
			} else {
				path = fmt.Sprintf("d%d.txt", num)
			}
			input = readLines(filepath.Join("inputs", path))
		}

		output := func(p1 any, p2 any) {
			fmt.Printf("Part 1: %v\n", p1)
			fmt.Printf("Part 2: %v\n", p2)
		}

		switch day := day.(type) {
		case days.Day[int]:
			fillInput(day.Num())
			p1 := day.Part1(input)
			p2 := day.Part2(input)
			output(p1, p2)
		case days.Day[int64]:
			fillInput(day.Num())
			p1 := day.Part1(input)
			p2 := day.Part2(input)
			output(p1, p2)

		}
	}

	run(12, false)
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	out := []string{}
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out
}
