package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type pos = [2]int // (y, x)
type rect = [2]pos
type tMatrix = [][]byte
type rnge = [2]int // [low, high]

/*
luke@mbp 9 % cat input | time go run main.go
Part One: 4755064176
Part Two: 1613305596
go run main.go  80.15s user 164.47s system 68% cpu 5:54.74 total
*/

func main() {
	lines := readLines()
	positions := parseInput(lines)
	ansOne := doPartOne(positions)
	ansTwo := doPartTwo(positions)
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)
}

func doPartOne(positions []pos) int {
	n := len(positions)
	bestArea := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1, p2 := positions[i], positions[j]
			bestArea = max(bestArea, rectangleSize(rect{p1, p2}))
		}
	}
	return bestArea
}

func rectangleSize(r rect) int {
	p1, p2 := r[0], r[1]
	dx := int(math.Abs(float64(p1[0]-p2[0]))) + 1
	dy := int(math.Abs(float64(p1[1]-p2[1]))) + 1
	return dx * dy
}

func doPartTwo(positions []pos) int {
	minX, minY, maxX, maxY := findBounds(positions)
	positions = normalize(positions, minX, minY)
	numRows := (maxY - minY) + 1
	numCols := (maxX - minX) + 1
	matrix := make(tMatrix, numRows, numRows)
	for i := range numRows {
		matrix[i] = make([]byte, numCols, numCols)
	}

	// Color in shape.
	fillMatrix(positions, matrix)
	for r := range numRows {
		for c := range numCols {
			if matrix[r][c] == 2 {
				matrix[r][c] = 0
			} else {
				matrix[r][c] = 1
			}
		}
	}

	rowRangesPerCol, colRangesPerRow := computeRanges(positions, matrix)

	rectangles := allRectangles(positions)
	for _, rct := range rectangles {
		if isGood(rct, rowRangesPerCol, colRangesPerRow) {
			return rectangleSize(rct)
		}
	}

	return -1
}

func findBounds(positions []pos) (int, int, int, int) {
	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, -1, -1
	for _, p := range positions {
		r, c := p[0], p[1]
		minX = min(minX, c)
		maxX = max(maxX, c)
		minY = min(minY, r)
		maxY = max(maxY, r)
	}
	return minX, minY, maxX, maxY
}

func normalize(positions []pos, minX, minY int) []pos {
	newPositions := []pos{}
	for _, p := range positions {
		newPositions = append(newPositions, pos{p[0] - minY, p[1] - minX})
	}
	return newPositions
}

func fillMatrix(positions []pos, matrix tMatrix) {
	n := len(positions)
	for _, p := range positions {
		matrix[p[0]][p[1]] = 1
	}
	for i := 0; i < n-1; i++ {
		connectAdjacent(matrix, positions[i], positions[i+1])
	}
	connectAdjacent(matrix, positions[0], positions[n-1])
	floodFill(positions, matrix)
}

func connectAdjacent(matrix tMatrix, p1, p2 pos) {
	if p1[0] == p2[0] { // Same row.
		r := p1[0]
		s, e := min(p1[1], p2[1]), max(p1[1], p2[1])
		for c := s; c < e; c++ {
			matrix[r][c] = 1
		}
	} else if p1[1] == p2[1] { // Same col.
		c := p1[1]
		s, e := min(p1[0], p2[0]), max(p1[0], p2[0])
		for r := s; r < e; r++ {
			matrix[r][c] = 1
		}
	} else {
		panic("bad input")
	}
}

func floodFill(positions []pos, matrix tMatrix) {
	numRows := len(matrix)
	numCols := len(matrix[0])
	for r := range numRows {
		if matrix[r][0] == 0 {
			bfs(matrix, r, 0)
		}
		if matrix[r][numCols-1] == 0 {
			bfs(matrix, r, 0)
		}
	}
	for c := range numCols {
		if matrix[0][c] == 0 {
			bfs(matrix, 0, c)
		}
		if matrix[numRows-1][c] == 0 {
			bfs(matrix, numRows-1, c)
		}
	}
}

func bfs(m tMatrix, startR, startC int) {
	numRows := len(m)
	numCols := len(m[0])

	toVisit := make([]pos, 0, 200000)
	toVisit = append(toVisit, pos{startR, startC})
	for len(toVisit) != 0 {
		front := toVisit[0]
		r, c := front[0], front[1]
		toVisit = toVisit[1:]
		if r < 0 || c < 0 || r >= numRows || c >= numCols {
			continue
		}
		if m[r][c] != 0 { // Skip explored.
			continue
		}
		m[r][c] = 2
		toVisit = append(toVisit, pos{r + 1, c}, pos{r - 1, c}, pos{r, c + 1}, pos{r, c - 1})
	}
}

