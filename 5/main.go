package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type idRange = [2]int

func main() {
	lines := readLines()
	ranges, ingredients := parseInput(lines)
	ansOne := doPartOne(ranges, ingredients)
	ansTwo := doPartTwo(ranges, ingredients)
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)

}

func doPartOne(idRanges []idRange, ids []int) int {
	numFresh := 0
	for _, id := range ids {
		if isFresh(id, idRanges) {
			numFresh += 1
		}
	}
	return numFresh
}

func isFresh(id int, idRanges []idRange) bool {
	for _, rnge := range idRanges {
		l, r := rnge[0], rnge[1]
		if id >= l && id <= r {
			return true
		}
	}
	return false
}

func doPartTwo(idRanges []idRange, _ []int) int {
	merged := mergeRanges(idRanges)
	return sumRanges(merged)
}

func mergeRanges(idRanges []idRange) []idRange {
	slices.SortFunc(idRanges, func(rnge1, rnge2 idRange) int {
		return cmp.Compare(rnge1[1], rnge2[1])
	})
	merged := []idRange{idRanges[0]}
	for _, rnge := range idRanges[1:] {
		top := &merged[len(merged)-1]
		if isOverlap(top, &rnge) {
			top[0] = min(top[0], rnge[0])
			top[1] = rnge[1]
		} else {
			merged = append(merged, rnge)
		}
	}
	return merged
}

func isOverlap(rnge1, rnge2 *idRange) bool {
	return rnge2[0] <= rnge1[1]
}

func sumRanges(idRanges []idRange) int {
	sum := 0
	for _, rnge := range idRanges {
		l, r := rnge[0], rnge[1]
		sum += (r - l + 1)
	}
	return sum
}

func parseInput(lines []string) ([]idRange, []int) {
	parseIds := false
	ranges := []idRange{}
	ids := []int{}
	for _, line := range lines {
		if line == "" {
			parseIds = true
			continue
		}
		if !parseIds {
			reg := regexp.MustCompile("^(\\d+)-(\\d+)$")
			matches := reg.FindStringSubmatch(line)
			l, e1 := strconv.Atoi(matches[1])
			r, e2 := strconv.Atoi(matches[2])
			must(e1)
			must(e2)
			ranges = append(ranges, idRange{l, r})
		} else {
			id, err := strconv.Atoi(line)
			must(err)
			ids = append(ids, id)
		}
		if line == "" {
			parseIds = true
		}
	}
	return ranges, ids
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
