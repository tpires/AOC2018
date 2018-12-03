package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	freqs, err := ReadFrom(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}
	fmt.Println(ProcessFrequencies(freqs))
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

func ProcessFrequencies(freqs []string) int {
	var frequency int
	for _, s := range freqs {
		change, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		frequency += change
	}
	return frequency
}
