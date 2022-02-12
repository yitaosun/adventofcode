package main

import (
	"fmt"
	"log"
)

type sub02A struct {
	x int
	y int
}

func (s *sub02A) forward(steps int) {
	s.x += steps
}

func (s *sub02A) up(steps int) {
	s.y -= steps
	if s.y < 0 {
		s.y = 0
	}
}

func (s *sub02A) down(steps int) {
	s.y += steps
}

func solve02A(input []string) int {
	s := &sub02A{}
	for _, line := range input {
		var cmd string
		var steps int
		if _, err := fmt.Sscanf(line, "%s %d", &cmd, &steps); err != nil {
			log.Fatal(err)
		}
		switch cmd {
		case "forward":
			s.forward(steps)
		case "up":
			s.up(steps)
		case "down":
			s.down(steps)
		default:
			log.Fatal("invalid input:", line)
		}
	}
	return s.x * s.y

}
