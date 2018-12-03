package main_test

import (
	"strings"
	"testing"

	"github.com/cheekybits/is"
	v "github.com/tpires/AOC2018/day2/one"
)

func TestReadFromSingleExample(t *testing.T) {
	is := is.New(t)
	str := "abcdef"
	expected := []string{"abcdef"}
	boxes, err := v.ReadFrom(strings.NewReader(str))

	is.NoErr(err)
	is.Equal(len(boxes), 1)

	is.Equal(boxes, expected)
}

func TestReadFromExample(t *testing.T) {
	is := is.New(t)
	str := "abcdef\nbababc\nabbcde"
	expected := []string{"abcdef", "bababc", "abbcde"}
	boxes, err := v.ReadFrom(strings.NewReader(str))

	is.NoErr(err)
	is.Equal(len(boxes), 3)

	is.Equal(boxes, expected)
}

func TestCountNoRepeatedLetters(t *testing.T) {
	is := is.New(t)

	two, three := v.Count("abcde")
	is.Equal(two, 0)
	is.Equal(three, 0)
}

func TestCountTwoAndThreeRepeatedLetters(t *testing.T) {
	is := is.New(t)

	two, three := v.Count("bababc")
	is.Equal(two, 1)
	is.Equal(three, 1)
}

func TestCountTwoAndNoThreeRepeatedLetters(t *testing.T) {
	is := is.New(t)

	two, three := v.Count("abbcde")
	is.Equal(two, 1)
	is.Equal(three, 0)
}

func TestCountNoTwoAndThreeRepeatedLetters(t *testing.T) {
	is := is.New(t)

	two, three := v.Count("abcccd")
	is.Equal(two, 0)
	is.Equal(three, 1)
}

func TestCountTwoAndNoThreeRepeatedLettersButOnlyCountOnce(t *testing.T) {
	is := is.New(t)

	two, three := v.Count("aabcdd")
	is.Equal(two, 1)
	is.Equal(three, 0)
}

func TestCountNoTwoAndThreeRepeatedLettersButOnlyCountOnce(t *testing.T) {
	is := is.New(t)

	two, three := v.Count("ababab")
	is.Equal(two, 0)
	is.Equal(three, 1)
}

func TestProcessChecksum(t *testing.T) {
	is := is.New(t)
	boxes := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}

	is.Equal(v.ProcessChecksum(boxes), 12)
}
