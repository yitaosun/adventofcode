package main

import (
	"sort"
)

func checker10(s string) (int, int) {
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	rev := map[rune]rune{
		'(': ')', '[': ']', '{': '}', '<': '>',
	}
	var stack []rune
	for _, r := range s {
		switch r {
		case '(', '[', '{', '<':
			stack = append(stack, rev[r])
		case ')', ']', '}', '>':
			if stack[len(stack)-1] != r {
				return points[r], 0
			}
			stack = stack[:len(stack)-1]
		}
	}
	points = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	var score int
	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		score += points[stack[i]]
	}
	return 0, score
}

func solve10A(input []string) int {
	var total int
	for _, line := range input {
		score, _ := checker10(line)
		total += score
	}
	return total
}

func solve10B(input []string) int {
	var scores []int
	for _, line := range input {
		_, score := checker10(line)
		if score > 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
