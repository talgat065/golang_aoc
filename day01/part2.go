package day01

import (
	"bufio"
	"os"
	"strconv"
)

func SonarSweep2(filepath string) int {
	f, err := os.Open(filepath)

	checkErr(err)

	defer closeFile()(f)

	scr := bufio.NewScanner(f)
	
	var records []int
	for scr.Scan() {
		text := scr.Text()
		num, err := strconv.Atoi(text)
		
		checkErr(err)
		
		records = append(records, num)
	}

	previous := 0
	increases := 0
	for i := 2; i < len(records); i++ {
		sum := records[i] + records[i - 1] + records[i - 2]

		if sum > previous {
			increases++
		}

		previous = sum
	}

	return increases - 1
}
