package day03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay03Part1TestData(t *testing.T) {
	d := &Day03{
		filepath: "input_testdata.txt",
	}
	result := d.Part1()
	assert.Equal(t, "161", result)
}

func TestDay03Part1RealData(t *testing.T) {
	d := &Day03{
		filepath: "input.txt",
	}
	result := d.Part1()
	assert.Equal(t, "166357705", result)
}

func TestDay03Part2TestData(t *testing.T) {
	d := &Day03{
		filepath: "input_testdata_p2.txt",
	}
	result := d.Part2()
	assert.Equal(t, "48", result)
}

func TestDay03Part2RealData(t *testing.T) {
	d := &Day03{
		filepath: "input.txt",
	}
	result := d.Part2()
	assert.Equal(t, "88811886", result)
}
