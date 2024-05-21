package count

type Index struct {
	Row int
	Col int
}

func FindPossibilitiesForField(sudokuArray [9][9]int, firstZeroIndex Index, possibleNumbers [9]int) [9]int {
	countOutRowPossibilities(sudokuArray, firstZeroIndex, &possibleNumbers)
	countOutCollPossibilities(sudokuArray, firstZeroIndex, &possibleNumbers)
	countOutBoxPossibilities(sudokuArray, firstZeroIndex, &possibleNumbers)
	return possibleNumbers
}

func countOutBoxPossibilities(sudokuArray [9][9]int, firstZeroIndex Index, possibleNumbers *[9]int) {
	row := firstZeroIndex.Row
	col := firstZeroIndex.Col
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
		val := sudokuArray[i][firstZeroIndex.Col]
		if val != 0 {
			possibleNumbers[val-1] = 0
		}
	}
}

func countOutRowPossibilities(sudokuArray [9][9]int, firstZeroIndex Index, possibleNumbers *[9]int) {
	for i := 0; i < 9; i++ {
		val := sudokuArray[firstZeroIndex.Row][i]
		if val != 0 {
			possibleNumbers[val-1] = 0
		}
	}
}

func FindZerosInLine(sudokuArray [9][9]int, lineIndex int, isRow bool) []int {
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
