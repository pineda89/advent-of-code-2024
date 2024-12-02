package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay02Part1TestData(t *testing.T) {
	d := &Day02{
		filepath: "input_testdata.txt",
	}
	result := d.Part1()
	assert.Equal(t, "2", result)
}

func TestDay02Part1RealData(t *testing.T) {
	d := &Day02{
		filepath: "input.txt",
	}
	result := d.Part1()
	assert.Equal(t, "516", result)
}

func TestDay02Part2TestData(t *testing.T) {
	d := &Day02{
		filepath: "input_testdata.txt",
	}
	result := d.Part2()
	assert.Equal(t, "4", result)
}

func TestDay02Part2RealData(t *testing.T) {
	d := &Day02{
		filepath: "input.txt",
	}
	result := d.Part2()
	assert.Equal(t, "561", result)
}
