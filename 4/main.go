package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos = [2]int

func main() {
	lines := readLines()
	grid := parseGrid(lines)
	ansOne := doPartOne(grid)
	ansTwo := doPartTwo(grid)
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)

}


func doPartOne(grid [][]byte) int {
	numRows := len(grid)
	numCols := len(grid[0])
	rolls := 0
	for r := range numRows {
		for c := range numCols {
			if grid[r][c] == '@' && isAccessible(grid, r, c) {
				rolls += 1	
			}
		}
	}
	return rolls
}

func doPartTwo(grid [][]byte) int {
	numRows := len(grid)
	numCols := len(grid[0])
	rolls := 0
	for {
		madeProgress := false
		for r := range numRows {
			for c := range numCols {
				if grid[r][c] == '@' && isAccessible(grid, r, c) {
					rolls += 1	
					grid[r][c] = '.'
					madeProgress = true
				}
			}
		}
		if !madeProgress {
			break
		}
	}
	return rolls
}

func parseGrid(grid []string) [][]byte {
	newGrid := [][]byte{}
	for _, row := range grid {
		newGrid = append(newGrid, []byte(row))
	}
	return newGrid
}

func isAccessible(grid [][]byte, r, c int) bool {
	positions := adjacentValidPositions(grid, r, c)
	numRolls := 0
	for _, pos := range positions {
		rPos, cPos := pos[0], pos[1]
		if grid[rPos][cPos] == '@' {
			numRolls += 1
		}
	}
	return numRolls < 4
}

func adjacentValidPositions(grid [][]byte, r, c int) []pos {
	numRows := len(grid)
	numCols := len(grid[0])
	directions := [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
	}
	positions := []pos{}
	for _, d := range directions {
		dr, dc := d[0], d[1]
		nr, nc := r + dr, c + dc 
		if nr >= 0 && nc >= 0 && nr < numRows && nc < numCols {
			positions = append(positions, pos{nr, nc})
		}
	}
	return positions
}

func readLines() []string {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}


