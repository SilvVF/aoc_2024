package main

import (
	"bufio"
	"fmt"
	"main/days"
	"os"
	"path/filepath"
)

func main() {

	days := []days.Day{
		days.NewDay1(),
		days.NewDay2(),
		days.NewDay3(),
		days.NewDay4(),
	}

	run := func(d int, test bool) {
		day := days[d-1]

		var path string
		if test {
			path = fmt.Sprintf("d%d%s.txt", day.Num(), "test")
		} else {
			path = fmt.Sprintf("d%d.txt", day.Num())
		}
		input := readLines(filepath.Join("inputs", path))
		fmt.Printf("Part 1: %d\n", day.Part1(input))
		fmt.Printf("Part 2: %d\n", day.Part2(input))

		day.Part2(input)
	}

	run(4, true)
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
