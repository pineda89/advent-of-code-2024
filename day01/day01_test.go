package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay01Part1TestData(t *testing.T) {
	d := &Day01{
		filepath: "input_testdata.txt",
	}
	result := d.Part1()
	assert.Equal(t, "11", result)
}

func TestDay01Part1RealData(t *testing.T) {
	d := &Day01{
		filepath: "input.txt",
	}
	result := d.Part1()
	assert.Equal(t, "1651298", result)
}

func TestDay01Part2TestData(t *testing.T) {
	d := &Day01{
		filepath: "input_testdata.txt",
	}
	result := d.Part2()
	assert.Equal(t, "31", result)
}

func TestDay01Part2RealData(t *testing.T) {
	d := &Day01{
		filepath: "input.txt",
	}
	result := d.Part2()
	assert.Equal(t, "21306195", result)
}
