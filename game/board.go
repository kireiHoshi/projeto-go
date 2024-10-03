package game

import (
	"strconv"
)

/**
Board Structure, holds * Lines
 */
type Board struct{
	board []Line
}

/**
Creates a new Board for a Game, it receives two arguments for the size
 */
func NewBoard(axisX, axisY int) *Board{
	b := new(Board)
	for x:=0; x<axisX; x++{
			newLine := Line{}
			b.board = append(b.board, newLine.GetNewLine(axisY))
	}
	return b
}

/**
Displays the style of the Board and the current game setting
 */
func (b Board) ShowBoard() string{
	var drawnBoard string

	drawnBoard = addYAxis(drawnBoard, b)
	drawnBoard = addXAxis(drawnBoard, b)
	drawnBoard = addYAxis(drawnBoard, b)

	return drawnBoard
}

/**
Function that fills the numbers of the X axis of the board
 */
func addXAxis(drawnBoard string, b Board) string{
	for i:=0; i < len(b.board); i++{
		drawnBoard = drawnBoard + strconv.Itoa(i) + " "
		for j:=0; j < len(b.board); j++{
			drawnBoard = drawnBoard + b.board[i][j]
		}
		drawnBoard = drawnBoard+ strconv.Itoa(i) + "  \n"
	}
	return drawnBoard
}

/**
Function that fills the numbers of the Y axis of the board
 */
func addYAxis(drawnBoard string, b Board) string{
	drawnBoard = drawnBoard + "+ "
	for roof:=0; roof < len(b.board); roof++{
		drawnBoard = drawnBoard +" "+ strconv.Itoa(roof)+" "
	};
	drawnBoard = drawnBoard + "+\n"
	return drawnBoard
}

/**
Checks for the board and sees if any of the players has won
 */
func (b Board) CheckBoardWinner() string{

	if b.checkDraw(){
		return "DRAW"
	}
	if b.checkLinesCrosses(){
		return "Crosses"
	}
	if b.checkLinesCircles(){
		return "Circles"
	}
	if b.checkDiagonalLineCircles(){
		return "Circles"
	}
	if b.checkDiagonalLineCrosses(){
		return "Crosses"
	}
	if b.checkColumnsCrosses() {
		return "Crosses"
	}
	if b.checkColumnsCircles() {
		return "Circles"
	}
	return ""
}

/**
Function that checks the board and sees if every spot is filled
 */
func (b Board) checkDraw() bool{
	for xAxis:=0; xAxis < len(b.board); xAxis++{
		for yAxis:=0; yAxis < len(b.board); yAxis++ {
			if b.board[xAxis][yAxis] == " - "{
				return false
			}
		}
	}
	return true
}

/**
Function that checks if a line has all crosses
 */
func (b Board) checkLinesCrosses() bool{
	var count int
	for xAxis:=0; xAxis < len(b.board); xAxis++{
		for yAxis:=0; yAxis < len(b.board); yAxis++ {
			if b.board[xAxis][yAxis] == " X "{
				count++
			}
			if count == len(b.board){
				return true
			}
		}
		count = 0
	}
	return false
}

/**
Function that checks if a line has all crosses
 */
func (b Board) checkLinesCircles() bool{
	var count int
	for xAxis:=0; xAxis < len(b.board); xAxis++{
		for yAxis:=0; yAxis < len(b.board); yAxis++ {
			if b.board[xAxis][yAxis] == " O "{
				count++
			}
			if count == len(b.board){
				return true
			}
		}
		count = 0
	}
	return false
}

/**
Function that checks the values of the columns if there's a column of Crosses
 */
func (b Board) checkColumnsCrosses() bool{
	var count int
	for columnCount:=0; columnCount < len(b.board); columnCount++ {
		for xAxis:=0; xAxis < len(b.board); xAxis++ {
			if b.board[xAxis][columnCount] == " X "{
				count++
			}
			if count == len(b.board){
				return true
			}
		}
		count = 0
	}
	return false
}

/*
Function that checks the values of the columns if there's a column of Circles
 */
func (b Board) checkColumnsCircles() bool{
	var count int
	for columnCount:=0; columnCount < len(b.board); columnCount++ {
		for xAxis:=0; xAxis < len(b.board); xAxis++ {
			if b.board[xAxis][columnCount] == " O "{
				count++
			}
			if count == len(b.board){
				return true
			}
		}
		count = 0
	}
	return false
}

/**
Function that checks the left to right diagonal line for crosses
 */
func (b Board) checkDiagonalLineCrosses() bool{
	var count int
	for xAxis:=0; xAxis < len(b.board); xAxis++{
		for yAxis:=0; yAxis < len(b.board); yAxis++ {
			if xAxis == yAxis{
				if b.board[xAxis][yAxis] == " X "{
					count++
				}
				if count == len(b.board) {
					return true
				}
			}
		}
	}
	count=0
	for xAxis:=0; xAxis < len(b.board); xAxis++{
		for yAxis:=0; yAxis < len(b.board); yAxis++ {
			if xAxis+yAxis == len(b.board)-1{
				if b.board[xAxis][yAxis] == " X "{
					count++
				}
				if count == len(b.board) {
					return true
				}
			}
		}
	}
	return false
}

/**
Function that checks the left to right diagonal line for circles
 */
func (b Board) checkDiagonalLineCircles() bool{
	var countLeft int
	for xAxis:=0; xAxis < len(b.board); xAxis++{
		for yAxis:=0; yAxis < len(b.board); yAxis++ {
			if xAxis == yAxis{
				if b.board[xAxis][yAxis] == " O "{
					countLeft++
				}
				if countLeft == len(b.board) {
					return true
				}
			}
		}
	}
	var countRight int
	for xAxis:=0; xAxis < len(b.board); xAxis++{
		for yAxis:=0; yAxis < len(b.board); yAxis++ {
			if xAxis+yAxis == len(b.board)-1{
				if b.board[xAxis][yAxis] == " O "{
					countRight++
				}
				if countRight == len(b.board) {
					return true
				}
			}
		}
	}
	return false
}

/**
Function that checks if the spot is already taken by a player
 */
func (b Board) CheckIfSpotIsTaken(x,y int64) bool{
	if b.board[x][y] == " - " {
		return false
	}
	return true
}

