package main

import (
	"log"
	"strconv"
	"strings"
)

func solve06(input []string, gen int) int {
	fishes := map[int]int{}
	for _, s := range strings.Split(input[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		fishes[i]++
	}
	for t := 0; t < gen; t++ {
		newFishes := map[int]int{}
		spawns := fishes[0]
		for i := 0; i < 8; i++ {
			newFishes[i] = fishes[i+1]
		}
		newFishes[8] = spawns
		newFishes[6] += spawns
		fishes = newFishes
	}
	var total int
	for _, v := range fishes {
		total += v
	}
	return int(total)
}

func solve06A(input []string) int {
	return solve06(input, 80)
}

func solve06B(input []string) int {
	return solve06(input, 256)
}
