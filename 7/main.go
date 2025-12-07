package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos = [2]int

func main() {
	lines := readLines()
	ansOne := doPartOne(lines)
	ansTwo := doPartTwo(lines)
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)
}

func doPartOne(grid [][]byte) int {
	numRows := len(grid)
	numCols := len(grid[0])
	splits := 0
	for r := 1; r < numRows; r++ {
		for c := range numCols {
			e := grid[r][c]
			aboveE := grid[r - 1][c]
			switch e {
			case '.':
				if aboveE == '|' || aboveE == 'S' {
					grid[r][c] = '|'
				}
			case '^':
				if aboveE == '|' {
					grid[r][c - 1] = '|'
					grid[r][c + 1] = '|'
					splits += 1
				}
			default:
			}
		}
	}
	return splits
}

func doPartTwo(grid [][]byte) int {
	cache := map[pos]int{}
	for c, e := range grid[0] {
		if e == 'S' {
			return dfs(grid, cache, 1, c)
		}
	}
	panic("start not found")
}

// Counts how many paths there are from point (r, c).
func dfs(grid [][]byte, cache map[pos]int, r, c int) int {
	numRows := len(grid)
	if r >= numRows {
		return 1
	}
	if v, ok := cache[pos{r, c}]; ok {
		return v	
	}
	e := grid[r][c]
	switch e {
	case '|':
		// Need to handle pipe case because of P1 mutation.
		return dfs(grid, cache, r + 1, c)
	case '.':
		return dfs(grid, cache, r + 1, c)
	case '^':
		left := dfs(grid, cache, r, c - 1)
		right := dfs(grid, cache, r, c + 1) // DO NOT CALL THIS `r` (lol).
		res := left + right
		cache[pos{r, c}] = res
		return res
	default:
		panic("bad e")
	}
}

func readLines() [][]byte {
	lines := [][]byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, []byte(scanner.Text()))
	}
	return lines
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
