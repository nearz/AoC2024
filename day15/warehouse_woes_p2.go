package day15

// func PartTwo(wmap [][]string, dirStr string) {
// 	dmap := doubleObjs(wmap)
// 	for _, yv := range dmap {
// 		for _, xv := range yv {
// 			fmt.Print(xv)
// 		}
// 		fmt.Println("")
// 	}
// }
//
// func movesDbl(bot *pos, bxs, wls set, dirStr string) {
// 	for _, d := range dirStr {
// 		var np pos
// 		np.x, np.y = step(string(d), bot.x, bot.y)
// 		_, bok := bxs[np]
// 		_, wok := wls[np]
// 		if !bok && !wok {
// 			*bot = np
// 		} else if bok {
// 			var nbx pos
// 			nbx.x, nbx.y = step(string(d), np.x, np.y)
// 			_, nbok := bxs[nbx]
// 			for nbok {
// 				nbx.x, nbx.y = step(string(d), nbx.x, nbx.y)
// 				_, nbok = bxs[nbx]
// 			}
// 			_, cwl := wls[nbx]
// 			if !cwl {
// 				bxs[nbx] = nbx
// 				nx, ny := step(string(d), bot.x, bot.y)
// 				*bot = pos{x: nx, y: ny}
// 				delete(bxs, *bot)
// 			}
// 		}
// 	}
// }
//
// func doubleObjs(wmap [][]string) [][]string {
// 	dmap := [][]string{}
// 	for _, yv := range wmap {
// 		i := 0
// 		row := make([]string, len(yv)*2)
// 		for _, xv := range yv {
// 			switch xv {
// 			case "@":
// 				fmt.Println(i)
// 				row[i] = "@"
// 				row[i+1] = "."
// 			case "O":
// 				row[i] = "["
// 				row[i+1] = "]"
// 			case "#":
// 				row[i] = "#"
// 				row[i+1] = "#"
// 			default:
// 				row[i] = "."
// 				row[i+1] = "."
// 			}
// 			i += 2
// 			if i >= len(row) {
// 				dmap = append(dmap, row)
// 				break
// 			}
// 		}
// 	}
// 	return dmap
// }
