package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	horizontal int
	depth int
}

func (p *Position) move(direction string, steps int) {
	if direction == "forward" {
		p.horizontal += steps
	}

	if direction == "down" {
		p.depth += steps
	}

	if direction == "up" {
		p.depth -= steps
	}
}

func Dive(filepath string) int {
	f, err := os.Open(filepath)

	checkErr(err)

	defer closeFile()(f)

	scr := bufio.NewScanner(f)

	p := Position{}

	for scr.Scan() {
		text := scr.Text()
		sep := " "

		parts := strings.Split(text, sep)
		direction := parts[0]

		steps, err := strconv.Atoi(parts[1])

		checkErr(err)

		p.move(direction, steps)
	}

	return p.depth * p.horizontal
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
