package day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay05Part1TestData(t *testing.T) {
	d := &Day05{
		filepath: "input_testdata.txt",
	}
	result := d.Part1()
	assert.Equal(t, "143", result)
}

func TestDay05Part1RealData(t *testing.T) {
	d := &Day05{
		filepath: "input.txt",
	}
	result := d.Part1()
	assert.Equal(t, "6041", result)
}

func TestDay05Part2TestData(t *testing.T) {
	d := &Day05{
		filepath: "input_testdata.txt",
	}
	result := d.Part2()
	assert.Equal(t, "123", result)
}

func TestDay05Part2RealData(t *testing.T) {
	d := &Day05{
		filepath: "input.txt",
	}
	result := d.Part2()
	assert.Equal(t, "4884", result)
}
