package days

type Day interface {
	Num() int
	Part1(input []string) int
	Part2(input []string) int
}
