package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type idRange = [2]int64

func main() {
	line := readFirstLine()
	ansOne := doPartOne(line)
	ansTwo := doPartTwo(line)
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)
}

func doPartOne(line string) int64 {
	ranges := parseRanges(line)
	invalidIdSum := int64(0)
	for _, rnge := range ranges {
		l := rnge[0]
		r := rnge[1]
		for i := l; i <= r; i++ {
			s := strconv.FormatInt(i, 10)
			n := len(s)
			half := (n / 2)
			if n % 2 == 0 && s[:half] == s[half:] {
				invalidIdSum = invalidIdSum + i
			}
		}
	}
	return invalidIdSum
}

func doPartTwo(line string) int64 {
	ranges := parseRanges(line)
	invalidIdSum := int64(0)
	for _, rnge := range ranges {
		l := rnge[0]
		r := rnge[1]
		for i := l; i <= r; i++ {
			s := strconv.FormatInt(i, 10)
			if isInvalidId(s) {
				invalidIdSum += i
			}
		}
	}
	return invalidIdSum
}

func isInvalidId(id string) bool {
	n := len(id)
	for j := 1; j <= n / 2; j++ {
		maybeRepeated := id[:j]
		if isComposedOfRepeated(id, maybeRepeated) {
			return true
		}
	}
	return false
}

// Check if `s` split by `repeated` yields a perfect split.
func isComposedOfRepeated(s, repeated string) bool {
	splits := strings.Split(s, repeated)
	hasNonEmptyElement := slices.ContainsFunc(splits, func(split string) bool {
		return split != ""
	})
	hasOnlyEmptyElements := !hasNonEmptyElement
	return hasOnlyEmptyElements
}

func parseRanges(line string) []idRange {
	ranges := []idRange{}
	for _, rnge := range strings.Split(line, `,`) {
		reg := regexp.MustCompile("^(\\d+)-(\\d+)$")
		matches := reg.FindStringSubmatch(rnge)
		l, e1 := strconv.ParseInt(matches[1], 0, 64)
		must(e1)
		r, e2 := strconv.ParseInt(matches[2], 0, 64)
		must(e2)
		ranges = append(ranges, idRange{int64(l), int64(r)})
	}
	return ranges
}

func readFirstLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return scanner.Text()
	}
	panic("no line")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

