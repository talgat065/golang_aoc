package day01__test

import (
	"advent_of_code/day01"
	"fmt"
	"testing"
)

func TestSonarSweep(t *testing.T) {
	increases := day01.SonarSweep("../../input_files/day01/input.txt")

	correctAnswer := 1121
	if increases != correctAnswer {
		fmt.Println(increases)
		t.Fatal()
	}
}