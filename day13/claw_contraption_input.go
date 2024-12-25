package day13

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func CCFile(fn string) []game {
	games := []game{}
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 1
	var g game
	var ab button
	var bb button
	for scanner.Scan() {
		line := scanner.Text()
		cline := i % 4
		switch cline {
		case 1:
			ab, err = parseBttn(line, 3)
			if err != nil {
				panic(err)
			}
		case 2:
			bb, err = parseBttn(line, 1)
			if err != nil {
				panic(err)
			}
		case 3:
			g, err = parseGame(line)
			if err != nil {
				panic(err)
			}
			g.a = ab
			g.b = bb
			games = append(games, g)
		}
		i++
	}
	return games
}

func parseBttn(s string, c int) (button, error) {
	var a button
	splt := strings.Split(s, " ")
	xs := strings.Split(splt[2], "+")[1]
	x, err := strconv.Atoi(xs[:len(xs)-1])
	if err != nil {
		return a, err
	}
	y, err := strconv.Atoi(strings.Split(splt[3], "+")[1])
	if err != nil {
		return a, err
	}
	a.x, a.y = x, y
	a.cost = c
	return a, nil
}

func parseGame(s string) (game, error) {
	var g game
	splt := strings.Split(s, " ")
	xs := strings.Split(splt[1], "=")[1]
	x, err := strconv.Atoi(xs[:len(xs)-1])
	if err != nil {
		return g, err
	}
	y, err := strconv.Atoi(strings.Split(splt[2], "=")[1])
	if err != nil {
		return g, err
	}
	g.x, g.y = x, y
	return g, nil
}
