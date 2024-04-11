package main

import (
	"fmt"
	"slices"
)

type Index struct {
	row int
	col int
}

func main() {
	fmt.Println("Hello, World!")
	possibleNumbers := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(possibleNumbers)
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
		col := maxCollIndex
		row := findFirstZeroInLine(sudokuArray, col, false)

		firstZeroIndex := Index{row, col}
		possibilitiesForField := findPossibilitiesForField(sudokuArray, firstZeroIndex, possibleNumbers)
		fmt.Println(possibilitiesForField)

	}
}

func findPossibilitiesForField(sudokuArray [9][9]int, firstZeroIndex Index, possibleNumbers [9]int) [9]int {
	countOutRowPossibilities(sudokuArray, firstZeroIndex, &possibleNumbers)
	countOutCollPossibilities(sudokuArray, firstZeroIndex, &possibleNumbers)
	countOutBoxPossibilities(sudokuArray, firstZeroIndex, &possibleNumbers)
	return possibleNumbers
}

func countOutBoxPossibilities(sudokuArray [9][9]int, firstZeroIndex Index, possibleNumbers *[9]int) {
	row := firstZeroIndex.row
	col := firstZeroIndex.col
	boxRow := row / 3
	boxCol := col / 3

	for i := boxRow * 3; i < boxRow*3+3; i++ {
		for j := boxCol * 3; j < boxCol*3+3; j++ {
			val := sudokuArray[i][j]
			if val != 0 {
				possibleNumbers[val-1] = 0
			}
		}
	}
}

func countOutCollPossibilities(sudokuArray [9][9]int, firstZeroIndex Index, possibleNumbers *[9]int) {
	for i := 0; i < 9; i++ {
		val := sudokuArray[i][firstZeroIndex.col]
		if val != 0 {
			possibleNumbers[val-1] = 0
		}
	}
}

func countOutRowPossibilities(sudokuArray [9][9]int, firstZeroIndex Index, possibleNumbers *[9]int) {
	for i := 0; i < 9; i++ {
		val := sudokuArray[firstZeroIndex.row][i]
		if val != 0 {
			possibleNumbers[val-1] = 0
		}
	}
}

func findFirstZeroInLine(sudokuArray [9][9]int, maxCollIndex int, isRow bool) int {
	for i := 0; i < 9; i++ {
		if isRow {
			if sudokuArray[maxCollIndex][i] == 0 {
				return i
			}
		} else {
			if sudokuArray[i][maxCollIndex] == 0 {
				return i
			}
		}
	}
	return -1
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
