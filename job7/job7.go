package main

import (
	"fmt"
)

func calculateFuelConsumption(destPos uint, horPositions []uint) uint {
	sum := uint(0)
	for _, pos := range horPositions {

		if pos < destPos {
			sum += destPos - pos
		} else {
			sum += pos - destPos
		}
	}
	return sum
}

func calculateFuelConsumptionVariant(destPos uint, horPositions []uint) uint {
	sum := uint(0)
	for _, pos := range horPositions {
		diff := uint(0)
		if pos < destPos {
			diff = destPos - pos
		} else {
			diff = pos - destPos
		}
		sum += (diff * (diff + 1)) / 2
	}
	return sum
}

func main() {
	horPositions := GetInput()
	min := ^uint(0)
	min2 := ^uint(0)

	for i := uint(0); i < uint(len(horPositions)); i++ {
		res := calculateFuelConsumption(i, horPositions)
		res2 := calculateFuelConsumptionVariant(i, horPositions)
		if res < min {
			min = res
			fmt.Println(min)
		}
		if res2 < min2 {
			min2 = res2
			fmt.Println(min2)
		}
	}
	fmt.Println(min, min2)
}
