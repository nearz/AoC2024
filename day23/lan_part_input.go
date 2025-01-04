package day23

import (
	"bufio"
	"os"
	"strings"
)

func LPFile(fn string) []cpair {
	cprs := []cpair{}
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var cpr cpair
		line := scanner.Text()
		splt := strings.Split(line, "-")
		cpr.c1 = splt[0]
		cpr.c2 = splt[1]
		cprs = append(cprs, cpr)
	}
	return cprs
}
