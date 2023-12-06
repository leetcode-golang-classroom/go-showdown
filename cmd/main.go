package main

import (
	"github.com/leetcode-golang-classroom/go-showdown/showdown"
)

func main() {
	hp := showdown.NewHumanPlayer()
	hp1 := showdown.NewAIPlayer()
	hp2 := showdown.NewAIPlayer()
	hp3 := showdown.NewAIPlayer()
	game := showdown.NewGame([]showdown.Player{hp, hp1, hp2, hp3})
	game.Start()
	game.DrawCards()
	game.TakeTurn()
	game.JudgeWin()
}
