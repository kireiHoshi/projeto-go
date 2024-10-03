package main

import (
	"github.com/franciscosaraiva-olx/tictactoe/game"
	"time"
	"math/rand"
)

/**
Initializer
 */
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

/**
Main start of the app
 */
func main() {
	game.Start()
}
