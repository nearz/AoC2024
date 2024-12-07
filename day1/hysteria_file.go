package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HysteriaFile() ([]int, []int, error) {
	var leftList []int
	var rightList []int

	f, err := os.Open("lists.input")
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, "   ")
		left, err := strconv.Atoi(splt[0])
		if err != nil {
			return nil, nil, err
		}

		right, err := strconv.Atoi(splt[1])
		if err != nil {
			return nil, nil, err
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while reading file: %s\n", err)
	}

	return leftList, rightList, nil
}
