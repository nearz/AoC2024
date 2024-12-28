package day15

import (
	"bufio"
	"os"
	"strings"
)

func WWFile(fn string) ([][]string, string) {
	wmap := [][]string{}
	var dirStr string
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
			splt := strings.Split(line, "")
			wmap = append(wmap, splt)
		} else {
			dirStr += line
		}
	}
	return wmap, dirStr
}
