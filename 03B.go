package main

import (
	"log"
	"strconv"
)

func reduce03B(inputs []string, high bool) string {
	var i int
	for {
		var ones []string
		var zeros []string
		for _, v := range inputs {
			if len(v) <= i {
				continue
			}
			switch v[i] {
			case '0':
				zeros = append(zeros, v)
			case '1':
				ones = append(ones, v)
			}
		}
		if high {
			inputs = ones
			if len(ones) < len(zeros) {
				inputs = zeros
			}
		} else {
			inputs = zeros
			if len(ones) < len(zeros) {
				inputs = ones
			}
		}
		if len(inputs) == 1 {
			return inputs[0]
		}
		i++
	}
}

func solve03B(input []string) int {
	gamma, err := strconv.ParseInt(reduce03B(input, true), 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(reduce03B(input, false), 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(gamma * epsilon)
}
