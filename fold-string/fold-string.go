package main

import (
	"fmt"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(a, b int) bool {
	return s[a] < s[b]
}

func (s sortRunes) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s sortRunes) Len() int {
	return len(s)
}

func Fold(s string) string {
	m := map[rune]int{}
	for _, r := range s {
		m[r]++
	}

	keys := make([]rune, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sortRunes(keys))

	var sb strings.Builder
	sb.Grow(len(m) * 2 * 4)
	for _, k := range keys {
		sb.WriteRune(k)
		sb.WriteRune(rune(m[k] + 0x30))
	}
	return sb.String()
}

func main() {
	input := "ЖbbbaaacccccaaЖ"
	actual := Fold(input)
	expected := "a5b3c5Ж2"
	fmt.Println(actual == expected)
}
