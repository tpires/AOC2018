package main_test

import (
	"strings"
	"testing"

	"github.com/cheekybits/is"
	v "github.com/tpires/AOC2018/day3/two"
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

func TestFindOverlapped(t *testing.T) {
	is := is.New(t)
	claims := []*v.Claim{{"1", v.Point{1, 3}, v.Point{4, 4}, false}, {"2", v.Point{3, 1}, v.Point{4, 4}, false}}

	v.FindNonOverlapped(claims)
	is.Equal(len(claims), 2)
	is.True(claims[0].Overlap)
	is.True(claims[1].Overlap)
}

func TestFindNonOverlapped(t *testing.T) {
	is := is.New(t)
	claims := []*v.Claim{{"1", v.Point{1, 3}, v.Point{4, 4}, false}, {"2", v.Point{3, 1}, v.Point{4, 4}, false}, {"3", v.Point{5, 5}, v.Point{2, 2}, false}}

	v.FindNonOverlapped(claims)
	is.Equal(len(claims), 3)
	is.True(claims[0].Overlap)
	is.True(claims[1].Overlap)
	is.False(claims[2].Overlap)
}

func TestGetNonOverlapped(t *testing.T) {
	is := is.New(t)
	claims := []*v.Claim{{"1", v.Point{1, 3}, v.Point{4, 4}, false}, {"2", v.Point{3, 1}, v.Point{4, 4}, false}}

	v.FindNonOverlapped(claims)
	is.Equal(len(claims), 2)
	is.True(claims[0].Overlap)
	is.True(claims[1].Overlap)
	is.Equal(v.GetNonOverlapped(claims), "")
}

func TestGetOverlapped(t *testing.T) {
	is := is.New(t)
	claims := []*v.Claim{{"1", v.Point{1, 3}, v.Point{4, 4}, false}, {"2", v.Point{3, 1}, v.Point{4, 4}, false}, {"3", v.Point{5, 5}, v.Point{2, 2}, false}}

	v.FindNonOverlapped(claims)
	is.Equal(len(claims), 3)
	is.True(claims[0].Overlap)
	is.True(claims[1].Overlap)
	is.False(claims[2].Overlap)
	is.Equal(v.GetNonOverlapped(claims), "3")
}
