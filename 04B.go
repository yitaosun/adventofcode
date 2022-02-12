package main

import (
	"log"
	"strconv"
	"strings"
)

func solve04B(input []string) int {
	var boards []*bingo04
	for i := 0; i < len(input[2:]); i += 6 {
		b, err := newBingo04(input[i+2 : i+7])
		if err != nil {
			log.Fatal(err)
		}
		boards = append(boards, b)
	}

	bingos := make([]bool, len(boards))
	numBingos := 0
	for _, s := range strings.Split(input[0], ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		for i, b := range boards {
			score := b.mark(n)
			if !bingos[i] && score > 0 {
				bingos[i] = true
				numBingos++
				if numBingos == len(boards) {
					return score
				}
			}
		}
	}
	log.Fatal("unexpected end")
	return 0
}
