package day3

import (
	"bufio"
	"os"
)

func MullItOverFile() (string, error) {
	fs := ""
	f, err := os.Open("mio.input")
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fs += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return fs, nil
}
