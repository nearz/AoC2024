package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func PartOne(mio string) int {
	mulSlc, err := findMuls(mio)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	total, err := multTotal(mulSlc)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return total
}

func PartTwo(mio string) int {
	rmds, err := removeDonts(mio)
	if err != nil {
		fmt.Println(err)
	}
	mulSlc, err := findMuls(rmds)
	if err != nil {
		fmt.Println(err)
	}
	total, err := multTotal(mulSlc)
	if err != nil {
		fmt.Println(err)
	}
	return total
}

func removeDonts(s string) (string, error) {
	r, err := regexp.Compile(`don't\(\).*?do\(\)`)
	if err != nil {
		return "", err
	}
	rmds := r.ReplaceAllString(s, "-")
	return rmds, nil
}

func findMuls(s string) ([]string, error) {
	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		return nil, err
	}

	mulSlc := r.FindAllString(s, -1)
	return mulSlc, nil
}

func multTotal(ms []string) (int, error) {
	total := 0
	r, err := regexp.Compile(`\d{1,3}`)
	if err != nil {
		return 0, err
	}
	for _, s := range ms {
		intSlc := r.FindAllString(s, -1)
		a, err := strconv.Atoi(intSlc[0])
		if err != nil {
			return 0, err
		}
		b, err := strconv.Atoi(intSlc[1])
		if err != nil {
			return 0, err
		}
		m := a * b
		total += m
	}
	return total, nil
}
