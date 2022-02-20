package main

import (
	"fmt"
	"math"
)

type map15 struct {
	risks [][]int
	m     int
	w     int
}

const (
	up15 = iota
	right15
	down15
	left15
)

type edge [3]int

func makeEdges(n [2]int) []edge {
	return []edge{
		{n[0], n[1], up15},
		{n[0], n[1], right15},
		{n[0], n[1], down15},
		{n[0], n[1], left15},
	}
}

func (e edge) src() [2]int {
	return [2]int{e[0], e[1]}
}

func (e edge) dst() [2]int {
	switch e[2] {
	case up15:
		return [2]int{e[0], e[1] - 1}
	case right15:
		return [2]int{e[0] + 1, e[1]}
	case down15:
		return [2]int{e[0], e[1] + 1}
	case left15:
		return [2]int{e[0] - 1, e[1]}
	}
	panic("Ahhhhh")
}

func (e edge) String() string {
	var dir string
	switch e[2] {
	case up15:
		dir = "up"
	case right15:
		dir = "right"
	case down15:
		dir = "down"
	case left15:
		dir = "left"
	}
	return fmt.Sprintf("[%d %d %s]", e[0], e[1], dir)
}

func newMap15(input []string, m int) *map15 {
	var risks [][]int
	for _, line := range input {
		var row []int
		for _, r := range line {
			row = append(row, int(r-'0'))
		}
		risks = append(risks, row)
	}
	return &map15{
		risks: risks,
		w:     len(risks) * m,
		m:     m,
	}
}

func (m *map15) riskAt(c [2]int) int {
	w := len(m.risks)
	risk := m.risks[c[1]%w][c[0]%w] + c[0]/w + c[1]/w
	return (risk-1)%9 + 1
}

func (m *map15) width() int {
	return len(m.risks) * m.m
}

func min(args ...int) int {
	rv := math.MaxInt
	for _, v := range args {
		if v < rv {
			rv = v
		}
	}
	return rv
}

func (m *map15) debug(paths map[[2]int]int) {
	for y := 0; y < m.w; y++ {
		for x := 0; x < m.w; x++ {
			r, ok := paths[[2]int{x, y}]
			if !ok {
				fmt.Print("  *")
			} else {
				fmt.Printf(" %2d", r)
			}
		}
		fmt.Println()
	}
}

func (m *map15) solve() int {
	paths := map[[2]int]int{
		{0, 0}: 0,
	}
	edges := NewSet(
		edge{0, 0, right15},
		edge{0, 0, down15},
	)
	for {
		newEdges := NewSet()
		var bestEdge edge
		var bestRisk int = math.MaxInt
		for el := range edges {
			e := el.(edge)
			if _, ok := paths[e.dst()]; ok {
				continue
			}
			risk := paths[e.src()] + m.riskAt(e.dst())
			if risk < bestRisk {
				bestEdge = e
				bestRisk = risk
			}
			newEdges.Add(e)
		}
		newEdges.Remove(bestEdge)
		n := bestEdge.dst()
		paths[n] = bestRisk
		if rv, ok := paths[[2]int{m.w - 1, m.w - 1}]; ok {
			return rv
		}
		for _, e := range makeEdges(n) {
			dst := e.dst()
			if _, ok := paths[dst]; !ok && dst[0] >= 0 && dst[0] < m.w && dst[1] >= 0 && dst[1] < m.w {
				newEdges.Add(e)
			}
		}
		edges = newEdges
		if len(paths)%1000 == 0 {
			fmt.Printf("Searched %d nodes. Edge set is %d\n", len(paths), len(edges))
		}
	}
}

func solve15A(input []string) int {
	return newMap15(input, 1).solve()
}

func solve15B(input []string) int {
	return newMap15(input, 5).solve()
}
