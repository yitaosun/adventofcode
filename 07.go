package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func solve07(input []string, costF func(int) int) int {
	var max int
	pos := map[int]int{}
	for _, s := range strings.Split(input[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		pos[i]++
		if i > max {
			max = i
		}
	}
	min := math.MaxInt
	for i := 0; i < max; i++ {
		var cost int
		for k, v := range pos {
			if k > i {
				cost += costF(k-i) * v
			} else if k < i {
				cost += costF(i-k) * v
			}
		}
		if cost < min {
			min = cost
		}
	}
	return min
}

func solve07A(input []string) int {
	return solve07(input, func(n int) int { return n })
}

func solve07B(input []string) int {
	return solve07(input, func(n int) int { return (n*n + n) / 2 })
}
