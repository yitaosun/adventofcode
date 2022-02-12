package main

import (
	"log"
	"strconv"
)

type window01B struct {
	buf   [3]int
	i     int
	count int
}

func (w *window01B) push(val int) {
	w.count++
	w.buf[w.i] = val
	w.i = (w.i + len(w.buf) + 1) % len(w.buf)
}

func (w *window01B) sum() int {
	var total int
	for _, v := range w.buf {
		total += v
	}
	return total
}

func solve01B(input []string) int {
	var last int
	var count int
	w := &window01B{}
	for _, line := range input {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		w.push(n)
		cur := w.sum()
		if w.count <= 3 {
			last = cur
			continue
		}
		if cur > last {
			count++
		}
		last = cur
	}
	return count
}
