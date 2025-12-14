package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type machineSpec struct {
	lights   string
	buttons  [][]int
	joltages []int
}

func main() {
	lines := readLines()
	ansOne := doPartOne(parseInput(lines))
	ansTwo := doPartTwo(parseInput(lines))
	fmt.Printf("Part One: %d\n", ansOne)
	fmt.Printf("Part Two: %d\n", ansTwo)
}

func doPartOne(loms []machineSpec) int {
	total := 0
	for _, ms := range loms {
		cache := map[string]int{}
		pressed := make([]byte, len(ms.buttons))
		for i := range pressed {
			pressed[i] = '0'
		}
		start := strings.Repeat(".", len(ms.lights))
		fewest := fewestPressesPartOne(start, ms.lights, ms.buttons, pressed, cache)
		total += fewest
	}
	return total
}

// Computes the number of presses needed to reach the target state
// from the current state.
func fewestPressesPartOne(
	lights, target string,
	buttons [][]int,
	pressed []byte,
	cache map[string]int,
) int {
	if lights == target {
		return 0
	}

	cacheKey := lights + "," + string(pressed)

	if e, ok := cache[cacheKey]; ok {
		return e
	}

	fewest := len(buttons) * 2
	for i, b := range buttons {
		if pressed[i] == '0' {
			pressed[i] = '1'
			newLights := press(lights, b)
			candidate := fewestPressesPartOne(newLights, target, buttons, pressed, cache)
			fewest = min(1+candidate, fewest)
			pressed[i] = '0'
		}
	}

	cache[cacheKey] = fewest

	return fewest
}

func press(lights string, b []int) string {
	lightsSlice := strings.Split(lights, "")
	for _, n := range b {
		if lightsSlice[n] == "." {
			lightsSlice[n] = "#"
		} else {
			lightsSlice[n] = "."
		}
	}
	return strings.Join(lightsSlice, "")
}

func allOff(lights string) bool {
	for _, l := range lights {
		if l == '.' {
			return false
		}
	}
	return true
}

func doPartTwo(loms []machineSpec) int {
	total := 0
	for _, ms := range loms {
		cache := map[string]int{}
		start := make([]int, len(ms.joltages))
		target := ms.joltages
		fewest := fewestPressesPartTwo(start, target, ms.buttons, cache)
		total += fewest
	}
	return total
}

// Computes the number of presses needed to reach the target state
// from the current state. This is too slow because there are
// too many states (product of the joltages).
func fewestPressesPartTwo(
	state []int,
	target []int,
	buttons [][]int,
	cache map[string]int,
) int {
	if slices.Equal(state, target) {
		return 0
	}

	stateAsString := intSliceToString(state)
	cacheKey := stateAsString

	if e, ok := cache[cacheKey]; ok {
		return e
	}

	fewest := math.MaxInt64
	for _, b := range buttons {
		newState := pressV2(state, target, b)
		if newState == nil {
			continue
		}
		candidate := fewestPressesPartTwo(newState, target, buttons, cache)
		if candidate == math.MaxInt64 {
			continue
		}
		fewest = min(1+candidate, fewest)
	}

	cache[cacheKey] = fewest

	return fewest
}

func pressV2(state []int, target []int, button []int) []int {
	newState := slices.Clone(state)
	for _, b := range button {
		newState[b] += 1
		if newState[b] > target[b] {
			return nil
		}
	}
	return newState
}

func readLines() []string {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseInput(lines []string) []machineSpec {
	loms := []machineSpec{}
	for _, line := range lines {
		toks := strings.Split(line, " ")
		n := len(toks)
		lights := parseLights(toks[0])
		buttons := parseButtons(toks[1 : n-1])
		joltages := parseJoltages(toks[n-1])
		ms := machineSpec{
			lights:   lights,
			buttons:  buttons,
			joltages: joltages,
		}
		loms = append(loms, ms)
	}
	return loms
}

func parseLights(lights string) string {
	n := len(lights)
	return lights[1 : n-1]
}

func parseButtons(los []string) [][]int {
	buttons := make([][]int, len(los), len(los))
	for i, s := range los {
		n := len(s)
		s = s[1 : n-1]
		button := stringSliceToIntSlice(strings.Split(s, ","))
		buttons[i] = button
	}
	return buttons
}

func parseJoltages(s string) []int {
	n := len(s)
	s = s[1 : n-1]
	joltageStrings := strings.Split(s, ",")
	return stringSliceToIntSlice(joltageStrings)
}

func stringSliceToIntSlice(los []string) []int {
	intSlice := make([]int, 0, len(los))
	for _, s := range los {
		intSlice = append(intSlice, stringToInt(s))
	}
	return intSlice
}

func intSliceToString(loi []int) string {
	los := make([]string, 0, len(loi))
	for _, n := range loi {
		los = append(los, strconv.Itoa(n))
	}
	return strings.Join(los, ",")
}

func stringToInt(s string) int {
	n, err := strconv.Atoi(s)
	must(err)
	return n
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
