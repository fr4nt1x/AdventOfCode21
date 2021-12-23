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
		indexSame := 0
		if line.start[0] == line.end[0] {
			indexSame = 0
		} else if line.start[1] == line.end[1] {
			indexSame = 1
		} else {
			continue
		}
		fmt.Println(line)
		indexOther := (indexSame + 1) % 2
		start := uint(0)
		end := uint(0)
		if line.start[indexOther] < line.end[indexOther] {
			start = line.start[indexOther]
			end = line.end[indexOther]
		} else {
			start = line.end[indexOther]
			end = line.start[indexOther]
		}

		for i := start; i <= end; i++ {
			if indexSame == 0 {
				result[line.start[0]][i] += 1
			} else {
				result[i][line.start[1]] += 1
			}
		}
	}
	printResult(result)
}
