package day2

func PartOne(reports [][]int) map[string]int {
	reportsSafety := noDampener(reports)
	countMap := safetyCount(reportsSafety)
	return countMap
}

func PartTwo(reports [][]int) map[string]int {
	reportsSafety := dampener(reports)
	countMap := safetyCount(reportsSafety)
	return countMap
}

func dampener(reports [][]int) []bool {
	safe := make([]bool, len(reports))
	for i, r := range reports {
		levelIdx := unsafeLevel(r)
		if levelIdx == -1 {
			safe[i] = true
		} else {
			dampened := unsafeLevel(removeLevel(r, levelIdx)) == -1
			levelIdx++
			dampened2 := unsafeLevel(removeLevel(r, levelIdx)) == -1
			levelIdx = 0
			dampened3 := unsafeLevel(removeLevel(r, levelIdx)) == -1
			if dampened || dampened2 || dampened3 {
				safe[i] = true
			} else {
				safe[i] = false
			}
		}
	}
	return safe
}

func removeLevel(report []int, idx int) []int {
	cp := make([]int, len(report))
	copy(cp, report)
	cp = append(cp[:idx], cp[idx+1:]...)
	return cp
}

func noDampener(reports [][]int) []bool {
	safe := make([]bool, len(reports))
	for i, v := range reports {
		if unsafeLevel(v) == -1 {
			safe[i] = true
		} else {
			safe[i] = false
		}
	}
	return safe
}

func unsafeLevel(report []int) int {
	ascend := report[0] < report[1]
	ln := len(report)
	for i := 0; i < ln-1; i++ {
		seq := (ascend && report[i] < report[i+1]) || (!ascend && report[i] > report[i+1])
		dif := abs(report[i+1]-report[i]) < 4
		if !seq || !dif {
			return i
		}
	}
	return -1
}

func safetyCount(reportsSafety []bool) map[string]int {
	safety := map[string]int{
		"safe":   0,
		"unsafe": 0,
	}
	for i := range len(reportsSafety) {
		if reportsSafety[i] {
			safety["safe"] += 1
		} else {
			safety["unsafe"] += 1
		}
	}
	return safety
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
