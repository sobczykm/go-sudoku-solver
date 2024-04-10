package main

import "fmt"

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
	printSudoku(sudokuArray)
}

func printSudoku(sudokuArray [9][9]int) {
	for row := 0; row < 9; row++ {
		printVerticalLine(row)
		numOfSolved := 0
		for col := 0; col < 9; col++ {
			val := sudokuArray[row][col]
			printHorizontalLine(col)

			if val == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print(sudokuArray[row][col], " ")
				numOfSolved++
			}
		}
		fmt.Print("    ", numOfSolved)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	for col := 0; col < 9; col++ {
		numOfSolved := 0
		for row := 0; row < 9; row++ {
			val := sudokuArray[col][row]
			if val != 0 {
				numOfSolved++
			}
		}
		printHorizontalLine(col)
		fmt.Print(numOfSolved, " ")
	}
	fmt.Println()

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
