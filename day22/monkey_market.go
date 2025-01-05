package day22

func PartOne(snums []int) int {
	total := 0
	for _, sn := range snums {
		res := sn
		for range 2000 {
			res = secretNumRound(res)
		}
		total += res
	}
	return total
}

// 1. sn * 64 -> mix sn - prune = res
// 2. res / 32floor -> mix - prune = res
// 3. res * 2048 -> mix - prune = res
// mix = res xor sn = res
// res % 16777216 = final sn
func secretNumRound(sn int) int {
	res := ((sn * 64) ^ sn) % 16777216
	res = ((res / 32) ^ res) % 16777216
	res = ((res * 2048) ^ res) % 16777216
	return res
}
