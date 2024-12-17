package day9

import (
	"bufio"
	"os"
	"strings"
)

func DMToFile(dm []string, fn string) {
	jdm := strings.Join(dm, "")
	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(jdm)
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}
}
