package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var verticalPosition int = 0
var horizontalPosition int = 0
var aim = 0

func readLinesToIntArray(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func processInstruction(instruction string, count int) {
	switch instruction {
	case "forward":
		horizontalPosition += count
		break
	case "up":
		verticalPosition -= count
		break
	case "down":
		verticalPosition += count
		break
	}

}

func processInstructionWithAim(instruction string, count int) {
	switch instruction {
	case "forward":
		horizontalPosition += count
		verticalPosition += (aim * count)
		break
	case "up":
		aim -= count
		break
	case "down":
		aim += count
		break
	}

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

	for i := 0; i < len(lines); i++ {
		command := strings.Split(lines[i], " ")
		int, _ := strconv.Atoi(command[1])
		processInstruction(command[0], int)
	}
	fmt.Printf("Excercise 1:\n")
	fmt.Printf("Vertical Position: %v\nHorizontal Position: %v\nCombined Position: %v\n", verticalPosition, horizontalPosition, (verticalPosition * horizontalPosition))

	verticalPosition = 0
	horizontalPosition = 0

	for i := 0; i < len(lines); i++ {
		command := strings.Split(lines[i], " ")
		int, _ := strconv.Atoi(command[1])
		processInstructionWithAim(command[0], int)
	}

	fmt.Printf("Excercise 2:\n")
	fmt.Printf("Vertical Position: %v\nHorizontal Position: %v\nCombined Position: %v\n", verticalPosition, horizontalPosition, (verticalPosition * horizontalPosition))
}
