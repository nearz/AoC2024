package day6

import (
	"bufio"
	"os"
	"strings"
)

func GMapToFile(gmap [][]string) {
	fn := "gmap.output"
	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range gmap {
		jl := strings.Join(line, "")
		_, err := w.WriteString(jl + "\n")
		if err != nil {
			panic(err)
		}
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
}
