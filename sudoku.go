package main

import (
	"encoding/json"
	"net/http"
)

type Sudoku struct{
	Puzzle [][]string `json:"puzzle"`
}

func Solve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sudoku Sudoku
	json.NewDecoder(r.Body).Decode(&sudoku)
	//fmt.Println(sudoku)
	helper(sudoku, 0, 0)
	json.NewEncoder(w).Encode(sudoku)
}

func helper(sudoku Sudoku, row int, col int) bool {
	if row==8 && col==9 {
		return true
	}

	if col==9 {
		return helper(sudoku, row+1, 0)
	}

	if sudoku.Puzzle[row][col]!="." {
		return helper(sudoku, row, col+1)
	}

	temp := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i:=0; i<len(temp); i++ {
		if !present(sudoku, row, col, temp[i]) {
			sudoku.Puzzle[row][col] = temp[i];
			if helper(sudoku, row, col+1){
				return true
			}
			sudoku.Puzzle[row][col] = "."
		}
	}
	return false
}

func present(sudoku Sudoku, row int, col int, ch string) bool{
	for i:=0; i<9; i++ {
		if sudoku.Puzzle[i][col]==ch {
			return true
		}
		if sudoku.Puzzle[row][i]==ch {
			return true
		}
	}

	r := 3*(row/3)
	c := 3*(col/3)

	for i:=r; i<r+3; i++ {
		for j:=c; j<c+3; j++ {
			if sudoku.Puzzle[i][j]==ch {
				return true
			}
		}
	}

	return false
}