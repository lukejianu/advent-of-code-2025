package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type rotation struct {
	direction direction
	numClicks uint
}

type direction int

const (
	left direction = iota
	right
)

func main() {
	lines := readLines()
	ansOne := doPartOne(lines)
	ansTwo := doPartTwo(lines)
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)
}


func doPartOne(lines []string) int {
	rotations := parseLines(lines)
	dial := 50
	timesAtZero := 0
	for _, rotation := range rotations {
		step := computeStep(rotation)
		for i := uint(0); i < rotation.numClicks; i++ {
			dial = dial + step
			if dial == -1 {
				dial = 99
			} else if dial == 100 {
				dial = 0
			}
		}
		if dial == 0 {
			timesAtZero++
		}

	}
	return timesAtZero
}

func doPartTwo(lines []string) int {
	rotations := parseLines(lines)
	dial := 50
	timesAtZero := 0
	for _, rotation := range rotations {
		step := computeStep(rotation)
		for i := uint(0); i < rotation.numClicks; i++ {
			dial = dial + step
			if dial == -1 {
				dial = 99
			} else if dial == 100 {
				dial = 0
			}
			if dial == 0 {
				timesAtZero++
			}
		}

	}
	return timesAtZero
}

func computeStep(r rotation) int {
	if r.direction == left {
		return -1
	} else if r.direction == right {
		return 1
	} else {
		panic("bad rotation")
	}
}

func parseLines(lines []string) []rotation {
	rotations := []rotation{}
	r := regexp.MustCompile("^([LR])(\\d+)$")
	for _, line := range lines {
		matches := r.FindStringSubmatch(line)
		numClicks, err := strconv.ParseUint(matches[2], 0, 32)
		must(err)
		rotations = append(rotations, rotation{
			direction: parseDirection(matches[1]),
			numClicks: uint(numClicks),
		})
	}
	return rotations
}

func parseDirection(d string) direction {
	if d == "L" {
		return left
	} else if d == "R" {
		return right
	} else {
		panic("bad direction")
	}
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

