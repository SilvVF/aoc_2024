package days

import (
	"strconv"
	"strings"
	"unicode"
)

const (
	MUL   TokenType = "MUL"
	OPEN  TokenType = "OPEN"
	CLOSE TokenType = "CLOSE"
	NUM   TokenType = "NUM"
	SEP   TokenType = "SEP"
	DO    TokenType = "DO"
	DONT  TokenType = "DONT"
	ANY   TokenType = "ANY"
)

type TokenType = string
type Token struct {
	Raw  string
	Type TokenType
}

func NewDay3() Day {
	return &Day3{}
}

type Day3 struct{}

func (d *Day3) Num() int {
	return 3
}

func (d *Day3) Part1(input []string) int {
	return d.parseTokens(strings.Join(input, ""), false)
}

func (d *Day3) Part2(input []string) int {
	return d.parseTokens(strings.Join(input, ""), true)
}

func (d *Day3) parseTokens(input string, useEnableCmd bool) int {

	tokens := []Token{}
	enabled := true
	t := 0

	for i := 0; i < len(input); i++ {
		c := string(input[i])
		switch {
		case c == "m":
			if input[i:i+3] == "mul" {
				tokens = []Token{}
				tokens = append(tokens, Token{Raw: "mul", Type: MUL})
				i += 2
			} else {
				tokens = []Token{}
			}
		case unicode.IsNumber([]rune(c)[0]):
			for i+1 < len(input) && unicode.IsNumber([]rune(string(input[i+1]))[0]) {
				c += string(input[i+1])
				i += 1
			}
			tokens = append(tokens, Token{Raw: c, Type: NUM})
		case c == ",":
			tokens = append(tokens, Token{Raw: ",", Type: SEP})
		case c == "(":
			tokens = append(tokens, Token{Raw: "(", Type: OPEN})
		case c == ")":
			tokens = append(tokens, Token{Raw: ")", Type: CLOSE})
			if len(tokens) == 6 && (!useEnableCmd || enabled) && d.check(tokens[1:]) {
				n1, _ := strconv.Atoi(tokens[2].Raw)
				n2, _ := strconv.Atoi(tokens[4].Raw)
				t += n1 * n2
			}
			tokens = []Token{}
		case c == "d":
			if i+6 < len(input) && input[i:i+7] == "don't()" {
				enabled = false
				i += 6
			} else if i+3 < len(input) && input[i:i+4] == "do()" {
				enabled = true
				i += 3
			}
			tokens = []Token{}
		default:
			tokens = []Token{}
		}
	}

	return t
}

func (d *Day3) check(arr []Token) bool {
	return arr[0].Type == OPEN && arr[1].Type == NUM && arr[2].Type == SEP && arr[3].Type == NUM && arr[4].Type == CLOSE
}
