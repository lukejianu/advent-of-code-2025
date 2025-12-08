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

type pos = [3]int
type edge = [2]int

// NOTE: This program will fail on the sample small input.
func main() {
	lines := readLines()
	positions := parseInput(lines)
	ansOne := doPartOne(positions)
	ansTwo := doPartTwo(positions)
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)

}

func doPartOne(positions []pos) int {
	id := map[int]pos{}
	for i, p := range positions {
		id[i] = p
	}

	matrix := make([][]float64, len(positions))
	for i := range positions {
		matrix[i] = make([]float64, len(positions))
	}

	for i := range positions {
		for j := range positions {
			p1 := id[i]
			p2 := id[j]
			d := distance(p1, p2)
			matrix[i][j] = d
		}
	}

	edges := []edge{}
	for i := range positions {
		for j := i + 1; j < len(positions); j++ {
			edges = append(edges, edge{i, j})
		}
	}

	slices.SortFunc(edges, func(e1, e2 edge) int {
		e1Weight := matrix[e1[0]][e1[1]]
		e2Weight := matrix[e2[0]][e2[1]]
		return cmp.Compare(e1Weight, e2Weight)
	})

	ufd := map[int]int{}
	for i := range positions {
		ufd[i] = i
	}

	for i := range 1000 {
		e := edges[i]
		addEdge(ufd, e)
	}

	sizes := map[int]int{}
	for _, p := range ufd {
		if _, ok := sizes[p]; !ok {
			numChildren := findChildren(ufd, p)
			sizes[p] = numChildren
		}
	}

	values := []int{}
	for _, size := range sizes {
		values = append(values, -size)
	}

	slices.Sort(values)

	return -1 * values[0] * values[1] * values[2]
}

func doPartTwo(positions []pos) int {
	id := map[int]pos{}
	for i, p := range positions {
		id[i] = p
	}

	matrix := make([][]float64, len(positions))
	for i := range positions {
		matrix[i] = make([]float64, len(positions))
	}

	for i := range positions {
		for j := range positions {
			p1 := id[i]
			p2 := id[j]
			d := distance(p1, p2)
			matrix[i][j] = d
		}
	}

	edges := []edge{}
	for i := range positions {
		for j := i + 1; j < len(positions); j++ {
			edges = append(edges, edge{i, j})
		}
	}

	slices.SortFunc(edges, func(e1, e2 edge) int {
		e1Weight := matrix[e1[0]][e1[1]]
		e2Weight := matrix[e2[0]][e2[1]]
		return cmp.Compare(e1Weight, e2Weight)
	})

	ufd := map[int]int{}
	for i := range positions {
		ufd[i] = i
	}

	ans := -1
	i := 0
	for !isDone(ufd) {
		e := edges[i]
		addEdge(ufd, e)
		ans = id[e[0]][0] * id[e[1]][0]
		i += 1
	}

	return ans
}

// Check if the graph is fully connected.
func isDone(ufd map[int]int) bool {
	checkP := ufd[0]
	for _, p := range ufd {
		if p != checkP {
			return false
		}
	}
	return true
}

// Counts the number of nodes connected to the given parent node.
func findChildren(ufd map[int]int, theP int) int {
	n := 0
	for _, p := range ufd {
		if p == theP {
			n += 1
		}
	}
	return n
}

// Add edge to UF data structure. Performs slow path compression.
func addEdge(ufd map[int]int, e edge) {
	a, b := e[0], e[1]
	aParent := ufd[a]
	bParent := ufd[b]
	for id, p := range ufd {
		if p == bParent {
			ufd[id] = aParent
		}
	}
}

func distance(p1, p2 pos) float64 {
	x1, y1, z1 := p1[0], p1[1], p1[2]
	x2, y2, z2 := p2[0], p2[1], p2[2]

	dz := z2 - z1

	n := math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))

	return math.Sqrt(n*n + float64(dz*dz))
}

func parseInput(lines []string) []pos {
	positions := []pos{}
	for _, line := range lines {
		s := strings.Split(line, ",")
		positions = append(positions, pos{
			stringToInt(s[0]),
			stringToInt(s[1]),
			stringToInt(s[2]),
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
