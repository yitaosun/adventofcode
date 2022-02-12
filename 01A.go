package main

import (
	"log"
	"strconv"
)

func solve01A(input []string) int {
	var first bool = true
	var last int
	var count int
	for _, line := range input {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if first {
			first = false
			last = n
			continue
		}
		if n > last {
			count++
		}
		last = n
	}
	return count
}
