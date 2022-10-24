package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := "1.txt"
	input, _ := os.ReadFile(fmt.Sprintf("./inputs/%s", file))

	board := parseInput(string(input))

	display(board)
	fmt.Println()

	fmt.Println("Solving...")
	if solve(&board) {
		display(board)

		file, err := os.Create(fmt.Sprintf("./outputs/%s", file))
		if err != nil {
			log.Println(err)
		}

		for _, slice := range board {
			s := ""
			for _, v := range slice {
				s = s + fmt.Sprintf("%v, ", v)
			}
			_, err = file.WriteString(fmt.Sprintf("%v\n", s))
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// parseInput converts a string to 2D slice.
func parseInput(input string) [][]int {
	board := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	} // dummy board

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			scanner.Scan()
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Println(err)
			}
			board[row][col] = val
		}
	}

	return board
}

// solve the sudoku board using backtracking.
// It returns solved board if the board is solved.
func solve(board *[][]int) bool {
	// tasks to do
	// 1. find the empty cell
	// 2. try to put a number in that cell
	// 3. check if it is safe to put that number in that cell
	// 4. if it is safe then put that number in that cell and solve the rest of the board
	// 5. if it is not safe then try another number
	// 6. if all the numbers are tried and none of them is safe then return false

	row := -1
	col := -1
	isEmpty := false

	// finding the empty cell
	for r, slice := range *board {
		for c, value := range slice {
			if value == 0 {
				isEmpty = true
				row = r
				col = c
				break
			}
			if isEmpty {
				break
			} // found empty cell.
		}
	}

	if !isEmpty {
		return true
	} // base condition

	// if there is no empty cell then the board is solved
	for i := 1; i <= 9; i++ {
		// trying to put a number in that cell

		// checking if it is safe to put that number in that cell
		if isSafe(board, row, col, i) {
			// putting that number in that cell
			(*board)[row][col] = i

			// solving the rest of the board
			if !solve(board) {
				// if the board is not solved then we need to backtrack
				(*board)[row][col] = 0 // bakcktrack
			} else {
				// board it solved !
				return true
			}

		}
	}

	// if all the numbers are tried and none of them is safe then return false
	return false
}

// isSafe returns true if it is possible to put n in that (row, col) position in the given board.
func isSafe(board *[][]int, row int, col int, num int) bool {
	// tasks to do
	// 1. check if num is present in row
	// 2. check if num is present in col
	// 3. check if num is present in 3*3 box

	// checking in the row
	for _, slice := range *board {
		if slice[col] == num {
			return false
		}
	}

	// checking in the col
	for _, value := range (*board)[row] {
		if value == num {
			return false
		}
	}

	// checking in the 3*3 box
	rowStart := row - row%3 // (row / 3) * 3
	colStart := col - col%3 // (col / 3) * 3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if (*board)[i][j] == num {
				return false
			}
		}
	}

	return true
}

// display prints sudoku board on console.
func display(board [][]int) {
	for _, slice := range board {
		for _, val := range slice {
			fmt.Printf("%d, ", val)
		}
		fmt.Println()
	}
}
