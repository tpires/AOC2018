package main

import (
	"fmt"
	"io"
	"os"
)

type Point struct {
	X int
	Y int
}

type Claim struct {
	ID    string
	Start Point
	Size  Point
}

func main() {
	claims, err := ReadFrom(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}
	fmt.Println(CountOverlapedInches(ConvertFromClaimsToMatrix(claims)))
}

func ReadFrom(input io.Reader) ([]Claim, error) {
	claims := make([]Claim, 0)
	for {
		c := Claim{}
		_, err := fmt.Fscanf(input, "#%s @ %d,%d: %dx%d\n", &c.ID, &c.Start.X, &c.Start.Y, &c.Size.X, &c.Size.Y)
		if err == nil {
			claims = append(claims, c)
		}

		if err != nil {
			if err != io.ErrUnexpectedEOF {
				return nil, err
			}
			return claims, nil
		}
	}
}

func ConvertFromClaimsToMatrix(claims []Claim) [][]int {
	// init 2d array
	m := make([][]int, 1000)
	for i := range m {
		m[i] = make([]int, 1000)
	}
	for _, c := range claims {
		for i := c.Start.X; i < c.Start.X+c.Size.X; i++ {
			for j := c.Start.Y; j < c.Start.Y+c.Size.Y; j++ {
				m[i][j]++
			}
		}
	}
	return m
}

func CountOverlapedInches(m [][]int) int {
	var count int
	for i := 0; i < len(m); i++ {
		l := len(m[i])
		for j := 0; j < l; j++ {
			if m[i][j] > 1 {
				count++
			}
		}
	}
	return count
}
