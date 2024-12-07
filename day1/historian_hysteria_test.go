package day1

import (
	"fmt"
	"testing"
)

func TestHysteriaFile(t *testing.T) {
	file := true
	var left, right []int
	total := Hysteria(file, left, right)
	if total != 1666427 {
		t.Errorf("Hysteria Total = %d, expected 1666427", total)
	}
	fmt.Printf("Total: %d\n", total)
}

func TestHysteria(t *testing.T) {
	file := false
	left := []int{5, 8, 1, 2}
	right := []int{4, 9, 2, 7}
	// 1, 2, 5, 8
	// 2, 4, 7, 9
	// Total = 6
	total := Hysteria(file, left, right)
	if total != 6 {
		t.Errorf("Hysteria Total = %d, expected 6", total)
	}
	fmt.Printf("Total: %d\n", total)
}

func TestHysteriaMultiples(t *testing.T) {
	file := false
	left := []int{3, 3, 3}
	right := []int{3, 3, 3}
	score := HysteriaMultiples(file, left, right)
	if score != 27 {
		t.Errorf("Similarity Score = %d, expected 27", score)
	}
	fmt.Printf("Score: %d\n", score)

	leftTwo := []int{2, 2, 3, 4, 5}
	rightTwo := []int{2, 3, 3, 4, 4}
	// 2 * 1, 2 * 1, 3 * 2, 4 * 2, 5 * 0
	// 2 + 2 + 6 + 8 + 0 = 18
	score = HysteriaMultiples(file, leftTwo, rightTwo)
	if score != 18 {
		t.Errorf("Similarity Score = %d, expected 18", score)
	}
	fmt.Printf("Score: %d\n", score)
}

func TestHysteriaMultiplesFile(t *testing.T) {
	file := true
	var left, right []int
	score := HysteriaMultiples(file, left, right)
	if score != 24316233 {
		t.Errorf("Similarity Score = %d, expected 24316233", score)
	}
	fmt.Printf("Score: %d\n", score)
}
