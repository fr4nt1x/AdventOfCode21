package main

import (
	"fmt"
	"sort"
	"strings"
)

func printUniqueDigits(horPositions []Digit) {
	digitToLength := map[int]int{
		1: 2,
		4: 4,
		7: 3,
		8: 7,
	}
	sum := 0
	for _, pos := range horPositions {
		for _, out := range pos.output {
			for _, v := range digitToLength {
				if len(out) == v {
					sum++
				}
			}

		}
	}
	fmt.Println(sum)
}
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func findUniqueCharsInFirstString(in1 string, in2 string) string {
	//Which characters are in in1 but not in in2
	s1 := strings.Split(in1, "")
	res := ""
	for _, s := range s1 {
		if !strings.Contains(in2, s) {
			res += s
		}
	}
	return res
}

func getDigit3Index(input []string, one_digit string) int {
	res := -1
	s1 := strings.Split(one_digit, "")
	for i, digit := range input {
		if len(digit) == 5 {
			if hasBothChars(s1, digit) {
				res = i
			}
		}
	}
	return res
}
func hasBothChars(chars []string, in2 string) bool {
	//in2 contains both chars in chars
	sum := 0
	for _, s := range strings.Split(in2, "") {
		for _, c := range chars {
			if c == s {
				sum++
				break
			}
		}
	}
	if sum == 2 {
		return true
	} else {
		return false
	}
}
func decode(horPositions []Digit) {
	digitToLength := map[int]int{
		1: 2,
		4: 4,
		7: 3,
		8: 7,
	}
	//sum := 0
	for _, pos := range horPositions {
		fmt.Println(pos)
		decoding := map[string]string{}

		uniqueStrings := make(map[int]string)
		for _, input := range pos.input {
			for k, v := range digitToLength {
				if len(input) == v {
					uniqueStrings[k] = SortString(input)
				}
			}
		}
		decoding["a"] = findUniqueCharsInFirstString(uniqueStrings[7], uniqueStrings[1])

		fmt.Println(uniqueStrings)
		fmt.Println("Decoding", decoding)
		fmt.Println(pos.input[getDigit3Index(pos.input, uniqueStrings[1])])
	}
}

func main() {
	horPositions := GetInputTest()

	decode(horPositions)
}
