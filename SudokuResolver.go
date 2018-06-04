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
}


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

func isSubGridValid(subGrid [9]int) bool {
  cptValue := [9]int{}

  for i := 0; i < 9; i++ {
    if subGrid[i] != 0 {
      cptValue[i - 1]++

      if cptValue[i - 1] > 1 {
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

func isRowValid(sudokuGrid SudokuGrid, x int, y int) bool {
  return isSubGridValid(getRowForIndex(sudokuGrid, x, y))
}

func getRowForIndex(sudokuGrid SudokuGrid, x int, y int) [9]int {
  return [9]int{}
}

func isColumnValid(sudokuGrid SudokuGrid, x int, y int) bool {
 return false
}

func getColumnForIndex(sudokuGrid SudokuGrid, x int, y int) [9]int {
  return [9]int{}
}
