package day19

import (
	"bufio"
	"os"
	"strings"
)

func LLFile(fn string) ([]string, []string) {
	twls := []string{}
	dsgns := []string{}
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
			splt := strings.Split(line, ", ")
			for _, s := range splt {
				twls = append(twls, s)
			}
		} else {
			dsgns = append(dsgns, line)
		}
	}
	return twls, dsgns
}
