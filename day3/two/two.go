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
	ID      string
	Start   Point
	Size    Point
	Overlap bool
}

func main() {
	claims, err := ReadFrom(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}
	FindNonOverlapped(claims)
	fmt.Println(GetNonOverlapped(claims))
}

func ReadFrom(input io.Reader) ([]*Claim, error) {
	claims := make([]*Claim, 0)
	for {
		c := Claim{}
		_, err := fmt.Fscanf(input, "#%s @ %d,%d: %dx%d\n", &c.ID, &c.Start.X, &c.Start.Y, &c.Size.X, &c.Size.Y)
		if err == nil {
			claims = append(claims, &c)
		}

		if err != nil {
			if err != io.ErrUnexpectedEOF {
				return nil, err
			}
			return claims, nil
		}
	}
}

func FindNonOverlapped(claims []*Claim) {
	// init 2d array
	m := make([][][]*Claim, 1000)
	for i := range m {
		m[i] = make([][]*Claim, 1000)
		for j := range m[i] {
			m[i][j] = make([]*Claim, 0)
		}
	}
	for _, c := range claims {
		for i := c.Start.X; i < c.Start.X+c.Size.X; i++ {
			for j := c.Start.Y; j < c.Start.Y+c.Size.Y; j++ {
				m[i][j] = append(m[i][j], c)
				if len(m[i][j]) > 1 {
					markAsOverlapped(m[i][j])
				}
			}
		}
	}
}

func GetNonOverlapped(claim []*Claim) string {
	for _, c := range claim {
		if !c.Overlap {
			return c.ID
		}
	}
	return ""
}

func markAsOverlapped(claims []*Claim) {
	for i := 0; i < len(claims); i++ {
		claims[i].Overlap = true
	}
}
