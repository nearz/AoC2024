package day2

import (
	"bufio"
	"fmt"
	"os"
)

func SafeListFile(safeList []bool, fn string) {
	f, err := os.Create(fn)
	if err != nil {
		fmt.Printf("Failed to create file: %s\n", err)
		return
	}

	defer f.Close()

	writer := bufio.NewWriter(f)

	for _, v := range safeList {
		_, err := writer.WriteString(fmt.Sprintf("%v\n", v))
		if err != nil {
			fmt.Printf("Failed to write to file: %s\n", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Failed to flush to file: %s\n", err)
		return
	}
}
