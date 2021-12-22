package main

import "fmt"
import "advent_of_code/day01"

func main() {
	fmt.Println("hello")

	increases := day01.SonarSweep("input_files/day01/input.txt")

	fmt.Printf("Result is: %d", increases)
}
