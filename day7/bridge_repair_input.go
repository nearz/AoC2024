package day7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func BridgeRepairFile() ([]Equation, error) {
	eqtns := []Equation{}
	f, err := os.Open("br.input")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		e := Equation{}
		line := scanner.Text()
		splt := strings.Split(line, ":")
		sopds := strings.Split(splt[1], " ")
		sopds = sopds[1:]
		opds := make([]int, len(sopds))
		for i, s := range sopds {
			opds[i], err = strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
		}
		e.res, err = strconv.Atoi(splt[0])
		if err != nil {
			return nil, err
		}
		e.operands = opds
		eqtns = append(eqtns, e)

	}
	return eqtns, nil
}
