package main_test

import (
	"testing"

	"github.com/cheekybits/is"
	v "github.com/tpires/AOC2018/day2/two"
)

func TestCompareEqualStrings(t *testing.T) {
	is := is.New(t)
	c, common := v.Compare("abcde", "abcde")

	is.Equal(c, 5)
	is.Equal(common, "abcde")
}

func TestCompareNonEqualStrings(t *testing.T) {
	is := is.New(t)
	c, common := v.Compare("abcde", "axcye")

	is.Equal(c, 3)
	is.Equal(common, "ace")
}

func TestComputeBoxes(t *testing.T) {
	is := is.New(t)
	ids := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	str := v.ComputeBoxes(ids)

	is.Equal(str, "fgij")
}
