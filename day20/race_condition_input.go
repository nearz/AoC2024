package day20

import (
	"bufio"
	"os"
	"strings"
)

func RCFile(fn string) [][]string {
	rcmap := [][]string{}
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, "")
		rcmap = append(rcmap, splt)
	}
	return rcmap
}
