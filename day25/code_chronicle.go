package day25

type keyLock map[int]int

func PartOne(schematic []string) int {
	locks, keys := convKeysLocks(schematic)
	total := 0
	for _, k := range keys {
		for _, l := range locks {
			match := true
			for i := range 5 {
				if k[i]+l[i] > 5 {
					i = 0
					match = false
					break
				}
			}
			if match {
				total += 1
			}
		}
	}
	return total
}

func convKeysLocks(schematic []string) ([]keyLock, []keyLock) {
	keys := []keyLock{}
	locks := []keyLock{}
	sl := len(schematic)
	i := 0
	for i < sl {
		if schematic[i] == "#####" {
			lock := make(keyLock)
			j := i + 1
			for ; j < i+6; j++ {
				row := schematic[j]
				for r, v := range row {
					if string(v) == "#" {
						if _, lok := lock[r]; !lok {
							lock[r] = 1
						} else {
							lock[r]++
						}
					} else {
						if _, lok := lock[r]; !lok {
							lock[r] = 0
						}
					}
				}
			}
			locks = append(locks, lock)
		} else {
			key := make(keyLock)
			j := i + 1
			for ; j < i+6; j++ {
				row := schematic[j]
				for r, v := range row {
					if string(v) == "#" {
						if _, kok := key[r]; !kok {
							key[r] = 1
						} else {
							key[r]++
						}
					} else {
						if _, kok := key[r]; !kok {
							key[r] = 0
						}
					}
				}
			}
			keys = append(keys, key)
		}
		i += 8
	}
	return locks, keys
}
