package day9

import (
	"bufio"
	"os"
)

func DFFile() (string, error) {
	dm := ""
	f, err := os.Open("df.input")
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		dm += line
	}
	return dm, nil
}
