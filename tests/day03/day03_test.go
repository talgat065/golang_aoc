package day03__test

import (
	"advent_of_code/day03"
	"fmt"
	"testing"
)

func TestBinaryDiagnostic(t *testing.T) {
	result := day03.BinaryDiagnostic("../../input_files/day03/input.txt")

	var correctAnswer int64 = 1131506

	if result != correctAnswer {
		fmt.Println(result)
		t.Fatal()
	}
}