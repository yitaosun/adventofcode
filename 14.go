package main

import (
	"math"
)

type poly14 struct {
	rules map[string]string
	memo  map[string]map[int]map[rune]int
}

func newPoly14(rules []string) *poly14 {
	p := &poly14{
		rules: map[string]string{},
		memo:  map[string]map[int]map[rune]int{},
	}
	for _, rule := range rules {
		p.rules[rule[:2]] = rule[0:1] + rule[6:7] + rule[1:2]
	}
	return p
}

func (p *poly14) solveRule(left string, steps int) map[rune]int {
	if steps == 0 {
		return map[rune]int{}
	}
	stepsMap, ok := p.memo[left]
	if !ok {
		stepsMap = map[int]map[rune]int{}
		p.memo[left] = stepsMap
	}
	rv, ok := stepsMap[steps]
	if ok {
		return rv
	}
	rv = map[rune]int{}
	if right, ok := p.rules[left]; ok {
		rv[rune(right[1])]++
		for r, c := range p.solveRule(right[:2], steps-1) {
			rv[r] += c
		}
		for r, c := range p.solveRule(right[1:], steps-1) {
			rv[r] += c
		}
	}
	stepsMap[steps] = rv
	return rv
}

func (p *poly14) solveString(input string, steps int) map[rune]int {
	rv := map[rune]int{}
	for _, r := range input {
		rv[r]++
	}
	if len(input) < 2 {
		return rv
	}
	for i := 0; i < len(input)-1; i++ {
		for r, c := range p.solveRule(input[i:i+2], steps) {
			rv[r] += c
		}
	}
	return rv
}

func solve14(input []string, steps int) int {
	p := newPoly14(input[2:])
	var most int
	var least int = math.MaxInt
	for _, c := range p.solveString(input[0], steps) {
		if c > most {
			most = c
		}
		if c < least {
			least = c
		}
	}
	return most - least
}

func solve14A(input []string) int {
	return solve14(input, 10)
}

func solve14B(input []string) int {
	return solve14(input, 40)

}
