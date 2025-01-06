package day24

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CWFile(fn string) (map[string]int, []gate) {
	init := make(map[string]int)
	gates := []gate{}
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sec := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sec = true
			continue
		}
		if !sec {
			splt := strings.Split(line, ":")
			val, err := strconv.Atoi(strings.TrimSpace(splt[1]))
			if err != nil {
				panic(err)
			}
			init[splt[0]] = val
		} else {
			var g gate
			splt := strings.Split(line, "->")
			g.out = wire{name: strings.TrimSpace(splt[1]), val: 0}
			gsplt := strings.Split(strings.TrimSpace(splt[0]), " ")
			g.logic = gsplt[1]
			g.in1 = wire{name: strings.TrimSpace(gsplt[0]), val: 0}
			g.in2 = wire{name: strings.TrimSpace(gsplt[2]), val: 0}
			gates = append(gates, g)
		}
	}
	return init, gates
}
