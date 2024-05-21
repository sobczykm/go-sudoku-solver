package print

import "fmt"

func PrintSudoku(sudokuArray [9][9]int) ([]int, []int) {
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
	return solvedRows, solvedColls
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
