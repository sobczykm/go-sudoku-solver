package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Hello, World!")
	sudokuArray := [9][9]int{
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

	fmt.Println("Sudoku before solving:")
	rows, colls := printSudoku(sudokuArray)

	maxRow := slices.Max(rows)
	maxColl := slices.Max(colls)

	maxRowIndex := slices.Index(rows, maxRow)
	maxCollIndex := slices.Index(colls, maxColl)

	if maxRow > maxColl {
		fmt.Println(maxRowIndex)
	} else {
		fmt.Println(maxCollIndex)
		findFirstZeroInLine(sudokuArray, maxCollIndex, false)
	}
}

func findFirstZeroInLine(sudokuArray [9][9]int, maxCollIndex int, false bool) {

}

func printSudoku(sudokuArray [9][9]int) ([]int, []int) {
	var solvedRows []int
	var solvedColls []int
	for row := 0; row < 9; row++ {
		printVerticalLine(row)
		for col := 0; col < 9; col++ {
			val := sudokuArray[row][col]
			printHorizontalLine(col)

			if val == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print(val, " ")
			}
		}
		solvedInLine := countSolvedCellsInLine(sudokuArray, true, row)
		solvedRows = append(solvedRows, solvedInLine)
		fmt.Print("    ", solvedInLine)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	for col := 0; col < 9; col++ {
		solvedInLine := countSolvedCellsInLine(sudokuArray, false, col)
		solvedColls = append(solvedColls, solvedInLine)
		printHorizontalLine(col)
		fmt.Print(solvedInLine, " ")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("SolvedRows: ", solvedRows)
	fmt.Println("SolvedColls: ", solvedColls)
	return solvedRows, solvedColls
}

func printVerticalLine(i int) {
	if i != 0 && i%3 == 0 {
		fmt.Println("---------------------")
	}
}

func printHorizontalLine(j int) {
	if j != 0 && j%3 == 0 {
		fmt.Print("| ")
	}
}

func countSolvedCellsInLine(sudokuArray [9][9]int, isRow bool, j int) int {
	count := 0
	for i := 0; i < 9; i++ {
		val := 0
		if isRow {
			val = sudokuArray[j][i]
		} else {
			val = sudokuArray[i][j]
		}
		if val != 0 {
			count++
		}
	}
	return count
}

func findPossibilitiesForField(sudokuArray [9][9]int, row int, col int) []int {
	panic("Not implemented")
}
