package day8

import (
	"bufio"
	"os"
	"strings"
)

func RCFile() ([][]string, error) {
	gmap := [][]string{}
	f, err := os.Open("rc.input")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splt := strings.Split(line, "")
		gmap = append(gmap, splt)
	}
	return gmap, nil
}
