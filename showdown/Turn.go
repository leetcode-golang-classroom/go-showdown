package showdown

import "fmt"

type Turn struct {
	result    map[int]*Card
	winnerIdx int
	turn      int
}

func NewTurn(result map[int]*Card, winnerIdx int, turn int) *Turn {
	return &Turn{
		result:    result,
		winnerIdx: winnerIdx,
		turn:      turn,
	}
}

func (t *Turn) ShowTurnResult(g *Game) {
	fmt.Printf("the %v's turn\n", t.turn)
	for pidx, card := range t.result {
		fmt.Printf("Player Name: %s, Card: %v\n", g.players[pidx].GetData().name, card)
	}
	fmt.Printf("turn winner is %v\n", g.players[t.winnerIdx].GetData().name)
}
