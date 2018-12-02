package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	freqs, err := ReadFileFromStdin()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}
	fmt.Println(ProcessFrequencies(freqs))
}

func ReadFileFromStdin() ([]string, error) {
	boxes := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
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
