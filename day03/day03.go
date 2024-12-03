package day03

import (
	"advent-of-code-2024/common"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// 165784915

var r = regexp.MustCompile(`mul\(\d+,\d+\)`)
var r_do = regexp.MustCompile(`do\(\)`)
var r_dont = regexp.MustCompile(`don't\(\)`)

type Day03 struct {
	filepath string
}

func (d *Day03) Part1() string {
	var result int
	for _, line := range strings.Split(common.GetInput(d.filepath), "\n") {
		all := r.FindAllStringSubmatch(line, -1)
		for j := range all {
			var v1, v2 int
			fmt.Sscanf(all[j][0], "mul(%d,%d)", &v1, &v2)
			if v1 > 0 && v1 <= 999 && v2 > 0 && v2 <= 999 {
				result += v1 * v2
			}
		}
	}

	return strconv.Itoa(result)
}

type pair struct {
	index     int
	values    [2]int
	operation int8
	line      int
}

func (p pair) String() string {
	return fmt.Sprintf("index %d values (%d %d) operation %d line %d\n", p.index, p.values[0], p.values[1], p.operation, p.line)
}

func (d *Day03) Part2() string {
	var pairs []pair
	for l, line := range strings.Split(common.GetInput(d.filepath), "\n") {
		do_indexes := r_do.FindAllStringIndex(line, -1)
		dont_indexes := r_dont.FindAllStringIndex(line, -1)
		r_indexes := r.FindAllStringIndex(line, -1)

		for i := range do_indexes {
			pairs = append(pairs, pair{
				index:     do_indexes[i][0],
				operation: 1,
				line:      l,
			})
		}

		for i := range dont_indexes {
			pairs = append(pairs, pair{
				index:     dont_indexes[i][0],
				operation: 2,
				line:      l,
			})
		}

		for i := range r_indexes {
			p := pair{
				index:     r_indexes[i][0],
				operation: 3,
				line:      l,
			}
			fmt.Sscanf(line[r_indexes[i][0]:r_indexes[i][1]], "mul(%d,%d)", &p.values[0], &p.values[1])
			pairs = append(pairs, p)
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].line == pairs[j].line {
			return pairs[i].index < pairs[j].index
		}
		return pairs[i].line < pairs[j].line
	})

	var mustDo = true
	var result int
	for i := range pairs {
		switch pairs[i].operation {
		case 1:
			mustDo = true
		case 2:
			mustDo = false
		case 3:
			if mustDo {
				result += pairs[i].values[0] * pairs[i].values[1]
			}
		}
	}

	return strconv.Itoa(result)
}
