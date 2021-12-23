package day03

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type bitsCount struct {
	zero int
	one int
}

func BinaryDiagnostic(filepath string) int64 {
	f, err := os.Open(filepath)

	checkErr(err)

	defer closeFile()(f)

	bits := map[int]*bitsCount{
		0: {zero: 0, one: 0},
		1: {zero: 0, one: 0},
		2: {zero: 0, one: 0},
		3: {zero: 0, one: 0},
		4: {zero: 0, one: 0},
		5: {zero: 0, one: 0},
		6: {zero: 0, one: 0},
		7: {zero: 0, one: 0},
		8: {zero: 0, one: 0},
		9: {zero: 0, one: 0},
		10: {zero: 0, one: 0},
		11: {zero: 0, one: 0},
	}

	scr := bufio.NewScanner(f)

	for scr.Scan() {
		text := scr.Text()

		bitsStr := strings.Split(text, "")

		for ind, byteStr := range bitsStr {
			num := stringToInt(byteStr)

			if num == 0 {
				if entry, ok := bits[ind]; ok {
					entry.zero += 1
				}
			}

			if num == 1 {
				if entry, ok := bits[ind]; ok {
					entry.one += 1
				}
			}
		}
	}

	mostUsed := make([]string, 12)
	lessUsed := make([]string, 12)

	for i := 0; i < len(bits); i++ {
		if entry, ok := bits[i]; ok {
			if entry.zero > entry.one {
				mostUsed = append(mostUsed, "1")

				lessUsed = append(lessUsed, "0")
			} else {
				mostUsed = append(mostUsed, "0")

				lessUsed = append(lessUsed, "1")
			}
		}
	}

	gammaRate := parseIntFromBinaryStringNum(strings.Join(mostUsed, ""))
	epsilonRate := parseIntFromBinaryStringNum(strings.Join(lessUsed, ""))

	return gammaRate * epsilonRate
}

func stringToInt(str string) int {
	num, err := strconv.Atoi(str)

	checkErr(err)
	return num
}

func parseIntFromBinaryStringNum(str string) int64 {
	result, err := strconv.ParseInt(str, 2, 64)

	checkErr(err)

	return result
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
