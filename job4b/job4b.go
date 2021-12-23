package main

import (
	"fmt"
)

func printBoards(boards [][5][5]uint) {
	for i, board := range boards {
		fmt.Println("Board", i)
		printBoard(board)
		fmt.Println("----------")
	}
}
func printBoard(board [5][5]uint) {
	for _, row := range board {
		fmt.Println(row)
	}
}
func printBoardSlice(boardSlice [][]uint) {
	for _, row := range boardSlice {
		fmt.Println(row)
	}
}

func checkBingoForBoards(matched [][5][5]uint, finishedBoards []uint) []int {
	var ret []int
	for i, board := range matched {
		if boardIsFinished(finishedBoards, uint(i)) {
			continue
		}
		if checkBingo(board) {
			ret = append(ret, i)
		}
	}
	return ret
}
func checkBingo(matched [5][5]uint) bool {
	bingo := false
	var matchedTr [5][5]uint

	for i, row := range matched {
		for j, _ := range row {
			matchedTr[i][j] = matched[j][i]
		}
	}

	for _, row := range matched {
		if row == [5]uint{1, 1, 1, 1, 1} {
			bingo = true
		}
	}
	for _, row := range matchedTr {
		if row == [5]uint{1, 1, 1, 1, 1} {
			bingo = true
		}
	}
	return bingo
}

func getUnmarkedSum(board [][]uint, matched [5][5]uint) uint {
	sum := uint(0)
	for i, row := range matched {
		for j, entry := range row {
			if entry == 0 {
				sum += board[i][j]
			}
		}
	}
	return sum
}
func boardIsFinished(finishedIndices []uint, boardIndex uint) bool {
	for _, index := range finishedIndices {
		if index == boardIndex {
			return true
		}
	}
	return false
}

