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