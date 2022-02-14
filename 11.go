package main

import "fmt"

func print11(op [][]int) {
	for _, row := range op {
		for _, r := range row {
			if r >= 10 {
				fmt.Print("*")
			} else {
				fmt.Print(r)
			}
		}
		fmt.Println()
	}
}

func update11(op [][]int) int {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			increment11(op, i, j)
		}
	}
	var total int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if op[i][j] >= 10 {
				total++
				op[i][j] = 0
			}
		}
	}
	return total
}

func increment11(op [][]int, i, j int) {
	op[i][j]++
	if op[i][j] == 10 {
		// NW
		if i > 0 && j > 0 {
			increment11(op, i-1, j-1)
		}
		// W
		if j > 0 {
			increment11(op, i, j-1)
		}
		// SW
		if i < 9 && j > 0 {
			increment11(op, i+1, j-1)
		}
		// S
		if i < 9 {
			increment11(op, i+1, j)
		}
		// SE
		if i < 9 && j < 9 {
			increment11(op, i+1, j+1)
		}
		// E
		if j < 9 {
			increment11(op, i, j+1)
		}
		// NE
		if i > 0 && j < 9 {
			increment11(op, i-1, j+1)
		}
		// N
		if i > 0 {
			increment11(op, i-1, j)
		}
	}
}

func solve11A(input []string) int {
	var op [][]int
	for _, line := range input {
		row := make([]int, 10)
		for i, r := range line {
			row[i] = int(r - '0')
		}
		op = append(op, row)
	}
	var total int
	for i := 0; i < 100; i++ {
		total += update11(op)
	}
	return total
}

func solve11B(input []string) int {
	var op [][]int
	for _, line := range input {
		row := make([]int, 10)
		for i, r := range line {
			row[i] = int(r - '0')
		}
		op = append(op, row)
	}
	step := 1
	for {
		if update11(op) == 100 {
			return step
		}
		step++
	}
}
