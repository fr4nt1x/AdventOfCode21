package main

import "fmt"

const NUMBER_STATES = uint(9)
const STATE_RESET_FISH = uint(6)

func calculate(inputState [NUMBER_STATES]uint, days uint) [NUMBER_STATES]uint {
	var state [NUMBER_STATES]uint
	state = inputState
	for i := uint(0); i < days; i++ {
		state0 := state[0]
		for j := uint(1); j < NUMBER_STATES; j++ {
			state[j-1] = state[j]
		}
		state[NUMBER_STATES-1] = state0
		state[STATE_RESET_FISH] += state0
	}
	return state
}

func getSum(state [NUMBER_STATES]uint) {
	sum := uint(0)
	for _, i := range state {
		sum += i
	}
	fmt.Println(sum)
}
func getInputState(fishes []uint) [NUMBER_STATES]uint {
	res := [NUMBER_STATES]uint{0}
	for _, fish := range fishes {
		res[fish]++
	}
	return res
}

func main() {
	fishes := GetInput()
	inputState := getInputState(fishes)
	getSum(calculate(inputState, 80))
	getSum(calculate(inputState, 256))
}
