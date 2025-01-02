package day18

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RRFile(fn string) []pos {
	pslc := []pos{}
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, ",")
		x, err := strconv.Atoi(splt[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(splt[1])
		if err != nil {
			panic(err)
		}
		p := pos{x: x, y: y}
		pslc = append(pslc, p)
	}
	return pslc
}
