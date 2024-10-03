package game

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"math/rand"
	"strconv"
	"time"
)

/**
Player Structure, has a name and number of points.
 */
type Player struct{
	name string
}

/**
List of Players in a game.
 */
var players []*Player
/**
Number of turns in a game.
 */
var turns int
/**
Winner of a game.
 */
var winner string
/**
Board of a game.
 */
var board *Board

/**
Starting function to run a game.
 */
func Start(){
	clearScreen()
	fmt.Println("##### TIC TAC TOE #####")
	menu() // assign the game objects
	runGame() // run the game
}

/**
Function that displays the game menu to start playing or to leave the game.
 */
func menu(){
	fmt.Println("1. Play")
	fmt.Println("0. Exit")

	var option int
	fmt.Scanf("%d\n", &option)

	clearScreen()
	switch option{
		case 1: gameStart()
				break
		default: os.Exit(0)
				break
	}
}

/**
Function that starts the game configuration and options.
 */
func gameStart(){
	fmt.Println("Tic Tac Toe Setup!")

	fmt.Printf("Player 1 name[Cross]: ")
	player1 := createPlayer()
	fmt.Printf("Player 2 name[Circles]: ")
	player2 := createPlayer()

	players = append(players, player1, player2)
	board = createBoard()
}

/**
Function that creates a new player
 */
func createPlayer() *Player{
	reader := bufio.NewReader(os.Stdin)

	playerName, _ := reader.ReadString('\n')
	playerName = strings.TrimSpace(playerName)

	return &Player{playerName}
}

/**
Function that creates a new board
 */
func createBoard() *Board{
	var axisX, axisY, valid int
	for valid==0{
		fmt.Printf("Dimensions? (X by X) 3 Min, 10 Max: ")
		fmt.Scanf("%d\n", &axisX)
		axisY = axisX
		if axisX>=3 && axisX<=10 {
			valid = 1
		} else {
			fmt.Println("Invalid value for the board! Try again")
			time.Sleep(2000*time.Millisecond)
		}
	}
	return NewBoard(axisX, axisY)
}

/**
Function that runs the game.
 */
func runGame(){
	var turn int
	turn = rand.Intn(2)
	for winner == ""{
		clearScreen() // clears the screen
		showPlayerInfo() //shows player info
		fmt.Println() //break line
		fmt.Println(board.ShowBoard()) //shows the state of the board
		fmt.Println("Turn number "+strconv.Itoa(turns))
		if turn == 1 {
			fmt.Println("Current play -> Cross")
			turn = 0
		} else if turn == 0 {
			fmt.Println("Current play -> Circle")
			turn = 1
		}
		play(turn)
		winner = board.CheckBoardWinner() //checks if there's a winner
		turns++
	}
	fmt.Println(board.ShowBoard())
	announceWinner(winner)
}

/**
Function that displays Player information, their names and points
 */
func showPlayerInfo(){
	fmt.Printf("## (Cross): %s | %s :(Circle) ##\n", players[0].name, players[1].name)
}

/**
Function that plays a move for a player
 */
func play(turn int){
	reader := bufio.NewReader(os.Stdin)
	var valid int
	for valid==0{

		fmt.Println("What region to play on?")
		fmt.Print("x-y -> ")
		region, _ := reader.ReadString('\n')
		region = strings.TrimSpace(region)

		x, _ := strconv.ParseInt(region[0:1], 10, 10)
		y, _ := strconv.ParseInt(region[len(region)-1:], 10,10)

		if  x >= 0 && y >= 0 && x <= int64(len(board.board)) && y <= int64(len(board.board)) {
			if turn == 0{
				if board.CheckIfSpotIsTaken(x, y){
					fmt.Println("Region is already taken")
					play(turn)
				} else {
					board.board[x][y] = " X "
				}
			}
			if turn == 1{
				if board.CheckIfSpotIsTaken(x, y){
					fmt.Println("Region is already taken")
					play(turn)
				} else {
					board.board[x][y] = " O "
				}
			}
			valid=1
		} else {
			fmt.Println("Wrong range selected, try again")
			time.Sleep(2000*time.Millisecond)
		}
	}
}

/**
Function that announces the winner of the game
 */
func announceWinner(winner string){
	if winner == "Crosses"{
		winner = players[0].name
	}
	if winner == "Circles"{
		winner = players[1].name
	}
	if winner == "Draw"{
		winner = ">> DRAW <<"

	}
	fmt.Println("Result of the game : "+winner)
}

/**
Function that clears the console screen for better display
 */
func clearScreen(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
