package day1

import (
	"sort"
)

func PartOne(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	total := distanceTotal(left, right)
	return total
}

func PartTwo(left, right []int) int {
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
