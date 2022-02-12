package main

import (
	"log"
	"strings"
)

func solve03A(input []string) int {
	var count int
	var bits []int
	for _, line := range input {
		line := strings.TrimSpace(line)
		n := len(line)
		if len(bits) < n {
			newBits := make([]int, n)
			copy(newBits[:len(bits)], bits)
			bits = newBits
		}
		for i, c := range line {
			switch c {
			case '0':
				// Do nothing on 0s
			case '1':
				bits[i]++
			default:
				log.Fatal("invalid input:", line)
			}
		}
		count++
	}
	var gamma int64
	var epsilon int64
	for _, c := range bits {
		if c > count/2 {
			gamma = gamma | 1
		} else {
			epsilon = epsilon | 1
		}
		gamma = gamma << 1
		epsilon = epsilon << 1
	}
	gamma = gamma >> 1
	epsilon = epsilon >> 1
	return int(gamma * epsilon)
}
