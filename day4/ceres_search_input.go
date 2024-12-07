package day4

import (
	"bufio"
	"os"
	"strings"
)

func CeresSearchFile() ([][]string, error) {
	ceres := [][]string{}
	f, err := os.Open("ceres.input")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, "")
		ceres = append(ceres, splt)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ceres, nil
}
