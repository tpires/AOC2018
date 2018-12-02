package main_test

import (
	"testing"

	"github.com/cheekybits/is"

	v "github.com/tpires/AOC2018/day1/two"
)

func TestErrorConvertStringToInt(t *testing.T) {
	lines := []string{"+1", "abcd", "+3", "0"}

	is := is.New(t)
	_, err := v.ConvertStringToInt(lines)

	is.Err(err)
}

func TestConvertStringToInt(t *testing.T) {
	lines := []string{"+1", "-1", "+3", "0"}
	expected := []int{1, -1, 3, 0}

	is := is.New(t)
	linesInt, err := v.ConvertStringToInt(lines)

	is.NoErr(err)
	is.Equal(linesInt, expected)
}

func TestExampleFindFirstFrequencyRepeated(t *testing.T) {
	is := is.New(t)

	freqs := []int{1, -2, 3, 1, 1, -2, 3, -4}
	is.Equal(v.FindFirstFrequencyRepeated(freqs), 2)
}

func TestExampleZeroFindFirstFrequencyRepeated(t *testing.T) {
	is := is.New(t)

	freqs := []int{1, -1}
	is.Equal(v.FindFirstFrequencyRepeated(freqs), 0)
}

func TestExampleTenFirstFrequencyRepeated(t *testing.T) {
	is := is.New(t)

	freqs := []int{3, 3, 4, -2, -4}
	is.Equal(v.FindFirstFrequencyRepeated(freqs), 10)
}

func TestExampleFiveFirstFrequencyRepeated(t *testing.T) {
	is := is.New(t)

	freqs := []int{-6, 3, 8, 5, -6}
	is.Equal(v.FindFirstFrequencyRepeated(freqs), 5)
}

func TestExampleFourteenFirstFrequencyRepeated(t *testing.T) {
	is := is.New(t)

	freqs := []int{7, 7, -2, -7, -4}
	is.Equal(v.FindFirstFrequencyRepeated(freqs), 14)
}
