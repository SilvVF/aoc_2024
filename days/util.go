package days

import (
	"cmp"
	"slices"
)

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func absDiffUint(x, y uint) uint {
	if x < y {
		return y - x
	}
	return x - y
}

func removeCpy(slice []int, s int) []int {
	new := make([]int, len(slice)-1)
	idx := 0
	for i, n := range slice {
		if i == s {
			continue
		}
		new[idx] = n
		idx += 1
	}
	return new
}

// Index returns the index of the last occurrence of v in s,
// or -1 if not present.
func IndexLast[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		idx := len(s) - 1 - i
		if v == s[idx] {
			return idx
		}
	}
	return -1
}

func IndexLastFunc[S ~[]E, E any](s S, f func(e E) bool) int {
	for i := range s {
		idx := len(s) - 1 - i
		if f(s[idx]) {
			return idx
		}
	}
	return -1
}

func TakeLastWhile[S ~[]E, E any](s S, f func(e E) bool) S {
	end := len(s)
	start := 0
	for i := range s {
		idx := len(s) - 1 - i
		if !f(s[idx]) {
			break
		} else {
			start = idx
		}
	}
	return s[start:end]
}

func IndexSlice[S ~[]E, E cmp.Ordered](s S, v S) int {
	l := len(v)
	for i := range s[0 : len(s)-l] {
		if slices.Compare(v, s[i:i+l]) == 0 {
			return i
		}
	}
	return -1
}
