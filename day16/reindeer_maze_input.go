package day16

import (
	"bufio"
	"os"
	"strings"
)

func RMFile(fn string) [][]string {
	rmap := [][]string{}
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, "")
		rmap = append(rmap, splt)
	}
	return rmap
}
