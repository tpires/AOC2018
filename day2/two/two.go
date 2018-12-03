package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	boxes, err := ReadFrom(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}

	fmt.Println(ComputeBoxes(boxes))
}

func ReadFrom(input io.Reader) ([]string, error) {
	boxes := make([]string, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		str := fmt.Sprintf("%s", scanner.Text())
		boxes = append(boxes, str)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return boxes, nil
}

func Compare(str string, str1 string) (int, string) {
	var count int
	var common string
	for i := 0; i < len(str); i++ {
		if str[i] == str1[i] {
			count++
			common += string(str[i])
		}
	}
	return count, common
}

func ComputeBoxes(boxes []string) string {
	var c string
	min := 0

	for i := 0; i < len(boxes); i++ {
		b := boxes[i]
		for j := i + 1; j < len(boxes); j++ {
			n, common := Compare(b, boxes[j])
			if min < n {
				min = n
				c = common
			}
		}
	}
	return c
}
