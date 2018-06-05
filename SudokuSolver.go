package main

import (
  "fmt"
  "strconv"
)

type SudokuGrid struct {
  grid [9][9]int // [x][y]int x for columns number y for rows number (from top to bottom)
}

func main() {
  //easy := [9][9]int{[9]int{0, 0, 4, 0, 9, 0, 8, 0, 0}, [9]int{7, 0, 9, 1, 0, 0, 0, 0, 0}, [9]int{1, 0, 0, 0, 2, 0, 5, 0, 7}, [9]int{0, 3, 0, 9, 0, 0, 0, 6, 0}, [9]int{9, 0, 0, 0, 0, 0, 0, 0, 4}, [9]int{0, 6, 0, 0, 0, 8, 0, 7, 0}, [9]int{8, 0, 7, 0, 6, 0, 0, 0, 3}, [9]int{0, 0, 0, 0, 0, 2, 7, 0, 5}, [9]int{0, 0, 5, 0, 3, 0, 6, 0, 0}}

  hard := [9][9]int{[9]int{0, 7, 0, 0, 5, 0, 0, 8, 0}, [9]int{0, 8, 0, 0, 0, 4, 0, 0, 0}, [9]int{0, 9, 0, 7, 8, 0, 3, 0, 0}, [9]int{8, 0, 0, 0, 7, 0, 2, 0, 0}, [9]int{0, 1, 0, 0, 0, 0, 0, 7, 0}, [9]int{0, 0, 6, 0, 9, 0, 0, 0, 1}, [9]int{0, 0, 1, 0, 3, 2, 0, 4, 0}, [9]int{0, 0, 0, 5, 0, 0, 0, 3, 0}, [9]int{0, 6, 0, 0, 4, 0, 0, 9, 0}}
  grid := SudokuGrid{grid: hard}
  displayGrid(grid)

  isResolved, resolvedGrid := resolveSudoku(grid)
  fmt.Println("Is resolved:", isResolved)
  displayGrid(resolvedGrid)
}

/*
 * Resolve section
 */

func resolveSudoku(sudokuGrid SudokuGrid) (bool, SudokuGrid) {
  x, y := nextEmptyIndex(sudokuGrid)
  return putValue(1, sudokuGrid, x, y)
}

func putValue(value int, sudokuGrid SudokuGrid, x int, y int) (bool, SudokuGrid) {
  if value > 9 {
    return false, sudokuGrid
  }

  if x == -1 && y == -1 {
    return true, sudokuGrid
  }

  sudokuGrid.grid[x][y] = value

  if isGridValid(sudokuGrid) {
    nextX, nextY := nextEmptyIndex(sudokuGrid)
    isValid, tmpGrid := putValue(1, sudokuGrid, nextX, nextY)
    if isValid {
      return isValid, tmpGrid
    }
  }
  return putValue(value + 1, sudokuGrid, x, y)

}



func nextEmptyIndex(sudokuGrid SudokuGrid) (int, int) {
  for idxX := 0; idxX < 9; idxX++ {
    for idxY := 0; idxY < 9; idxY++ {
      if sudokuGrid.grid[idxX][idxY] == 0 {
        return idxX, idxY
      }
    }
  }

  // There is no more empty index
  return -1, -1
}

/*
 * Display section
 */

func displayGrid(sudokuGrid SudokuGrid) {
  fmt.Println("-------------------------------")

  for y := 0; y < 9; y++ {
    fmt.Print("|")

    for x := 0; x < 9; x++ {
      fmt.Print(" ", prettyDisplay(sudokuGrid.grid[x][y]), " ")

      if (x + 1) % 3 == 0 {
        fmt.Print("|")
      }
    }
    fmt.Println("")

    if (y + 1) % 3 == 0 {
      fmt.Println("-------------------------------")
    }
  }
}

func prettyDisplay(i int) string {
  if (i == 0) {
    return " "
  }

  return strconv.Itoa(i)
}


/*
 * Check section
 */

func isSudokuFinished(sudokuGrid SudokuGrid) bool {
  emptyCaseLeft := false

  for x := 0; x < 9; x++ {
    for y := 0; y < 9; y++ {
      if (sudokuGrid.grid[x][y] == 0) {
        emptyCaseLeft = true
      }
    }
  }

  return !emptyCaseLeft && isGridValid(sudokuGrid)
}

func isGridValid(sudokuGrid SudokuGrid) bool {
  for x := 0; x < 3; x++ {
    for y := 0; y < 3; y++ {
      if !isSquareValid(sudokuGrid, x, y) {
        return false
      }
    }
  }

  for idx := 0; idx < 9; idx++ {
    if !isRowValid(sudokuGrid, idx) || !isColumnValid(sudokuGrid, idx) {
      return false
    }
  }

  return true
}

func isSubGridValid(subGrid [9]int) bool {
  cptValue := [9]int{}

  for i := 0; i < 9; i++ {
    if subGrid[i] != 0 {
      cptValue[subGrid[i] - 1]++

      if cptValue[subGrid[i] - 1] > 1 {
        return false
      }

    }
  }

  return true
}

func isSquareValid(sudokuGrid SudokuGrid, x int, y int) bool {
  if x >= 3 || y >= 3 {
    return false
  }

  subGrid := [9]int{}
  idx := 0

  for idxY := y * 3; idxY < y * 3 + 3; idxY++ {
    for idxX := x * 3; idxX < x * 3 + 3; idxX++ {
      subGrid[idx] = sudokuGrid.grid[idxX][idxY]
      idx++
    }
  }

  return isSubGridValid(subGrid)
}

func isRowValid(sudokuGrid SudokuGrid, y int) bool {
  if y >= 9 {
    return false
  }

  subGrid := [9]int{}

  for idx := 0; idx < 9; idx++ {
    subGrid[idx] = sudokuGrid.grid[idx][y]
  }

  return isSubGridValid(subGrid)
}

func isColumnValid(sudokuGrid SudokuGrid, x int) bool {
  if x >= 9 {
    return false
  }

  subGrid := [9]int{}

  for idx := 0; idx < 9; idx++ {
    subGrid[idx] = sudokuGrid.grid[x][idx]
  }

  return isSubGridValid(subGrid)
}
