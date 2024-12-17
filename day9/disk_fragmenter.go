package day9

import (
	"strconv"
)

func PartOne(dm string) int {
	exp, err := expand(dm)
	if err != nil {
		panic(err)
	}
	moveFrags(exp)
	return checksum(exp)
}

func PartTwo(dm string) int {
	exp, err := expand(dm)
	if err != nil {
		panic(err)
	}
	moveFiles(exp)
	return checksum(exp)
}

func checksum(exp []int) int {
	total := 0
	l := len(exp)
	for i := 0; i < l; i++ {
		if exp[i] != -1 {
			total += i * exp[i]
		} else {
			continue
		}
	}
	return total
}

func moveFiles(exp []int) {
	hi := len(exp) - 1
	hi2 := hi
	for true {
		for hi >= 0 && exp[hi] == -1 {
			hi--
			hi2--
		}
		for hi2 >= 0 && exp[hi2] == exp[hi] {
			hi2--
		}
		if hi2 < 0 || hi < 0 {
			return
		}
		lo, lo2 := findFree(len(exp[hi2:hi]), hi2+1, exp)
		if lo != -1 {
			copy(exp[lo:lo2], exp[hi2+1:hi+1])
			for i := hi2 + 1; i < hi+1; i++ {
				exp[i] = -1
			}
		}
		hi = hi2
	}
}

func findFree(size, end int, exp []int) (int, int) {
	lo := 0
	lo2 := 0
	for true {
		for lo < end && exp[lo] != -1 {
			lo++
			lo2++
		}
		for lo2 < end && exp[lo2] == -1 {
			lo2++
		}
		if lo2-lo >= size {
			return lo, lo2
		}
		if lo2 > end {
			break
		}
		lo = lo2 + 1
		lo2++
	}
	return -1, -1
}

func moveFrags(exp []int) {
	hi := len(exp) - 1
	lo := 0
	for true {
		for lo < hi && exp[lo] != -1 {
			lo++
		}
		for hi > lo && exp[hi] == -1 {
			hi--
		}
		if lo >= hi {
			return
		}
		exp[lo], exp[hi] = exp[hi], exp[lo]
	}
}

func expand(dm string) ([]int, error) {
	exp := []int{}
	isFile := true
	id := 0
	for _, d := range dm {
		if isFile {
			d, err := strconv.Atoi(string(d))
			if err != nil {
				return nil, err
			}
			for range d {
				exp = append(exp, id)
			}
			id++
		} else {
			d, err := strconv.Atoi(string(d))
			if err != nil {
				return nil, err
			}
			for range d {
				exp = append(exp, -1)
			}
		}
		isFile = !isFile
	}
	return exp, nil
}

// func expand(dm string) ([]string, error) {
// 	var exp string
// 	idx := 0
// 	inc := false
// 	for i, v := range dm {
// 		iv, err := strconv.Atoi(string(v))
// 		if err != nil {
// 			return nil, err
// 		}
// 		for range iv {
// 			if i%2 != 1 {
// 				exp += strconv.Itoa(idx)
// 				inc = true
// 			} else {
// 				exp += "."
// 			}
// 		}
// 		if inc {
// 			idx++
// 			inc = false
// 		}
// 	}
// 	splt := strings.Split(exp, "")
// 	return splt, nil
// }
