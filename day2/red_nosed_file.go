package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RedNosedFile() ([][]int, error) {
	reports := [][]int{}
	f, err := os.Open("reports.input")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, " ")
		ints := make([]int, len(splt))
		for i := range len(splt) {
			n, err := strconv.Atoi(splt[i])
			if err != nil {
				return nil, err
			}
			ints[i] = n
		}
		reports = append(reports, ints)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return reports, nil
}
