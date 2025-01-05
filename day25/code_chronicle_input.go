package day25

import (
	"bufio"
	"os"
)

func CCFile(fn string) []string {
	schematic := []string{}
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			schematic = append(schematic, "-")
			continue
		}
		schematic = append(schematic, line)
	}
	return schematic
}
