package main

import (
	"encoding/json"
	"net/http"
)

type Sudoku struct{
	Puzzle [][]int `json:"puzzle"`
}

func Solve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sudoku Sudoku
	json.NewDecoder(r.Body).Decode(&sudoku)
	//fmt.Println(sudoku)
	if check(sudoku) {
		helper(sudoku, 0, 0)
		json.NewEncoder(w).Encode(sudoku)
	}else{
		json.NewEncoder(w).Encode("please provide a valid input")
	}
	
}

func check(sudoku Sudoku) bool {
	//if the input has less than 9 rows then return false
	if len(sudoku.Puzzle)!=9 {return false}

	//to check any duplicate number in same row or column
	for i:=0; i<9; i++ {
		if len(sudoku.Puzzle[i])!=9 {return false}
		r := map[int]bool{}
		c := map[int]bool{}
		for j:=0; j<9; j++ {

			//to check if the number is between 0 and 9
			if sudoku.Puzzle[i][j]<0 ||sudoku.Puzzle[i][j]>9 {
				return false
			}

			//to check any duplicate number in same row
			if r[sudoku.Puzzle[i][j]] && sudoku.Puzzle[i][j]!=0 {
				return false
			}else {
				r[sudoku.Puzzle[i][j]] = true
			}

			//to check any duplicate number in same column
			if c[sudoku.Puzzle[j][i]] && sudoku.Puzzle[j][i]!=0 {
				return false
			}else {
				c[sudoku.Puzzle[j][i]] = true
			}
		}
	}

	//to check any duplicate number in same grid
	for r:=0; r<=6; r=r+3 {
		for c:=0; c<=6; c=c+3 {
			g := map[int]bool{}
			for i:=r; i<r+3; i++ {
				for j:=c; j<c+3; j++ {
					if g[sudoku.Puzzle[i][j]] && sudoku.Puzzle[i][j]!=0 {
						return false
					}else {
						g[sudoku.Puzzle[i][j]] = true
					} 
				}
			}
		}
	}

	return true
}

func helper(sudoku Sudoku, row int, col int) bool {
	if row==8 && col==9 {
		return true
	}

	if col==9 {
		return helper(sudoku, row+1, 0)
	}

	if sudoku.Puzzle[row][col]!=0 {
		return helper(sudoku, row, col+1)
	}

	temp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i:=0; i<len(temp); i++ {
		if !present(sudoku, row, col, temp[i]) {
			sudoku.Puzzle[row][col] = temp[i];
			if helper(sudoku, row, col+1){
				return true
			}
			sudoku.Puzzle[row][col] = 0
		}
	}
	return false
}

//func to check if ch is present in same row || col || grid or not
func present(sudoku Sudoku, row int, col int, ch int) bool{
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