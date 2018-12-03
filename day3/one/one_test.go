package main_test

import (
	"strings"
	"testing"

	"github.com/cheekybits/is"
	v "github.com/tpires/AOC2018/day3/one"
)

func TestReadFromOneLine(t *testing.T) {
	is := is.New(t)
	str := "#1 @ 1,3: 4x4"
	claims, err := v.ReadFrom(strings.NewReader(str))

	is.NoErr(err)
	is.Equal(len(claims), 1)

	is.Equal(claims[0].ID, "1")
	is.Equal(claims[0].Start.X, 1)
	is.Equal(claims[0].Start.Y, 3)
	is.Equal(claims[0].Size.X, 4)
	is.Equal(claims[0].Size.Y, 4)
}

func TestReadFromTwoLine(t *testing.T) {
	is := is.New(t)
	str := "#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4"
	claims, err := v.ReadFrom(strings.NewReader(str))

	is.NoErr(err)
	is.Equal(len(claims), 2)

	is.Equal(claims[1].ID, "2")
	is.Equal(claims[1].Start.X, 3)
	is.Equal(claims[1].Start.Y, 1)
	is.Equal(claims[1].Size.X, 4)
	is.Equal(claims[1].Size.Y, 4)
}

func TestConvertFromSingleClaimsToMatrix(t *testing.T) {
	is := is.New(t)
	claims := []v.Claim{{"1", v.Point{1, 3}, v.Point{4, 4}}}
	m := v.ConvertFromClaimsToMatrix(claims)

	is.Equal(len(m), 1000)

	is.Equal(m[1][3], 1)
	is.Equal(m[1+3][3], 1)
	is.Equal(m[1][3+3], 1)
	is.Equal(m[1+3][3+3], 1)
}

func TestConvertFromDoubleClaimsToMatrix(t *testing.T) {
	is := is.New(t)
	claims := []v.Claim{{"1", v.Point{1, 3}, v.Point{4, 4}}, {"2", v.Point{3, 1}, v.Point{4, 4}}}
	m := v.ConvertFromClaimsToMatrix(claims)

	is.Equal(len(m), 1000)

	is.Equal(m[1][3], 1)
	is.Equal(m[1+3][3], 2)
	is.Equal(m[1][3+3], 1)
	is.Equal(m[1+3][3+3], 1)
	is.Equal(m[3][1], 1)
	is.Equal(m[3+3][1], 1)
	is.Equal(m[3][1+3], 2)
	is.Equal(m[3+3][1+3], 1)
}

func TestCountOverlapedInches(t *testing.T) {
	is := is.New(t)
	// init 2d array
	m := make([][]int, 6)
	for i := range m {
		m[i] = make([]int, 6)
	}
	m[2][2] = 2
	m[2][3] = 2
	m[3][2] = 2
	m[3][3] = 2

	is.Equal(v.CountOverlapedInches(m), 4)
}
