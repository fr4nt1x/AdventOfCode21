package main

import (
	"fmt"
)

func printGrid(grid [][]uint) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func getGridWidth(lines []Line) [2]uint {
	min := ^uint(0)
	max := uint(0)
	for _, line := range lines {
		if line.start[0] > max {
			max = line.start[0]
		}
		if line.end[0] < min {
			min = line.end[0]
		}
	}
	res := [2]uint{min, max}
	return res
}

func getGridLength(lines []Line) [2]uint {
	min := ^uint(0)
	max := uint(0)
	for _, line := range lines {
		if line.start[1] > max {
			max = line.start[1]
		}
		if line.end[1] < min {
			min = line.end[1]
		}
	}
	res := [2]uint{min, max}
	return res
}
func initializeResult(lines []Line) [][]uint {
	gridWidth := getGridWidth(lines)
	gridLength := getGridLength(lines)
	fmt.Println(gridWidth, gridLength)

	result := make([][]uint, gridWidth[1]+1)
	for i := range result {
		result[i] = make([]uint, gridLength[1]+1)
		for j := range result[i] {
			result[i][j] = 0
		}
	}
	return result
}

func printResult(result [][]uint) {
	sum := 0
	for _, row := range result {
		for _, entry := range row {
			if entry > 1 {
				sum += 1
			}
		}
	}
	fmt.Println("Result:", sum)
}

func main() {
	lines := GetInput()
	result := initializeResult(lines)
	for _, line := range lines {
		direction := []int{0, 0}
		length := uint(0)
		if line.start[0] < line.end[0] {
			direction[0] = 1
			length = line.end[0] - line.start[0]
		} else if line.start[0] > line.end[0] {
			direction[0] = -1
			length = line.start[0] - line.end[0]
		}
		if line.start[1] < line.end[1] {
			direction[1] = 1
			length = line.end[1] - line.start[1]
		} else if line.start[1] > line.end[1] {
			direction[1] = -1
			length = line.start[1] - line.end[1]
		}

		fmt.Println(line)

		for i := 0; i <= int(length); i++ {
			currentIndexX := line.start[0] + uint(direction[0]*i)
			currentIndexY := line.start[1] + uint(direction[1]*i)
			result[currentIndexX][currentIndexY]++
		}
		//printGrid(result)
	}

	printResult(result)
}
