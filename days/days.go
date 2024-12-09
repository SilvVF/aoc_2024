package days

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

type Day[T Number] interface {
	Num() int
	Part1(input []string) T
	Part2(input []string) T
}
