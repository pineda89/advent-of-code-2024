package day02

import (
	"advent-of-code-2024/common"
	"slices"
	"strconv"
	"strings"
)

type Day02 struct {
	filepath string
}

func (d *Day02) Part1() string {
	var count int
	for _, line := range strings.Split(common.GetInput(d.filepath), "\n") {
		var report []int
		for _, measure := range strings.Fields(line) {
			m, _ := strconv.Atoi(measure)
			report = append(report, m)
		}

		if d.validate(report) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func (d *Day02) Part2() string {
	var count int
	for _, line := range strings.Split(common.GetInput(d.filepath), "\n") {
		var report []int
		for _, measure := range strings.Fields(line) {
			m, _ := strconv.Atoi(measure)
			report = append(report, m)
		}

		if d.validate(report) || d.validateAfterSingleRemoval(report) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func (d *Day02) validate(report []int) bool {
	var valids []int
	if report[1] > report[0] {
		valids = []int{1, 2, 3}
	} else if report[1] < report[0] {
		valids = []int{-1, -2, -3}
	}

	for i := 1; i < len(report); i++ {
		if !slices.Contains(valids, report[i]-report[i-1]) {
			return false
		}
	}

	return true
}

func (d *Day02) validateAfterSingleRemoval(report []int) bool {
	for i := range report {
		copiedReport := make([]int, len(report))
		copy(copiedReport, report)

		if d.validate(slices.Delete(copiedReport, i, i+1)) {
			return true
		}
	}
	return false
}
