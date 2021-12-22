package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func (p *Position) move2(direction string, steps int) {
	if direction == "forward" {
		p.horizontal += steps
		p.depth += p.aim * steps
	}

	if direction == "down" {
		p.aim += steps
	}

	if direction == "up" {
		p.aim -= steps
	}
}

func Dive2(filepath string) int {
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

		p.move2(direction, steps)
	}

	return p.depth * p.horizontal
}
