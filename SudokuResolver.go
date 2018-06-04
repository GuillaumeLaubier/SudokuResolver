package main

import "fmt"

type SudokuGrid struct {
  grid [9][9]int
}

func main() {
  grid := SudokuGrid{grid: [9][9]int{}}
  displayGrid(grid)
}


func displayGrid(sudokuGrid SudokuGrid) {
  fmt.Println("-------------------------")

  for i := 0; i < 9; i += 3 {
    for j := 0; j < 9; j += 3 {
      fmt.Println("|", sudokuGrid.grid[i][j], sudokuGrid.grid[i][j + 2], sudokuGrid.grid[i][j + 2], "|", sudokuGrid.grid[i + 1][j], sudokuGrid.grid[i + 1][j + 2], sudokuGrid.grid[i + 1][j + 2], "|", sudokuGrid.grid[i + 2][j], sudokuGrid.grid[i + 2][j + 2], sudokuGrid.grid[i + 2][j + 2], "|")
    }

    fmt.Println("-------------------------")
  }
}
