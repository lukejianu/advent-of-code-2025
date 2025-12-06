package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	lines := readLines()
	ansOne := doPartOne(lines)
	ansTwo := doPartTwo(lines)
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)

}

func doPartOne(lines []string) int {
	totalJoltage := 0
	for _, line := range lines {
		totalJoltage = totalJoltage + computeMaxJoltagePartOne(line)
	}
	return totalJoltage
}

func computeMaxJoltagePartOne(bank string) int {
	largest, idx := -1, -1
	for i, ch := range bank {
		if i == len(bank) - 1 {
			break // Ignore the last digit.
		}
		n := int(ch - '0')
		if n > largest {
			largest = n
			idx = i
		}
	}
	
	largestAfterLargest := -1
	for j := idx + 1; j < len(bank); j++ {
		n := int(bank[j] - '0')
		if n > largestAfterLargest {
			largestAfterLargest = n
		}
	}

	return largest * 10 + largestAfterLargest
}

func doPartTwo(lines []string) int {
	totalJoltage := 0
	for _, line := range lines {
		totalJoltage = totalJoltage + computeMaxJoltagePartTwo(line)
	}
	return totalJoltage
}

// Computes the max joltage in a bank. Starts with the most significant
// digit (12), finds the best choice, then continues. 
func computeMaxJoltagePartTwo(bank string) int {
	bankSize := len(bank)	
	l := 0
	maxJoltage := 0
	for digit := 12; digit >= 1; digit-- {
		r := bankSize - digit
		largest, idx := getLargestOccurrence(bank, l, r)
		l = idx + 1
		maxJoltage = maxJoltage + (largest * powInt(10, digit - 1))
	}
	return maxJoltage
}

// Find the largest digit and its index in bank[l:r + 1]. If there are
// duplicates, choose the leftmost.
func getLargestOccurrence(bank string, l, r int) (int, int) {
	largest, idx := -1, -1
	for i := l; i <= r; i++ {
		n := int(bank[i] - '0')
		if n > largest {
			largest = n
			idx = i
		}
	}
	return largest, idx
}

func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
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


