package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	boxes, err := ReadFileFromStdin()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		panic(err)
	}
	fmt.Println(ProcessChecksum(boxes))
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

func Count(str string) (int, int) {
	var two, three int
	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		c := str[i]
		m[c]++
	}
	for _, v := range m {
		if v == 2 {
			two = 1
		} else if v == 3 {
			three = 1
		}
	}
	return two, three
}

func ProcessChecksum(boxes []string) int {
	var totalTwo, totalThree int
	for _, v := range boxes {
		two, three := Count(v)
		totalTwo += two
		totalThree += three
	}
	return totalTwo * totalThree
}
