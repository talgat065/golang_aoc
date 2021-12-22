package day02__test

import (
	"advent_of_code/day02"
	"fmt"
	"testing"
)

func TestDive(t *testing.T) {
	result := day02.Dive("../../input_files/day02/input.txt")

	correctAnswer := 1728414

	if result != correctAnswer {
		fmt.Println(result)
		t.Fatal()
	}
}

func TestDive2(t *testing.T) {
	result := day02.Dive2("../../input_files/day02/input.txt")

	correctAnswer := 1765720035

	if result != correctAnswer {
		fmt.Println(result)
		t.Fatal()
	}
}