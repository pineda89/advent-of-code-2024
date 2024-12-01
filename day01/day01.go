package day01

import (
	"advent-of-code-2024/common"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day01 struct {
	filepath string
}

func (d *Day01) Part1() string {
	var left, right []int
	for _, line := range strings.Split(common.GetInput(d.filepath), "\n") {
		var l, r int
		fmt.Sscanf(line, "%d   %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	var distance int
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	return strconv.Itoa(distance)
}

func (d *Day01) Part2() string {
	var left []int
	var right = make(map[int]int)
	for _, line := range strings.Split(common.GetInput(d.filepath), "\n") {
		var l, r int
		fmt.Sscanf(line, "%d   %d", &l, &r)
		left = append(left, l)
		right[r]++
	}

	var similarity int
	for i := range left {
		similarity += left[i] * right[left[i]]
	}

	return strconv.Itoa(similarity)
}
