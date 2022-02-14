package main

import (
	"strings"
	"unicode"
)

type path struct {
	src     string
	smalls  Set
	doubled bool
}

func solve12(input []string, allowDouble bool) int {
	m := map[string]map[string]struct{}{}
	for _, line := range input {
		p := strings.Split(line, "-")
		if _, ok := m[p[0]]; !ok {
			m[p[0]] = map[string]struct{}{}
		}
		if _, ok := m[p[1]]; !ok {
			m[p[1]] = map[string]struct{}{}
		}
		m[p[0]][p[1]] = struct{}{}
		m[p[1]][p[0]] = struct{}{}
	}
	var total int
	paths := []*path{
		{src: "start", smalls: Set{}},
	}
	for {
		next := []*path{}
		for _, p := range paths {
			for dest := range m[p.src] {
				if dest == "start" {
					continue
				}
				if dest == "end" {
					total++
					continue
				}
				pp := &path{
					src:     dest,
					smalls:  Set{}.AddAll(p.smalls),
					doubled: p.doubled,
				}
				if unicode.IsLower(rune(dest[0])) {
					if pp.smalls.Has(dest) {
						if !allowDouble || pp.doubled {
							continue
						} else {
							pp.doubled = true
						}
					}
					pp.smalls.Add(dest)
				}
				next = append(next, pp)
			}
		}
		paths = next
		if len(paths) == 0 {
			break
		}
	}
	return total
}

func solve12A(input []string) int {
	return solve12(input, false)
}

func solve12B(input []string) int {
	return solve12(input, true)
}
