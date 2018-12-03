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
	freqsInt, err := ConvertStringToInt(freqs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "processing string to int:", err)
		panic(err)
	}
	fmt.Println(FindFirstFrequencyRepeated(freqsInt))
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

func ConvertStringToInt(freqs []string) ([]int, error) {
	freqsInt := make([]int, 0)
	for _, v := range freqs {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		freqsInt = append(freqsInt, i)
	}
	return freqsInt, nil
}

func FindFirstFrequencyRepeated(freqsInt []int) int {
	var frequency int
	frequencyMap := make(map[int]int, 1)
	frequencyMap[0] = 1

	for {
		for _, change := range freqsInt {
			frequency += change
			if _, ok := frequencyMap[frequency]; ok {
				return frequency
			}
			frequencyMap[frequency] = 1
		}
	}
}
