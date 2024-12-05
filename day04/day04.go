package day04

import (
	"advent-of-code-2024/common"
	"strconv"
	"strings"
)

type Day04 struct {
	filepath string
}

var (
	directions = [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	directions_p2 = [][2]int{
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}
)

func (d *Day04) Part1() string {
	var result int
	var matrix [][]rune
	for i, line := range strings.Split(common.GetInput(d.filepath), "\n") {
		matrix = append(matrix, make([]rune, len(line)))
		for j, c := range line {
			matrix[i][j] = c
		}
	}

	const word = "XMAS"
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == rune(word[0]) {
				for _, direction := range directions {
					result += d.findWordInDirection(matrix, i, j, word, direction)
				}
			}
		}
	}

	return strconv.Itoa(result)
}

func (d *Day04) Part2() string {
	var result int
	var matrix [][]rune
	for i, line := range strings.Split(common.GetInput(d.filepath), "\n") {
		matrix = append(matrix, make([]rune, len(line)))
		for j, c := range line {
			matrix[i][j] = c
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 'A' {
				if i > 0 && j > 0 && i < len(matrix)-1 && j < len(matrix[i])-1 {
					values := make(map[rune]int)
					for _, direction := range directions_p2 {
						values[matrix[i+direction[0]][j+direction[1]]]++
					}

					if values['M'] == 2 && values['S'] == 2 && matrix[i-1][j-1] != matrix[i+1][j+1] {
						result++
					}
				}
			}
		}
	}

	return strconv.Itoa(result)
}

func (d *Day04) findWordInDirection(matrix [][]rune, i int, j int, word string, direction [2]int) int {
	if i < 0 || j < 0 || j >= len(matrix[0]) || i >= len(matrix) {
		return 0
	}

	if len(word) == 1 {
		if matrix[i][j] == rune(word[0]) {
			return 1
		}
		return 0
	}

	var result int

	if matrix[i][j] == rune(word[0]) {
		result += d.findWordInDirection(matrix, i+direction[0], j+direction[1], word[1:], direction)
	}

	return result
}
