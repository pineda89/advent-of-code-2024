package day05

import (
	"advent-of-code-2024/common"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day05 struct {
	filepath string
	rules    map[[2]int]bool
}

func (d *Day05) Part1() string {
	var result int
	d.rules = make(map[[2]int]bool)
	for _, line := range strings.Split(common.GetInput(d.filepath), "\n") {
		if strings.Contains(line, "|") {
			var left, right int
			fmt.Sscanf(line, "%d|%d", &left, &right)
			d.rules[[2]int{left, right}] = true
		} else if strings.Contains(line, ",") {
			var values []int
			for _, element := range strings.Split(line, ",") {
				value, _ := strconv.Atoi(element)
				values = append(values, value)
			}

			if d.isSorted(values) {
				result += values[len(values)/2]
			}
		}
	}

	return strconv.Itoa(result)
}

func (d *Day05) Part2() string {
	var result int
	d.rules = make(map[[2]int]bool)
	for _, line := range strings.Split(common.GetInput(d.filepath), "\n") {
		if strings.Contains(line, "|") {
			var left, right int
			fmt.Sscanf(line, "%d|%d", &left, &right)
			d.rules[[2]int{left, right}] = true
		} else if strings.Contains(line, ",") {
			var values []int
			for _, element := range strings.Split(line, ",") {
				value, _ := strconv.Atoi(element)
				values = append(values, value)
			}

			if !d.isSorted(values) {
				d.sort(values)
				result += values[len(values)/2]
			}
		}
	}

	return strconv.Itoa(result)
}

func (d *Day05) isSorted(values []int) bool {
	return slices.IsSortedFunc(values, d.sortFunc)
}

func (d *Day05) sort(values []int) {
	slices.SortFunc(values, d.sortFunc)
}

func (d *Day05) sortFunc(a int, b int) int {
	if d.rules[[2]int{a, b}] {
		return -1
	}
	return 0
}
