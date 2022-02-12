package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type bingo04 struct {
	board   [5][5]int
	markers [5][5]bool
}

var spaces05A = regexp.MustCompile(`\s+`)

func newBingo04(input []string) (*bingo04, error) {
	b := &bingo04{}
	for i := 0; i < 5; i++ {
		row := spaces05A.Split(input[i], 5)
		for j, cell := range row {
			n, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
			b.board[i][j] = n
		}
	}
	return b, nil
}

func (b *bingo04) mark(v int) int {
	var marked bool
	var ii, jj int
	for i, row := range b.board {
		for j, cell := range row {
			if cell == v {
				b.markers[i][j] = true
				marked = true
				ii = i
				jj = j
				break
			}
		}
		if marked {
			break
		}
	}
	if !marked {
		return 0
	}
	colBingo := true
	rowBingo := true
	for i := 0; i < 5; i++ {
		colBingo = colBingo && b.markers[i][jj]
		rowBingo = rowBingo && b.markers[ii][i]
	}
	if !(colBingo || rowBingo) {
		return 0
	}
	var score int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.markers[i][j] {
				score += b.board[i][j]
			}
		}
	}
	return score * v
}

func solve04A(input []string) int {
	var boards []*bingo04
	for i := 0; i < len(input[2:]); i += 6 {
		b, err := newBingo04(input[i+2 : i+7])
		if err != nil {
			log.Fatal(err)
		}
		boards = append(boards, b)
	}

	for _, s := range strings.Split(input[0], ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		var score int
		for _, b := range boards {
			newScore := b.mark(n)
			if newScore > score {
				score = newScore
			}
		}
		if score > 0 {
			return score
		}
	}
	log.Fatal("unexpected end")
	return 0
}
