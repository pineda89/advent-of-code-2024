package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay04Part1TestData(t *testing.T) {
	d := &Day04{
		filepath: "input_testdata.txt",
	}
	result := d.Part1()
	assert.Equal(t, "18", result)
}

func TestDay04Part1RealData(t *testing.T) {
	d := &Day04{
		filepath: "input.txt",
	}
	result := d.Part1()
	assert.Equal(t, "2549", result)
}

func TestDay04Part2TestData(t *testing.T) {
	d := &Day04{
		filepath: "input_testdata.txt",
	}
	result := d.Part2()
	assert.Equal(t, "9", result)
}

func TestDay04Part2RealData(t *testing.T) {
	d := &Day04{
		filepath: "input.txt",
	}
	result := d.Part2()
	assert.Equal(t, "2003", result)
}