func computeRanges(lop []pos, m tMatrix) ([][]rnge, [][]rnge) {
	rowRangesPerCol, colRangesPerRow := initRanges(m)
	for _, p := range lop {
		r, c := p[0], p[1]
		rowRangesPerCol[c] = updateRowRanges(c, rowRangesPerCol[c], m)
		colRangesPerRow[r] = updateColRanges(r, colRangesPerRow[r], m)
	}
	return rowRangesPerCol, colRangesPerRow
}

func initRanges(m tMatrix) ([][]rnge, [][]rnge) {
	numCols := len(m[0])
	rowRangesPerCol := make([][]rnge, numCols, numCols)
	numRows := len(m)
	colRangesPerRow := make([][]rnge, numRows, numRows)
	return rowRangesPerCol, colRangesPerRow
}

func updateRowRanges(c int, lor []rnge, m tMatrix) []rnge {
	if len(lor) != 0 {
		return lor
	}
	numRows := len(m)
	for r := range numRows {
		if m[r][c] == 1 {
			if len(lor) == 0 || lor[len(lor)-1][1] != r-1 {
				lor = append(lor, rnge{r, r - 1})
			}
			aRnge := &lor[len(lor)-1]
			aRnge[1] += 1
		}
	}
	return lor
}

func updateColRanges(r int, lor []rnge, m tMatrix) []rnge {
	if len(lor) != 0 {
		return lor
	}
	numCols := len(m[0])
	for c := range numCols {
		if m[r][c] == 1 {
			if len(lor) == 0 || lor[len(lor)-1][1] != c-1 {
				lor = append(lor, rnge{c, c - 1})
			}
			aRnge := &lor[len(lor)-1]
			aRnge[1] += 1
		}
	}
	return lor
}

func allRectangles(positions []pos) []rect {
	n := len(positions)
	rectangles := []rect{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1, p2 := positions[i], positions[j]
			rectangles = append(rectangles, rect{p1, p2})
		}
	}
	slices.SortFunc(rectangles, func(r1, r2 rect) int {
		return -1 * cmp.Compare(rectangleSize(r1), rectangleSize(r2))
	})
	return rectangles
}

func isGood(rct rect, rowRangesPerCol, colRangesPerRow [][]rnge) bool {
	p0, p1 := rct[0], rct[1]

	rectLeftmostCol := min(p0[1], p1[1])
	rectRightmostCol := max(p0[1], p1[1])
	rectLeftmostRow := min(p0[0], p1[0])
	rectRightmostRow := max(p0[0], p1[0])

	rectColRange := rnge{rectLeftmostCol, rectRightmostCol}
	rectRowRange := rnge{rectLeftmostRow, rectRightmostRow}

	actualColRange := overlappingRange(
		computeActualRange(p0[1], colRangesPerRow[p0[0]]),
		computeActualRange(p1[1], colRangesPerRow[p1[0]]))

	actualRowRange := overlappingRange(
		computeActualRange(p0[0], rowRangesPerCol[p0[1]]),
		computeActualRange(p1[0], rowRangesPerCol[p1[1]]))

	if actualRowRange[0] == -1 || actualColRange[0] == -1 {
		return false
	}

	return rectRowRange[0] >= actualRowRange[0] &&
		rectRowRange[1] <= actualRowRange[1] &&
		rectColRange[0] >= actualColRange[0] &&
		rectColRange[1] <= actualColRange[1]
}

func computeActualRange(i int, lor []rnge) rnge {
	for _, r := range lor {
		// This is the only AI generated line of code in this program.
		// Such a noob bug...
		thankYouGemini := i <= r[1]
		if i >= r[0] && thankYouGemini {
			return r
		}
	}
	panic("not possible")
}

func overlappingRange(r1, r2 rnge) rnge {
	overlap := rnge{max(r1[0], r2[0]), min(r1[1], r2[1])}
	if overlap[0] > overlap[1] {
		return rnge{-1, -1}
	}
	return overlap
}

func printMatrix(matrix tMatrix) {
	for r := range len(matrix) {
		sRow := byteSliceToStringSlice(matrix[r])
		fmt.Println(strings.Join(sRow, ""))
	}
}

func byteSliceToStringSlice(nums []byte) []string {
	s := []string{}
	for _, n := range nums {
		s = append(s, strconv.Itoa(int(n-0)))
	}
	return s
}

func parseInput(lines []string) []pos {
	positions := []pos{}
	for _, line := range lines {
		s := strings.Split(line, ",")
		positions = append(positions, pos{
			stringToInt(s[1]),
			stringToInt(s[0]),
		})
	}
	return positions
}

func stringToInt(s string) int {
	n, err := strconv.Atoi(s)
	must(err)
	return n
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
