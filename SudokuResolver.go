package main

import (
  "fmt"
  "strconv"
)

type SudokuGrid struct {
  grid [9][9]int // [x][y]int x for columns number y for rows number (from top to bottom)
}

func main() {
  easy := [9][9]int{[9]int{0, 4, 6, 0, 2, 3, 0, 7, 9}, [9]int{2, 5, 8, 0, 7, 9, 4, 6, 3}, [9]int{0, 7, 9, 5, 4, 6, 0, 1, 8}, [9]int{4, 0, 2, 3, 0, 5, 6, 0, 7}, [9]int{5, 8, 3, 6, 9, 7, 1, 4, 2}, [9]int{6, 0, 7, 2, 0, 4, 8, 0, 5}, [9]int{7, 2, 0, 9, 6, 8, 3, 5, 0}, [9]int{8, 3, 4, 7, 5, 0, 9, 2, 6}, [9]int{9, 6, 0, 4, 3, 0, 7, 8, 0}}

  grid := SudokuGrid{grid: easy}
  displayGrid(grid)

  fmt.Println(isSquareValid(grid, 0, 0))
  fmt.Println(isRowValid(grid, 4))
  fmt.Println(isColumnValid(grid, 2))
  fmt.Println(isGridValid(grid))
  fmt.Println(isSudokuFinished(grid))
}

/*
 * Resolve section
 */

func resolveSudoku(sudokuGrid SudokuGrid) SudokuGrid {
  
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

  for idxY := y * 3; idxY < y + 3; idxY++ {
    for idxX := x * 3; idxX < x + 3; idxX++ {
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
