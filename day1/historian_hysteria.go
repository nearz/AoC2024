package day1

import (
	"fmt"
	"sort"
)

func Hysteria(file bool, paramLeft, paramRight []int) int {
	var left, right []int
	var err error
	if file {
		left, right, err = HysteriaFile()
	} else {
		left, right = paramLeft, paramRight
		err = nil
	}
	if err != nil {
		fmt.Println(err)
		return -1
	}
	sort.Ints(left)
	sort.Ints(right)
	total := distanceTotal(left, right)
	return total
}

func HysteriaMultiples(file bool, paramLeft, paramRight []int) int {
	var left, right []int
	var err error
	if file {
		left, right, err = HysteriaFile()
	} else {
		left, right = paramLeft, paramRight
		err = nil
	}
	if err != nil {
		fmt.Println(err)
	}
	scoreTotal := similarityScore(left, right)
	return scoreTotal
}

func similarityScore(left, right []int) int {
	score := 0
	count := 0
	for i := range len(left) {
		for k := range len(right) {
			if left[i] == right[k] {
				count++
			}
		}
		score += left[i] * count
		count = 0
	}
	return score
}

func distanceTotal(left, right []int) int {
	total := 0
	for i := range len(left) {
		total += abs(left[i] - right[i])
	}
	return total
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
