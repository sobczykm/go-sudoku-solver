package main

import (
	"fmt"
	"slices"
	"sobczykm/sudoku-solver/count"
	"sobczykm/sudoku-solver/print"
)

var possibleNumbers = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

var sudokuArray = [9][9]int{
	{3, 0, 0, 9, 2, 0, 0, 0, 6},
	{1, 0, 5, 0, 0, 0, 0, 9, 0},
	{0, 0, 3, 0, 0, 5, 0, 0, 7},
	{0, 0, 7, 5, 0, 9, 0, 3, 0},
	{6, 0, 9, 8, 3, 1, 0, 7, 0},
	{5, 0, 1, 6, 0, 0, 0, 0, 8},
	{4, 9, 0, 1, 0, 0, 0, 8, 3},
	{2, 1, 0, 0, 9, 0, 0, 0, 4},
	{7, 0, 3, 2, 8, 0, 0, 0, 9},
}

func main() {
	solveSudoku(sudokuArray)
}

type Possibility struct {
	index              count.Index
	possibilities      [9]int
	numOfPossibilities int
}

type LineInfo struct {
	lineIndex int
	lineValue int
}

func solveSudoku(sudokuArray [9][9]int) {
	rows, colls := print.PrintSudoku(sudokuArray)

	rowInfo, collInfo := getMaxIndexes(rows, colls)
	printLinesInfo(rowInfo, collInfo)
	//pick line with bigger value
	var possibilities []Possibility

	if rowInfo.lineValue > collInfo.lineValue {
		fmt.Println("Solving row with index: ", collInfo.lineIndex)
		zerosInLine := count.FindZerosInLine(sudokuArray, collInfo.lineIndex, false)
		fmt.Println("Zeros in line: ", zerosInLine)
		row := rowInfo.lineIndex

		for _, col := range zerosInLine {
			zeroIndex := count.Index{Row: row, Col: col}
			fmt.Println("First zero index: ", zeroIndex)

			possibilitiesForField := count.FindPossibilitiesForField(sudokuArray, zeroIndex, possibleNumbers)

			numOfPossibilities := getNumOfPossibilities(possibilitiesForField)
			possibilities = append(possibilities, Possibility{zeroIndex, possibilitiesForField, numOfPossibilities})
		}
	} else {
		fmt.Println("Solving coll with index: ", collInfo.lineIndex)
		zerosInLine := count.FindZerosInLine(sudokuArray, collInfo.lineIndex, false)
		fmt.Println("Zeros in line: ", zerosInLine)
		col := collInfo.lineIndex

		for _, row := range zerosInLine {
			zeroIndex := count.Index{Row: row, Col: col}
			fmt.Println("First zero index: ", zeroIndex)

			possibilitiesForField := count.FindPossibilitiesForField(sudokuArray, zeroIndex, possibleNumbers)

			numOfPossibilities := getNumOfPossibilities(possibilitiesForField)
			possibilities = append(possibilities, Possibility{zeroIndex, possibilitiesForField, numOfPossibilities})
		}
	}
	fmt.Println("Possibilities: ", possibilities)
}

func getNumOfPossibilities(possibilitiesForField [9]int) int {
	var numOfPossibilities int
	for _, possibility := range possibilitiesForField {
		if possibility != 0 {
			numOfPossibilities++
		}
	}
	return numOfPossibilities
}

func getMaxIndexes(rows []int, colls []int) (LineInfo, LineInfo) {
	maxRowValue := slices.Max(rows)
	maxCollValue := slices.Max(colls)

	maxRowIndex := slices.Index(rows, maxRowValue)
	maxCollIndex := slices.Index(colls, maxCollValue)

	return LineInfo{maxRowIndex, maxRowValue}, LineInfo{maxCollIndex, maxCollValue}
}

func printLinesInfo(rowInfo, collInfo LineInfo) {
	fmt.Println("Max row value: ", rowInfo.lineValue, "with index: ", rowInfo.lineIndex)
	fmt.Println("Max coll value: ", collInfo.lineValue, "with index: ", collInfo.lineIndex)
}
