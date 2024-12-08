package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func printQueueFile() ([]string, error) {
	s := []string{}
	f, err := os.Open("pq.input")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return s, nil
}

func Rules() [][]int {
	s, err := printQueueFile()
	if err != nil {
		panic(err)
	}
	intSlc := [][]int{}
	for _, v := range s {
		if v == "" {
			break
		}
		container := make([]int, 2)
		splt := strings.Split(v, "|")
		x, err := strconv.Atoi(splt[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(splt[1])
		if err != nil {
			panic(err)
		}
		container[0] = x
		container[1] = y
		intSlc = append(intSlc, container)

	}
	return intSlc
}

func ActivePrintQueues() [][]int {
	s, err := printQueueFile()
	if err != nil {
		panic(err)
	}
	intSlc := [][]int{}
	i := len(s) - 1
	for ; s[i] != ""; i-- {
		splt := strings.Split(s[i], ",")
		container := make([]int, len(splt))
		for i, sv := range splt {
			sti, err := strconv.Atoi(sv)
			if err != nil {
				panic(err)
			}
			container[i] = sti
		}
		intSlc = append(intSlc, container)
	}
	return intSlc
}