func main() {
	input := []uint{17, 2, 33, 86, 38, 41, 4, 34, 91, 61, 11, 81, 3, 59, 29, 71, 26, 44, 54, 89, 46, 9, 85, 62, 23, 76, 45, 24, 78, 14, 58, 48, 57, 40, 21, 49, 7, 99, 8, 56, 50, 19, 53, 55, 10, 94, 75, 68, 6, 83, 84, 88, 52, 80, 73, 74, 79, 36, 70, 28, 37, 0, 42, 98, 96, 92, 27, 90, 47, 20, 5, 77, 69, 93, 31, 30, 95, 25, 63, 65, 51, 72, 60, 16, 12, 64, 18, 13, 1, 35, 15, 66, 67, 43, 22, 87, 97, 32, 39, 82}
	boards := [][][]uint{[][]uint{
		[]uint{10, 27, 53, 91, 86},
		[]uint{15, 94, 47, 38, 61},
		[]uint{32, 68, 8, 88, 9},
		[]uint{35, 84, 3, 7, 87},
		[]uint{62, 78, 90, 66, 64},
	},
		[][]uint{
			[]uint{30, 51, 26, 16, 57},
			[]uint{66, 88, 47, 75, 23},
			[]uint{61, 77, 64, 9, 73},
			[]uint{44, 32, 28, 80, 81},
			[]uint{3, 99, 67, 49, 78},
		},
		[][]uint{
			[]uint{68, 92, 82, 74, 83},
			[]uint{12, 99, 80, 72, 3},
			[]uint{56, 96, 36, 28, 43},
			[]uint{2, 7, 14, 24, 9},
			[]uint{63, 76, 40, 37, 73},
		},
		[][]uint{
			[]uint{88, 66, 96, 86, 7},
			[]uint{94, 21, 70, 25, 46},
			[]uint{28, 16, 12, 69, 8},
			[]uint{59, 43, 89, 30, 55},
			[]uint{45, 52, 0, 83, 67},
		},
		[][]uint{
			[]uint{21, 42, 92, 30, 81},
			[]uint{15, 98, 26, 79, 48},
			[]uint{90, 99, 5, 88, 53},
			[]uint{2, 67, 74, 55, 33},
			[]uint{54, 20, 69, 39, 75},
		},
		[][]uint{
			[]uint{53, 12, 4, 86, 46},
			[]uint{62, 7, 98, 6, 23},
			[]uint{17, 68, 39, 63, 20},
			[]uint{29, 25, 84, 87, 24},
			[]uint{54, 5, 42, 8, 45},
		},
		[][]uint{
			[]uint{14, 63, 36, 84, 27},
			[]uint{72, 96, 95, 99, 40},
			[]uint{28, 68, 78, 8, 46},
			[]uint{41, 45, 33, 15, 82},
			[]uint{65, 66, 64, 49, 7},
		},
		[][]uint{
			[]uint{22, 35, 72, 75, 47},
			[]uint{53, 59, 17, 95, 55},
			[]uint{25, 91, 57, 10, 96},
			[]uint{39, 3, 18, 90, 64},
			[]uint{34, 26, 71, 52, 69},
		},
		[][]uint{
			[]uint{72, 8, 67, 92, 83},
			[]uint{87, 89, 25, 39, 78},
			[]uint{86, 53, 55, 22, 43},
			[]uint{21, 63, 40, 9, 74},
			[]uint{29, 56, 44, 30, 80},
		},
		[][]uint{
			[]uint{33, 87, 52, 80, 83},
			[]uint{70, 91, 74, 63, 36},
			[]uint{48, 49, 29, 42, 6},
			[]uint{54, 47, 96, 4, 19},
			[]uint{53, 35, 30, 43, 61},
		},
		[][]uint{
			[]uint{82, 7, 38, 86, 79},
			[]uint{53, 87, 21, 45, 44},
			[]uint{10, 18, 46, 30, 36},
			[]uint{12, 1, 50, 2, 59},
			[]uint{94, 3, 39, 62, 32},
		},
		[][]uint{
			[]uint{68, 74, 24, 97, 99},
			[]uint{45, 75, 41, 62, 34},
			[]uint{3, 28, 49, 1, 66},
			[]uint{10, 91, 95, 58, 38},
			[]uint{61, 79, 50, 27, 71},
		},
		[][]uint{
			[]uint{69, 59, 96, 5, 26},
			[]uint{67, 16, 2, 72, 28},
			[]uint{45, 58, 55, 18, 53},
			[]uint{74, 76, 98, 38, 42},
			[]uint{82, 22, 79, 89, 87},
		},
		[][]uint{
			[]uint{3, 33, 73, 66, 52},
			[]uint{69, 29, 78, 75, 34},
			[]uint{1, 64, 15, 17, 68},
			[]uint{27, 32, 46, 54, 18},
			[]uint{55, 74, 60, 28, 40},
		},
		[][]uint{
			[]uint{9, 54, 84, 1, 42},
			[]uint{15, 91, 77, 74, 10},
			[]uint{55, 64, 60, 22, 86},
			[]uint{18, 58, 73, 0, 23},
			[]uint{11, 61, 2, 68, 43},
		},
		[][]uint{
			[]uint{75, 62, 34, 89, 53},
			[]uint{39, 10, 84, 56, 21},
			[]uint{86, 98, 87, 90, 83},
			[]uint{17, 79, 1, 19, 15},
			[]uint{42, 67, 55, 6, 77},
		},
		[][]uint{
			[]uint{36, 3, 60, 1, 70},
			[]uint{63, 40, 7, 88, 61},
			[]uint{65, 96, 18, 73, 30},
			[]uint{42, 35, 44, 45, 81},
			[]uint{77, 95, 39, 24, 5},
		},
		[][]uint{
			[]uint{81, 24, 39, 53, 89},
			[]uint{99, 11, 27, 22, 86},
			[]uint{5, 8, 36, 97, 28},
			[]uint{92, 58, 38, 34, 62},
			[]uint{32, 4, 1, 74, 68},
		},
		[][]uint{
			[]uint{97, 20, 54, 99, 67},
			[]uint{63, 78, 61, 57, 21},
			[]uint{28, 24, 4, 98, 19},
			[]uint{64, 77, 14, 81, 30},
			[]uint{16, 36, 89, 79, 26},
		},
		[][]uint{
			[]uint{73, 90, 0, 28, 5},
			[]uint{11, 27, 56, 96, 1},
			[]uint{29, 87, 12, 69, 8},
			[]uint{63, 95, 72, 86, 64},
			[]uint{48, 46, 50, 37, 57},
		},
		[][]uint{
			[]uint{22, 3, 7, 87, 14},
			[]uint{90, 11, 67, 76, 13},
			[]uint{58, 49, 16, 56, 59},
			[]uint{45, 46, 19, 41, 23},
			[]uint{75, 66, 61, 51, 54},
		},
		[][]uint{
			[]uint{4, 6, 84, 59, 86},
			[]uint{18, 16, 40, 79, 85},
			[]uint{38, 98, 95, 89, 5},
			[]uint{82, 21, 76, 36, 13},
			[]uint{71, 0, 17, 47, 29},
		},
		[][]uint{
			[]uint{73, 41, 26, 87, 95},
			[]uint{62, 99, 58, 9, 20},
			[]uint{45, 10, 71, 28, 39},
			[]uint{89, 17, 29, 46, 81},
			[]uint{49, 35, 24, 74, 32},
		},
		[][]uint{
			[]uint{62, 22, 95, 86, 0},
			[]uint{2, 39, 9, 41, 25},
			[]uint{59, 42, 94, 74, 13},
			[]uint{72, 69, 75, 97, 21},
			[]uint{6, 71, 90, 4, 19},
		},
		[][]uint{
			[]uint{62, 75, 92, 98, 10},
			[]uint{80, 12, 57, 82, 25},
			[]uint{3, 65, 67, 81, 15},
			[]uint{1, 69, 43, 14, 45},
			[]uint{93, 53, 36, 66, 4},
		},
		[][]uint{
			[]uint{72, 12, 47, 40, 78},
			[]uint{68, 43, 24, 28, 99},
			[]uint{5, 98, 70, 25, 59},
			[]uint{8, 10, 58, 46, 7},
			[]uint{36, 56, 37, 84, 32},
		},
		[][]uint{
			[]uint{37, 2, 68, 52, 23},
			[]uint{66, 80, 18, 98, 84},
			[]uint{97, 77, 96, 3, 26},
			[]uint{12, 14, 40, 42, 99},
			[]uint{29, 9, 30, 11, 44},
		},
		[][]uint{
			[]uint{24, 82, 7, 51, 16},
			[]uint{96, 0, 10, 92, 43},
			[]uint{34, 80, 5, 59, 57},
			[]uint{30, 18, 72, 37, 38},
			[]uint{31, 28, 81, 87, 94},
		},
		[][]uint{
			[]uint{40, 93, 85, 27, 69},
			[]uint{70, 6, 41, 14, 17},
			[]uint{58, 95, 79, 24, 65},
			[]uint{62, 48, 11, 78, 43},
			[]uint{30, 21, 19, 16, 97},
		},
		[][]uint{
			[]uint{90, 14, 51, 98, 39},
			[]uint{45, 56, 69, 24, 38},
			[]uint{73, 29, 88, 9, 62},
			[]uint{72, 84, 27, 18, 81},
			[]uint{22, 7, 23, 91, 68},
		},
		[][]uint{
			[]uint{55, 19, 29, 40, 18},
			[]uint{63, 51, 26, 93, 12},
			[]uint{11, 50, 60, 88, 65},
			[]uint{9, 35, 22, 97, 23},
			[]uint{61, 69, 82, 32, 28},
		},
		[][]uint{
			[]uint{37, 17, 81, 94, 1},
			[]uint{19, 6, 0, 49, 8},
			[]uint{40, 25, 34, 98, 63},
			[]uint{59, 15, 53, 23, 64},
			[]uint{66, 52, 69, 84, 68},
		},
		[][]uint{
			[]uint{83, 86, 19, 87, 93},
			[]uint{85, 92, 24, 50, 33},
			[]uint{1, 41, 40, 96, 26},
			[]uint{99, 59, 9, 98, 3},
			[]uint{45, 75, 60, 52, 90},
		},
		[][]uint{
			[]uint{41, 40, 36, 70, 57},
			[]uint{64, 63, 72, 16, 99},
			[]uint{50, 84, 69, 89, 43},
			[]uint{12, 55, 54, 67, 53},
			[]uint{59, 13, 42, 78, 91},
		},
		[][]uint{
			[]uint{98, 19, 96, 21, 39},
			[]uint{28, 48, 83, 50, 97},
			[]uint{57, 7, 12, 6, 63},
			[]uint{38, 32, 52, 66, 10},
			[]uint{2, 18, 42, 75, 94},
		},
		[][]uint{
			[]uint{75, 31, 77, 20, 90},
			[]uint{35, 14, 28, 54, 95},
			[]uint{96, 24, 86, 11, 58},
			[]uint{7, 50, 97, 76, 63},
			[]uint{27, 51, 34, 21, 83},
		},
		[][]uint{
			[]uint{60, 89, 11, 38, 88},
			[]uint{57, 36, 77, 55, 18},
			[]uint{42, 27, 67, 32, 94},
			[]uint{12, 9, 24, 10, 14},
			[]uint{69, 35, 79, 97, 50},
		},
		[][]uint{
			[]uint{46, 82, 60, 45, 6},
			[]uint{84, 88, 0, 7, 51},
			[]uint{37, 52, 64, 25, 74},
			[]uint{31, 8, 75, 53, 72},
			[]uint{11, 47, 34, 40, 50},
		},
		[][]uint{
			[]uint{70, 96, 35, 20, 26},
			[]uint{73, 62, 54, 72, 4},
			[]uint{29, 27, 8, 46, 48},
			[]uint{31, 0, 90, 81, 16},
			[]uint{82, 44, 88, 22, 32},
		},
		[][]uint{
			[]uint{73, 95, 77, 66, 37},
			[]uint{30, 68, 12, 85, 11},
			[]uint{34, 5, 57, 15, 38},
			[]uint{22, 89, 78, 7, 40},
			[]uint{71, 1, 54, 90, 39},
		},
		[][]uint{
			[]uint{13, 80, 22, 73, 30},
			[]uint{49, 36, 98, 75, 33},
			[]uint{32, 95, 74, 54, 56},
			[]uint{21, 55, 68, 34, 61},
			[]uint{60, 50, 3, 38, 11},
		},
		[][]uint{
			[]uint{21, 80, 17, 8, 46},
			[]uint{7, 88, 18, 22, 20},
			[]uint{41, 73, 72, 0, 34},
			[]uint{66, 75, 45, 47, 30},
			[]uint{44, 10, 93, 28, 58},
		},
		[][]uint{
			[]uint{32, 50, 78, 90, 29},
			[]uint{28, 71, 77, 2, 69},
			[]uint{79, 66, 30, 40, 37},
			[]uint{14, 11, 63, 10, 60},
			[]uint{84, 88, 65, 8, 54},
		},
		[][]uint{
			[]uint{86, 89, 64, 69, 76},
			[]uint{53, 82, 24, 16, 51},
			[]uint{67, 75, 3, 33, 21},
			[]uint{23, 63, 99, 13, 43},
			[]uint{4, 39, 7, 73, 87},
		},
		[][]uint{
			[]uint{3, 38, 22, 72, 80},
			[]uint{56, 48, 1, 50, 60},
			[]uint{49, 98, 67, 53, 30},
			[]uint{79, 61, 66, 9, 45},
			[]uint{96, 24, 23, 43, 78},
		},
		[][]uint{
			[]uint{62, 10, 16, 52, 93},
			[]uint{64, 81, 45, 21, 23},
			[]uint{90, 39, 98, 70, 28},
			[]uint{57, 42, 37, 47, 87},
			[]uint{99, 48, 94, 75, 9},
		},
		[][]uint{
			[]uint{69, 91, 72, 58, 67},
			[]uint{13, 16, 52, 86, 68},
			[]uint{17, 40, 23, 15, 83},
			[]uint{80, 37, 85, 82, 60},
			[]uint{22, 76, 3, 89, 35},
		},
		[][]uint{
			[]uint{79, 61, 4, 0, 89},
			[]uint{47, 6, 10, 12, 83},
			[]uint{13, 24, 31, 50, 90},
			[]uint{54, 99, 45, 42, 98},
			[]uint{21, 73, 39, 15, 16},
		},
		[][]uint{
			[]uint{25, 67, 43, 16, 93},
			[]uint{15, 98, 5, 54, 57},
			[]uint{87, 60, 64, 36, 7},
			[]uint{65, 73, 41, 44, 4},
			[]uint{38, 52, 47, 19, 30},
		},
		[][]uint{
			[]uint{22, 20, 1, 92, 94},
			[]uint{52, 73, 90, 14, 16},
			[]uint{54, 59, 29, 9, 44},
			[]uint{65, 83, 89, 75, 45},
			[]uint{72, 33, 77, 15, 69},
		},
		[][]uint{
			[]uint{84, 46, 85, 11, 41},
			[]uint{13, 95, 28, 38, 6},
			[]uint{96, 74, 19, 32, 15},
			[]uint{37, 70, 29, 83, 14},
			[]uint{48, 62, 92, 8, 64},
		},
		[][]uint{
			[]uint{26, 92, 89, 37, 23},
			[]uint{39, 97, 2, 40, 42},
			[]uint{46, 85, 52, 47, 45},
			[]uint{77, 36, 67, 10, 27},
			[]uint{8, 28, 24, 53, 86},
		},
		[][]uint{
			[]uint{52, 21, 54, 91, 72},
			[]uint{96, 53, 17, 89, 51},
			[]uint{23, 58, 5, 18, 2},
			[]uint{13, 68, 32, 47, 75},
			[]uint{50, 97, 30, 84, 86},
		},
		[][]uint{
			[]uint{91, 21, 13, 3, 74},
			[]uint{33, 1, 4, 95, 31},
			[]uint{29, 52, 62, 14, 10},
			[]uint{23, 11, 56, 51, 35},
			[]uint{47, 93, 8, 70, 58},
		},
		[][]uint{
			[]uint{1, 83, 91, 43, 7},
			[]uint{58, 18, 66, 47, 39},
			[]uint{67, 62, 89, 41, 35},
			[]uint{32, 50, 96, 56, 49},
			[]uint{11, 21, 12, 80, 86},
		},
		[][]uint{
			[]uint{23, 3, 63, 99, 42},
			[]uint{98, 97, 66, 86, 60},
			[]uint{73, 32, 96, 52, 75},
			[]uint{8, 31, 59, 84, 19},
			[]uint{93, 48, 35, 0, 92},
		},
		[][]uint{
			[]uint{9, 55, 36, 31, 78},
			[]uint{24, 81, 3, 10, 80},
			[]uint{88, 42, 91, 14, 87},
			[]uint{6, 59, 44, 30, 12},
			[]uint{71, 68, 58, 1, 57},
		},
		[][]uint{
			[]uint{85, 36, 3, 58, 11},
			[]uint{16, 44, 69, 60, 39},
			[]uint{51, 31, 65, 95, 87},
			[]uint{82, 63, 8, 14, 49},
			[]uint{67, 7, 64, 91, 59},
		},
		[][]uint{
			[]uint{52, 65, 60, 39, 22},
			[]uint{1, 77, 81, 91, 46},
			[]uint{19, 18, 87, 31, 88},
			[]uint{23, 11, 32, 10, 79},
			[]uint{4, 50, 8, 59, 68},
		},
		[][]uint{
			[]uint{54, 60, 99, 68, 42},
			[]uint{40, 20, 88, 5, 69},
			[]uint{14, 27, 73, 80, 30},
			[]uint{47, 62, 33, 86, 35},
			[]uint{72, 74, 12, 8, 15},
		},
		[][]uint{
			[]uint{37, 32, 15, 90, 21},
			[]uint{14, 61, 52, 82, 76},
			[]uint{44, 27, 58, 51, 55},
			[]uint{49, 2, 10, 17, 79},
			[]uint{29, 48, 71, 86, 30},
		},
		[][]uint{
			[]uint{36, 63, 48, 89, 92},
			[]uint{38, 71, 1, 46, 41},
			[]uint{3, 83, 79, 14, 34},
			[]uint{51, 11, 96, 69, 35},
			[]uint{61, 74, 99, 22, 95},
		},
		[][]uint{
			[]uint{25, 3, 2, 88, 13},
			[]uint{7, 98, 22, 89, 40},
			[]uint{30, 47, 42, 43, 31},
			[]uint{55, 65, 75, 99, 24},
			[]uint{23, 64, 29, 90, 10},
		},
		[][]uint{
			[]uint{57, 85, 31, 17, 98},
			[]uint{70, 3, 81, 51, 34},
			[]uint{43, 90, 23, 50, 37},
			[]uint{13, 75, 89, 25, 88},
			[]uint{12, 99, 46, 62, 36},
		},
		[][]uint{
			[]uint{97, 48, 96, 15, 53},
			[]uint{45, 87, 35, 0, 77},
			[]uint{46, 72, 89, 55, 54},
			[]uint{98, 81, 69, 92, 42},
			[]uint{95, 47, 19, 33, 63},
		},
		[][]uint{
			[]uint{65, 58, 47, 51, 17},
			[]uint{61, 60, 43, 10, 9},
			[]uint{4, 2, 53, 3, 25},
			[]uint{37, 93, 18, 59, 75},
			[]uint{42, 96, 11, 32, 35},
		},
		[][]uint{
			[]uint{10, 96, 37, 83, 17},
			[]uint{2, 87, 64, 18, 99},
			[]uint{81, 73, 1, 0, 66},
			[]uint{78, 80, 42, 72, 56},
			[]uint{59, 97, 53, 9, 12},
		},
		[][]uint{
			[]uint{59, 97, 49, 11, 58},
			[]uint{6, 99, 83, 27, 12},
			[]uint{21, 67, 79, 16, 57},
			[]uint{96, 9, 39, 69, 81},
			[]uint{18, 43, 42, 45, 95},
		},
		[][]uint{
			[]uint{10, 37, 77, 48, 85},
			[]uint{15, 19, 71, 92, 44},
			[]uint{57, 94, 39, 28, 1},
			[]uint{52, 46, 79, 60, 38},
			[]uint{11, 55, 65, 74, 93},
		},
		[][]uint{
			[]uint{10, 92, 67, 91, 2},
			[]uint{8, 28, 47, 80, 98},
			[]uint{48, 33, 1, 21, 37},
			[]uint{41, 15, 44, 73, 17},
			[]uint{31, 96, 5, 68, 65},
		},
		[][]uint{
			[]uint{87, 55, 85, 48, 7},
			[]uint{10, 53, 42, 80, 84},
			[]uint{81, 91, 68, 54, 27},
			[]uint{32, 45, 67, 76, 34},
			[]uint{30, 62, 31, 72, 12},
		},
		[][]uint{
			[]uint{15, 13, 94, 65, 7},
			[]uint{42, 83, 84, 55, 8},
			[]uint{56, 78, 38, 54, 87},
			[]uint{97, 37, 67, 10, 29},
			[]uint{3, 96, 2, 30, 14},
		},
		[][]uint{
			[]uint{96, 20, 38, 1, 41},
			[]uint{51, 29, 98, 21, 36},
			[]uint{87, 32, 85, 13, 66},
			[]uint{15, 94, 61, 0, 83},
			[]uint{5, 43, 73, 10, 39},
		},
		[][]uint{
			[]uint{74, 19, 4, 13, 53},
			[]uint{31, 92, 66, 40, 39},
			[]uint{42, 3, 21, 33, 95},
			[]uint{14, 34, 23, 45, 60},
			[]uint{16, 82, 89, 44, 7},
		},
		[][]uint{
			[]uint{64, 7, 12, 85, 32},
			[]uint{78, 23, 26, 39, 34},
			[]uint{42, 97, 41, 54, 59},
			[]uint{83, 4, 86, 57, 98},
			[]uint{87, 72, 0, 55, 96},
		},
		[][]uint{
			[]uint{32, 65, 88, 4, 57},
			[]uint{15, 79, 17, 58, 70},
			[]uint{8, 64, 89, 14, 82},
			[]uint{10, 40, 18, 94, 75},
			[]uint{84, 85, 92, 63, 56},
		},
		[][]uint{
			[]uint{19, 95, 11, 31, 38},
			[]uint{15, 0, 82, 75, 13},
			[]uint{25, 67, 78, 59, 18},
			[]uint{99, 69, 57, 21, 81},
			[]uint{14, 63, 12, 85, 35},
		},
		[][]uint{
			[]uint{41, 82, 78, 99, 90},
			[]uint{15, 10, 3, 87, 65},
			[]uint{54, 2, 6, 32, 22},
			[]uint{39, 89, 4, 14, 8},
			[]uint{85, 75, 76, 25, 74},
		},
		[][]uint{
			[]uint{69, 43, 56, 78, 26},
			[]uint{41, 11, 40, 8, 73},
			[]uint{64, 28, 55, 52, 44},
			[]uint{13, 33, 18, 77, 88},
			[]uint{50, 16, 60, 79, 83},
		},
		[][]uint{
			[]uint{34, 24, 48, 22, 11},
			[]uint{74, 60, 61, 42, 26},
			[]uint{37, 89, 84, 53, 7},
			[]uint{38, 41, 43, 31, 69},
			[]uint{17, 64, 88, 52, 14},
		},
		[][]uint{
			[]uint{40, 69, 43, 12, 29},
			[]uint{39, 79, 82, 0, 48},
			[]uint{17, 87, 73, 31, 71},
			[]uint{74, 35, 34, 85, 3},
			[]uint{76, 47, 13, 80, 20},
		},
		[][]uint{
			[]uint{6, 21, 8, 58, 86},
			[]uint{10, 84, 38, 5, 74},
			[]uint{19, 62, 88, 49, 1},
			[]uint{48, 44, 59, 56, 4},
			[]uint{60, 63, 61, 16, 73},
		},
		[][]uint{
			[]uint{40, 99, 77, 5, 11},
			[]uint{63, 30, 68, 94, 39},
			[]uint{36, 66, 13, 47, 89},
			[]uint{70, 22, 18, 53, 96},
			[]uint{24, 56, 87, 4, 93},
		},
		[][]uint{
			[]uint{9, 29, 90, 57, 60},
			[]uint{31, 97, 52, 16, 22},
			[]uint{36, 99, 50, 87, 13},
			[]uint{64, 84, 72, 0, 71},
			[]uint{43, 45, 68, 5, 7},
		},
		[][]uint{
			[]uint{64, 38, 78, 3, 89},
			[]uint{97, 25, 48, 65, 57},
			[]uint{39, 93, 77, 54, 6},
			[]uint{49, 10, 19, 53, 47},
			[]uint{84, 69, 76, 11, 86},
		},
		[][]uint{
			[]uint{88, 86, 29, 33, 72},
			[]uint{14, 93, 40, 36, 59},
			[]uint{19, 71, 47, 17, 91},
			[]uint{92, 16, 67, 27, 55},
			[]uint{51, 15, 2, 5, 84},
		},
		[][]uint{
			[]uint{55, 2, 36, 73, 61},
			[]uint{49, 25, 96, 56, 27},
			[]uint{42, 4, 89, 39, 83},
			[]uint{13, 14, 9, 52, 51},
			[]uint{71, 20, 92, 3, 5},
		},
		[][]uint{
			[]uint{81, 59, 60, 45, 25},
			[]uint{98, 94, 86, 89, 8},
			[]uint{57, 78, 51, 73, 53},
			[]uint{14, 15, 61, 71, 47},
			[]uint{79, 0, 92, 5, 55},
		},
		[][]uint{
			[]uint{67, 45, 73, 55, 53},
			[]uint{27, 88, 35, 85, 60},
			[]uint{71, 24, 6, 23, 21},
			[]uint{82, 76, 3, 9, 22},
			[]uint{86, 78, 8, 44, 47},
		},
		[][]uint{
			[]uint{31, 89, 58, 12, 71},
			[]uint{30, 92, 81, 61, 14},
			[]uint{4, 39, 60, 44, 94},
			[]uint{62, 85, 65, 98, 3},
			[]uint{88, 25, 40, 56, 47},
		},
		[][]uint{
			[]uint{75, 85, 40, 89, 19},
			[]uint{45, 86, 81, 74, 92},
			[]uint{62, 33, 78, 37, 1},
			[]uint{80, 2, 39, 76, 68},
			[]uint{91, 5, 79, 0, 54},
		},
		[][]uint{
			[]uint{56, 45, 33, 86, 47},
			[]uint{63, 73, 96, 15, 95},
			[]uint{69, 85, 22, 80, 20},
			[]uint{51, 43, 64, 0, 58},
			[]uint{3, 6, 29, 52, 74},
		},
		[][]uint{
			[]uint{74, 33, 86, 65, 16},
			[]uint{91, 80, 17, 53, 88},
			[]uint{23, 61, 90, 62, 79},
			[]uint{2, 95, 82, 26, 49},
			[]uint{15, 47, 77, 9, 46},
		},
		[][]uint{
			[]uint{27, 49, 21, 51, 53},
			[]uint{8, 26, 97, 74, 34},
			[]uint{38, 48, 81, 98, 46},
			[]uint{14, 80, 18, 11, 9},
			[]uint{36, 82, 66, 85, 86},
		},
		[][]uint{
			[]uint{44, 7, 89, 5, 45},
			[]uint{29, 48, 93, 37, 41},
			[]uint{77, 67, 21, 68, 81},
			[]uint{96, 28, 38, 49, 58},
			[]uint{19, 80, 1, 0, 50},
		},
		[][]uint{
			[]uint{95, 10, 63, 75, 76},
			[]uint{77, 43, 62, 46, 18},
			[]uint{91, 79, 57, 74, 85},
			[]uint{93, 81, 35, 61, 98},
			[]uint{86, 67, 32, 80, 84},
		},
		[][]uint{
			[]uint{78, 41, 61, 20, 40},
			[]uint{26, 34, 69, 7, 13},
			[]uint{49, 60, 92, 22, 56},
			[]uint{35, 99, 24, 82, 29},
			[]uint{0, 85, 53, 1, 75},
		},
		[][]uint{
			[]uint{19, 18, 55, 70, 84},
			[]uint{28, 68, 71, 20, 6},
			[]uint{27, 90, 86, 52, 2},
			[]uint{44, 43, 15, 48, 39},
			[]uint{14, 37, 63, 83, 75},
		},
		[][]uint{
			[]uint{73, 61, 41, 96, 68},
			[]uint{89, 40, 53, 12, 91},
			[]uint{29, 37, 59, 10, 19},
			[]uint{69, 98, 88, 82, 24},
			[]uint{65, 72, 25, 42, 4},
		},
		[][]uint{
			[]uint{62, 50, 34, 16, 8},
			[]uint{75, 88, 84, 33, 29},
			[]uint{2, 64, 31, 41, 86},
			[]uint{94, 45, 76, 70, 3},
			[]uint{39, 89, 66, 4, 24},
		},
	}

	const numberOfBoards = 100
	if numberOfBoards != len(boards) {
		fmt.Println(len(boards))
		panic("Change number of boards to match the actual number.")
	}
	const lengthOfBoard = 5

	var matched [numberOfBoards][lengthOfBoard][lengthOfBoard]uint
	var finishedBoards []uint

	for vIndex, v := range input {
		fmt.Println("Input", v)
		fmt.Println("**************")
		for boardIndex, board := range boards {
			if boardIsFinished(finishedBoards, uint(boardIndex)) {
				continue
			}
			for rowIndex, row := range board {
				for colIndex, entry := range row {
					if entry == v {
						matched[boardIndex][rowIndex][colIndex] = 1
					}
				}
			}
		}
		bingoIndex := checkBingoForBoards(matched[:], finishedBoards)
		if len(bingoIndex) > 0 {
			for _, bIndex := range bingoIndex {
				finishedBoards = append(finishedBoards, uint(bIndex))

				fmt.Println("input Values", input[:vIndex+1])
				fmt.Println("Bingo", bingoIndex)
				fmt.Println("Winning Board:")
				printBoard(matched[bIndex])
				fmt.Println("Full Board")
				printBoardSlice(boards[bIndex])
				fmt.Println("Result", getUnmarkedSum(boards[bIndex], matched[bIndex])*v)
			}

		}
		if len(finishedBoards) == numberOfBoards {
			break
		}
	}
}