package day10

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func HoofItFile() ([][]int, error) {
	tmap := [][]int{}
	f, err := os.Open("hi.input")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, "")
		isplt := make([]int, len(splt))
		for i, v := range splt {
			stoi, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			isplt[i] = stoi
		}
		tmap = append(tmap, isplt)
	}
	return tmap, nil
}
