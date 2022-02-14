package main

import "sort"

func solve09A(input []string) int {
	var hmap [][]int
	for _, line := range input {
		row := make([]int, len(line))
		for i, r := range line {
			row[i] = int(r - '0')
		}
		hmap = append(hmap, row)
	}
	var risk int
	for y, row := range hmap {
		for x := range row {
			h := hmap[y][x]
			if x > 0 && h >= hmap[y][x-1] ||
				x < len(row)-1 && h >= hmap[y][x+1] ||
				y > 0 && h >= hmap[y-1][x] ||
				y < len(hmap)-1 && h >= hmap[y+1][x] {
				continue
			}
			risk += h + 1
		}
	}
	return risk
}

func searchBasin(hmap [][]int, basins [][]*int, basin *int, y, x int) {
	(*basin) += 1
	basins[y][x] = basin
	if x > 0 && hmap[y][x-1] < 9 && basins[y][x-1] == nil {
		searchBasin(hmap, basins, basin, y, x-1)
	}
	if x < len(hmap[y])-1 && hmap[y][x+1] < 9 && basins[y][x+1] == nil {
		searchBasin(hmap, basins, basin, y, x+1)
	}
	if y > 0 && hmap[y-1][x] < 9 && basins[y-1][x] == nil {
		searchBasin(hmap, basins, basin, y-1, x)
	}
	if y < len(hmap)-1 && hmap[y+1][x] < 9 && basins[y+1][x] == nil {
		searchBasin(hmap, basins, basin, y+1, x)
	}
}

func solve09B(input []string) int {
	var hmap [][]int
	var basins [][]*int
	var sizes []int
	for _, line := range input {
		row := make([]int, len(line))
		for i, r := range line {
			row[i] = int(r - '0')
		}
		hmap = append(hmap, row)
		basins = append(basins, make([]*int, len(line)))
	}
	for y, row := range hmap {
		for x := range row {
			if hmap[y][x] < 9 && basins[y][x] == nil {
				var basin int
				searchBasin(hmap, basins, &basin, y, x)
				sizes = append(sizes, basin)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}
