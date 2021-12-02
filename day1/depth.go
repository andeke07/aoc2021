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

	count := 0

	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			count++
			fmt.Printf("i: %v, depth: %v, depth-1: %v, count: %v\n", i, lines[i], lines[i-1], count)
		}
	}

	fmt.Printf("Count: %v", count)
}
