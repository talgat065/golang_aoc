package day01

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func SonarSweep(filepath string) int {
	f, err := os.Open(filepath)

	checkErr(err)

	defer closeFile()(f)

	scr := bufio.NewScanner(f)

	increases := 0
	previous := 0
	for scr.Scan() {
		text := scr.Text()
		current, err := strconv.Atoi(text)

		checkErr(err)

		if current > previous {
			increases++
		}

		previous = current
	}

	return increases - 1
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func closeFile() func(f *os.File) {
	return func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
