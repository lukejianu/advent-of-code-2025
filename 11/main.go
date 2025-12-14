package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type graph map[string][]string

func main() {
	lines := readLines()
	ansOne := doPartOne(parseInput(lines))
	ansTwo := doPartTwo(parseInput(lines))
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)
}

func doPartOne(g graph) int {
	return dfs("you", g)
}

func dfs(node string, g graph) int {
	if node == "out" {
		return 1	
	}

	total := 0
	to := g[node]
	for _, t := range to {
		total += dfs(t, g)
	}
	return total
}

func doPartTwo(g graph) int {
	return dfsPartTwo("svr", g, false, false)
}

func dfsPartTwo(
	node string,
	g graph,
	seenDac,
	seenFft bool,
) int {
	if node == "out" && seenDac && seenFft {
		return 1	
	}

	if node == "out" {
		return 0
	}

	if node == "dac" {
		seenDac = true
	}

	if node == "fft" {
		seenFft = true
	}

	total := 0
	to := g[node]
	for _, t := range to {
		total += dfsPartTwo(t, g, seenDac, seenFft)
	}

	return total
}

func readLines() []string {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseInput(lines []string) graph {
	g := graph{}
	for _, line := range lines {
		from, to := parseLine(line)
		g[from] = to
	}
	return g
}

func parseLine(line string) (string, []string) {
	toks := strings.Split(line, " ")
	from := toks[0][:len(toks[0]) - 1]
	to := toks[1:]
	return from, to
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
