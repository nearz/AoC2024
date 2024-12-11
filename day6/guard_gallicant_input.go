package day6

import (
	"bufio"
	"os"
	"strings"
)

func GuardGallivantFile() ([][]string, error) {
	gmap := [][]string{}
	f, err := os.Open("gg.input")
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
