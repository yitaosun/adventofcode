package main

import (
	"strings"
)

func newSet08(s string) Set {
	rv := Set{}
	for _, r := range s {
		rv.Add(r)
	}
	return rv
}

func deduce08(seq []string) map[rune]rune {
	m := map[rune]rune{}
	m1, m4, m7 := Set{}, Set{}, Set{}
	counts := map[rune]int{}
	for _, s := range seq {
		switch len(s) {
		case 2:
			m1 = newSet08(s)
		case 3:
			m7 = newSet08(s)
		case 4:
			m4 = newSet08(s)
		}
		for _, r := range s {
			counts[r]++
		}
	}
	for r, c := range counts {
		switch c {
		case 4:
			m['e'] = r
		case 6:
			m['b'] = r
		case 9:
			m['f'] = r
		}
	}
	m['a'] = m7.RemoveAll(m1).ToSlice()[0].(rune)
	m['c'] = m1.Remove(m['f']).ToSlice()[0].(rune)
	m['d'] = m4.Remove(m['c']).Remove(m['b']).Remove(m['f']).ToSlice()[0].(rune)
	m['g'] = newSet08("abcdefg").
		Remove(m['a']).
		Remove(m['b']).
		Remove(m['c']).
		Remove(m['d']).
		Remove(m['e']).
		Remove(m['f']).ToSlice()[0].(rune)
	return m
}

func solve08B(input []string) int {
	var segs = []string{
		"abcefg",
		"cf",
		"acdeg",
		"acdfg",
		"bcdf",
		"abdfg",
		"abdefg",
		"acf",
		"abcdefg",
		"abcdfg",
	}
	var sum int
	for _, line := range input {
		parts := strings.Split(line, " ")
		m := deduce08(parts[0:10])
		var mods []Set
		for _, seg := range segs {
			mod := Set{}
			for _, r := range seg {
				mod.Add(m[r])
			}
			mods = append(mods, mod)
		}
		var n int
		for i, s := range parts[11:15] {
			for j, mod := range mods {
				if mod.Equals(newSet08(s)) {
					n += j
				}
			}
			if i < 3 {
				n *= 10
			}
		}
		sum += n
	}
	return sum
}

func solve08A(input []string) int {
	uniques := map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}
	var count int
	for _, line := range input {
		parts := strings.Split(line, " ")
		for _, out := range parts[11:15] {
			if _, ok := uniques[len(out)]; ok {
				count++
			}
		}
	}
	return count
}
