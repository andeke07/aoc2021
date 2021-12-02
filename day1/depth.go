package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readLinesToIntArray(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		int, err := strconv.Atoi(scanner.Text())
		if err == nil {
			lines = append(lines, int)
		}

	}

	return lines, scanner.Err()
}

func sum(a, b, c int) (sum int) {
	return a + b + c
}

func single_measurements(lines []int) (count int) {
	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			count++
		}
	}
	return count
}

func three_measurements(lines []int) (count int) {
	for i := 3; i < len(lines); i++ {
		if sum(lines[i], lines[i-1], lines[i-2]) > sum(lines[i-1], lines[i-2], lines[i-3]) {
			count++
		}
	}
	return count
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
		os.Exit(1)
	}

	file := os.Args[1]

	lines, err := readLinesToIntArray(file)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	problem_one := single_measurements(lines)
	fmt.Printf("Problem 1 count: %v\n", problem_one)

	problem_two := three_measurements(lines)
	fmt.Printf("Problem 2 count: %v\n", problem_two)

}
