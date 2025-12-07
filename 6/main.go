package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readLines()
	valueRows, ops := parseInput(lines)
	ansOne := doPartOne(valueRows, ops)
	ansTwo := doPartTwo(lines)
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)

}

func doPartOne(valueRows [][]int, ops []string) int {
	numRows := len(valueRows)
	numCols := len(valueRows[0])
	sum := 0
	for c := range numCols {
		op := ops[c]
		fullCol := []int{}
		for r := range numRows {
			fullCol = append(fullCol, valueRows[r][c])
		}
		v := compute(op, fullCol)
		sum += v
	}
	return sum
}

func doPartTwo(lines []string) int {
	total := 0
	ops := parseRow(lines[len(lines)-1])
	lines = lines[:len(lines)-1]
	values := parsePartTwoInput(lines)
	for i, op := range ops {
		total += compute(op, values[i])
	}
	return total
}

func compute(op string, values []int) int {
	res := 0
	if op == "*" {
		res = 1
	}
	for _, v := range values {
		switch op {
		case "*":
			res *= v
		case "+":
			res += v
		default:
			panic("bad op")
		}
	}
	return res
}

func parsePartTwoInput(lines []string) [][]int {
	numRows := len(lines)
	numCols := len(lines[0])
	valueGroups := [][]int{[]int{}}
	for c := range numCols {
		num := []byte{}
		// Attempt to read a number down the column.
		for r := range numRows {
			e := lines[r][c]
			if e != ' ' {
				num = append(num, e)
			}
		}
		// Empty column, start a new group.
		if len(num) == 0 {
			valueGroups = append(valueGroups, []int{})
		} else {
			s := string(num)
			n, err := strconv.Atoi(s)
			must(err)
			values := valueGroups[len(valueGroups)-1]
			values = append(values, n)
			valueGroups[len(valueGroups)-1] = values
		}
	}
	return valueGroups
}

func parseInput(lines []string) ([][]int, []string) {
	valueRows := [][]int{}
	for _, line := range lines[:len(lines)-1] {
		valueRows = append(valueRows, toIntSlice(parseRow(line)))
	}
	ops := parseRow(lines[len(lines)-1])
	return valueRows, ops
}

func parseRow(line string) []string {
	row := strings.Split(line, " ")
	row = slices.DeleteFunc(row, func(s string) bool {
		return s == ""
	})
	return row
}

func toIntSlice(stringSlice []string) []int {
	intSlice := make([]int, 0, len(stringSlice))
	for _, s := range stringSlice {
		n, err := strconv.Atoi(s)
		must(err)
		intSlice = append(intSlice, n)
	}
	return intSlice
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
