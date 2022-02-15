package main

import (
	"fmt"
	"log"
	"strconv"
)

type points13 struct {
	xs []int
	ys []int
	w  int
	h  int
}

func newPoints13(input []string) *points13 {
	p := &points13{}
	for _, line := range input {
		var x, y int
		if _, err := fmt.Sscanf(line, "%d,%d", &x, &y); err != nil {
			log.Fatal(err)
		}
		p.add(x, y)
	}
	return p
}

func (p *points13) add(x, y int) {
	p.xs = append(p.xs, x)
	p.ys = append(p.ys, y)
	if x+1 > p.w {
		p.w = x + 1
	}
	if y+1 > p.h {
		p.h = y + 1
	}
}

func fold(vals []int, pivot, max int) []int {
	rv := []int{}
	if pivot >= max/2 {
		for _, v := range vals {
			if v > pivot {
				v = 2*pivot - v
			}
			rv = append(rv, v)
		}
	} else {
		for _, v := range vals {
			if v < pivot {
				v = 2*pivot - v
			} else {
				v -= pivot
			}
			rv = append(rv, v)
		}
	}
	return rv
}

func (p *points13) foldX(x int) {
	p.xs = fold(p.xs, x, p.w)
	if x >= p.w/2 {
		p.w = x
	} else {
		p.w -= x
	}
}

func (p *points13) foldY(y int) {
	p.ys = fold(p.ys, y, p.h)
	if y >= p.h/2 {
		p.h = y
	} else {
		p.h -= y
	}
}

func (p *points13) toMap() map[[2]int]struct{} {
	m := map[[2]int]struct{}{}
	for i := 0; i < len(p.xs); i++ {
		m[[2]int{p.xs[i], p.ys[i]}] = struct{}{}
	}
	return m
}

func (p *points13) print() {
	m := p.toMap()
	for y := 0; y < p.h; y++ {
		for x := 0; x < p.w; x++ {
			if _, ok := m[[2]int{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func solve13A(input []string) int {
	var points []string
	var instructions []string
	for i, line := range input {
		if line == "" {
			points = input[:i]
			instructions = input[i+1:]
			break
		}
	}
	p := newPoints13(points)
	for _, inst := range instructions[:1] {
		v, err := strconv.Atoi(inst[13:])
		if err != nil {
			log.Fatal(err)
		}
		switch inst[11] {
		case 'x':
			p.foldX(v)
		case 'y':
			p.foldY(v)
		}
	}
	return len(p.toMap())
}

func solve13B(input []string) int {
	var points []string
	var instructions []string
	for i, line := range input {
		if line == "" {
			points = input[:i]
			instructions = input[i+1:]
			break
		}
	}
	p := newPoints13(points)
	for _, inst := range instructions {
		v, err := strconv.Atoi(inst[13:])
		if err != nil {
			log.Fatal(err)
		}
		switch inst[11] {
		case 'x':
			p.foldX(v)
		case 'y':
			p.foldY(v)
		}
	}
	p.print()
	return 0
}
