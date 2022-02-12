package main

import (
	"fmt"
	"log"
)

type sub02B struct {
	x   int
	y   int
	aim int
}

func (s *sub02B) forward(steps int) {
	s.x += steps
	s.y += s.aim * steps
	if s.y < 0 {
		s.y = 0
	}
}

func (s *sub02B) up(steps int) {
	s.aim -= steps
}

func (s *sub02B) down(steps int) {
	s.aim += steps
}

func solve02B(input []string) int {
	s := &sub02B{}
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
