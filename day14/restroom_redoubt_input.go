package day14

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RRInput(fn string) []bot {
	bots := []bot{}
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, " ")
		sp := strings.Split(splt[0], "=")[1]
		sv := strings.Split(splt[1], "=")[1]

		var p pos
		var v pos
		p.x, p.y = mAtoi(strings.Split(sp, ",")[0]), mAtoi(strings.Split(sp, ",")[1])
		v.x, v.y = mAtoi(strings.Split(sv, ",")[0]), mAtoi(strings.Split(sv, ",")[1])

		b := bot{
			p: p,
			v: v,
		}
		bots = append(bots, b)

	}
	return bots
}

func mAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
