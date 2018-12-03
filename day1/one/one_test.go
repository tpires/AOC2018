package main_test

import (
	"strings"
	"testing"

	"github.com/cheekybits/is"
	v "github.com/tpires/AOC2018/day1/one"
)

func TestReadFromExample(t *testing.T) {
	is := is.New(t)
	str := "+1\n-2\n+3\n+1"
	expected := []string{"+1", "-2", "+3", "+1"}
	freqs, err := v.ReadFrom(strings.NewReader(str))

	is.NoErr(err)
	is.Equal(len(freqs), 4)

	is.Equal(freqs, expected)
}

func TestExampleProcessFrequencies(t *testing.T) {
	is := is.New(t)
	freqs := []string{"+1", "-2", "+3", "+1"}
	is.Equal(v.ProcessFrequencies(freqs), 3)
}

func TestSumExampleProcessFrequencies(t *testing.T) {
	is := is.New(t)
	freqs := []string{"+1", "+1", "+1"}
	is.Equal(v.ProcessFrequencies(freqs), 3)
}

func TestSumAndSubtractExampleProcessFrequencies(t *testing.T) {
	is := is.New(t)
	freqs := []string{"+1", "+1", "-2"}
	is.Equal(v.ProcessFrequencies(freqs), 0)
}

func TestSubtractExampleProcessFrequencies(t *testing.T) {
	is := is.New(t)
	freqs := []string{"-1", "-2", "-3"}
	is.Equal(v.ProcessFrequencies(freqs), -6)
}
