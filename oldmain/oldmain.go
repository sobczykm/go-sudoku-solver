package oldmain

import (
	"fmt"
	"slices"
	"sobczykm/sudoku-solver/print"
	"time"
)

type Index struct {
	row int
	col int
}

type Possibility struct {
	index              Index
	possibilities      [9]int
	numOfPossibilities int
}

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
	fmt.Println("Hello, World!")

	fmt.Println("Sudoku before solving:")
	solveSudoku(sudokuArray)
}

func solveSudoku(sudokuArray [9][9]int) {
	rows, colls := print.PrintSudoku(sudokuArray)

	var unsolvedRows []int
	var unsolvedColls []int

	for _, row := range rows {
		if row != 9 {
			unsolvedRows = append(unsolvedRows, row)
		} else {
			unsolvedRows = append(unsolvedRows, 0)
		}
	}

	for _, coll := range colls {
		if coll != 9 {
			unsolvedColls = append(unsolvedColls, coll)
		} else {
			unsolvedColls = append(unsolvedColls, 0)
		}
	}

	fmt.Println("Unsolved rows: ", unsolvedRows)
	fmt.Println("Unsolved colls: ", unsolvedColls)

	maxRow := slices.Max(unsolvedRows)
	maxColl := slices.Max(unsolvedColls)
	fmt.Println("Max row: ", maxRow)
	fmt.Println("Max coll: ", maxColl)

	maxRowIndex := slices.Index(unsolvedRows, maxRow)
	maxCollIndex := slices.Index(unsolvedColls, maxColl)

	var possibilities []Possibility

	if maxRow > maxColl {
		doRow(maxRowIndex, sudokuArray, &possibilities)
	} else {
		doColl(maxCollIndex, sudokuArray, &possibilities)
	}

	fmt.Println(possibilities)
	var possibilityWithOnePossibility Possibility
	for _, possibility := range possibilities {
		if possibility.numOfPossibilities == 1 {
			possibilityWithOnePossibility = possibility
			break
		}
	}

	if possibilityWithOnePossibility == (Possibility{}) {
		if maxRow > maxColl {
			doColl(maxCollIndex, sudokuArray, &possibilities)
		} else {
			doRow(maxRowIndex, sudokuArray, &possibilities)
		}

		panic("No possibility with one possibility found")
	}

	var value int
	//return first non zero value
	for _, possibility := range possibilityWithOnePossibility.possibilities {
		if possibility != 0 {
			value = possibility
			break
		}
	}
	fmt.Println("Value: ", value)
	fmt.Println("Index: ", possibilityWithOnePossibility.index)
	sudokuArray[possibilityWithOnePossibility.index.row][possibilityWithOnePossibility.index.col] = value

	time.Sleep(500 * time.Millisecond)
	solveSudoku(sudokuArray)
}

func doColl(maxCollIndex int, sudokuArray [9][9]int, possibilities *[]Possibility) {
	fmt.Println("Max col is bigger than max ROW")

	col := maxCollIndex
	zerosInLine := findZerosInLine(sudokuArray, col, false)

	fmt.Println(zerosInLine)

	for _, row := range zerosInLine {
		firstZeroIndex := Index{row, col}
		possibilitiesForField := findPossibilitiesForField(sudokuArray, firstZeroIndex, possibleNumbers)
		var numOfPossibilities int
		for _, possibility := range possibilitiesForField {
			if possibility != 0 {
				numOfPossibilities++
			}
		}
		*possibilities = append(*possibilities, Possibility{firstZeroIndex, possibilitiesForField, numOfPossibilities})
	}

}

func doRow(maxRowIndex int, sudokuArray [9][9]int, possibilities *[]Possibility) {
	fmt.Println("Max row is bigger than max COLL")
	row := maxRowIndex
	zerosInLine := findZerosInLine(sudokuArray, row, false)

	fmt.Println(zerosInLine)

	for _, col := range zerosInLine {
		firstZeroIndex := Index{row, col}
		possibilitiesForField := findPossibilitiesForField(sudokuArray, firstZeroIndex, possibleNumbers)
		var numOfPossibilities int
		for _, possibility := range possibilitiesForField {
			if possibility != 0 {
				numOfPossibilities++
			}
		}
		*possibilities = append(*possibilities, Possibility{firstZeroIndex, possibilitiesForField, numOfPossibilities})
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

func findZerosInLine(sudokuArray [9][9]int, lineIndex int, isRow bool) []int {
	zeroIndexes := make([]int, 0)
	for i := 0; i < 9; i++ {
		if isRow {
			if sudokuArray[lineIndex][i] == 0 {
				zeroIndexes = append(zeroIndexes, i)
			}
		} else {
			if sudokuArray[i][lineIndex] == 0 {
				zeroIndexes = append(zeroIndexes, i)
			}
		}
	}
	return zeroIndexes
}
