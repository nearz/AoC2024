package day22

import (
	"bufio"
	"os"
	"strconv"
)

func MMFile(fn string) []int {
	snums := []int{}
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		snums = append(snums, n)
	}
	return snums
}
