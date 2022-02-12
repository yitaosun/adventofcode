package main

import (
	"fmt"
	"log"
)

type seafloor05 map[[2]int]int

func (s seafloor05) add(x0, y0, x1, y1 int, diagonal bool) {
	if x0 == x1 {
		if y0 > y1 {
			y0, y1 = y1, y0
		}
		for i := y0; i <= y1; i++ {
			s[[2]int{x0, i}]++
		}
	} else if y0 == y1 {
		if x0 > x1 {
			x0, x1 = x1, x0
		}
		for i := x0; i <= x1; i++ {
			s[[2]int{i, y0}]++
		}
	} else if diagonal {
		if x0 > x1 {
			x0, y0, x1, y1 = x1, y1, x0, y0
		}
		if x1-x0 == y1-y0 {
			for i := x1 - x0; i >= 0; i-- {
				s[[2]int{x0 + i, y0 + i}]++
			}
		} else if x1-x0 == y0-y1 {
			for i := x1 - x0; i >= 0; i-- {
				s[[2]int{x0 + i, y0 - i}]++
			}
		}
	}
}

func (s seafloor05) print() {
	var maxX, maxY int
	for k := range s {
		if k[0] > maxX {
			maxX = k[0]
		}
		if k[1] > maxY {
			maxY = k[1]
		}
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if v, ok := s[[2]int{x, y}]; ok {
				fmt.Print(v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func solve05(input []string, diagonal bool) int {
	s := seafloor05{}
	for _, line := range input {
		var x0, y0, x1, y1 int
		if _, err := fmt.Sscanf(line, "%d,%d -> %d,%d", &x0, &y0, &x1, &y1); err != nil {
			log.Fatal(err)
		}
		s.add(x0, y0, x1, y1, diagonal)
	}
	var count int
	for _, v := range s {
		if v >= 2 {
			count++
		}
	}
	return count
}

func solve05A(input []string) int {
	return solve05(input, false)
}

func solve05B(input []string) int {
	return solve05(input, true)
}
