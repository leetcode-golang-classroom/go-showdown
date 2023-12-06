package showdown

import "fmt"

type Game struct {
	exchangeHands []ExchangeHands
	deck          *Deck
	players       []Player
}

func NewGame(players []Player) *Game {
	return &Game{
		exchangeHands: []ExchangeHands{},
		deck:          NewDeck(),
		players:       players,
	}
}
func (g *Game) Start() {
	// start game
	// setup players name
	for _, player := range g.players {
		player.GetData().SetPlayerName(player.DecideName())
	}
	g.deck.Shuffle()
}

func (g *Game) DrawCards() {
	// draw cards
	for idx := 0; idx < 13; idx++ {
		for _, player := range g.players {
			if g.deck.HasCards() {
				g.deck.DrawCard(player)
			}
		}
	}
}
func (g *Game) ExchangeLogic(targetPlayer Player, players []Player, self int) bool {
	isExchange := targetPlayer.DecideExchange()
	if isExchange {
		i := targetPlayer.ChooseExchange(self)
		targetPlayer.ExchangeHands(players[i])
		targetPlayer.GetData().SetHasExchangedHand(true)
	}
	return isExchange
}
func (g *Game) TakeTurn() {
	// take turn for each player
	for turn := 1; turn <= 13; turn++ {
		turnResult := make(map[int]*Card)
		var maxCard *Card
		maxIdx := 0
		// check exchangehands countdown
		for _, exchangehand := range g.exchangeHands {
			exchangehand.CountDown()
		}
		// check if need exchange
		for idx, player := range g.players {
			isExchange := false
			isShowCard := true
			hasExchangedHande := player.GetData().GetHasExchangedHand()
			if !hasExchangedHande {
				isExchange = g.ExchangeLogic(player, g.players, idx)
				if isExchange {
					isShowCard = player.DecideShow()
				}
			}
			if isShowCard && len(player.GetData().hands) > 0 {
				currentCard := player.ChooseHandCard()
				turnResult[idx] = currentCard
				if maxCard == nil {
					maxCard = currentCard
				}
				if maxCard != currentCard && maxCard.Compare(currentCard) {
					maxCard = currentCard
					maxIdx = idx
				}
			}
		}
		// show result
		g.ShowTurnResult(turnResult, maxIdx, turn)
		orginPoint := g.players[maxIdx].GetData().GetPoint()
		g.players[maxIdx].GetData().SetPoint(orginPoint + 1)
	}
}
func (g *Game) ShowTurnResult(turnResult map[int]*Card, maxIdx int, turn int) {
	fmt.Printf("the %v's turn\n", turn)
	for pidx, card := range turnResult {
		fmt.Printf("Player Name: %s, Card: %v\n", g.players[pidx].GetData().name, card)
	}
	fmt.Printf("turn winner is %v\n", g.players[maxIdx].GetData().name)
}
func (g *Game) JudgeWin() {
	// find out the winner and display
	winner := g.players[0]
	for idx := 1; idx < len(g.players); idx++ {
		if winner.GetData().GetPoint() < g.players[idx].GetData().GetPoint() {
			winner = g.players[idx]
		}
	}
	fmt.Printf("winner is %s, point:%v\n", winner.GetData().name, winner.GetData().point)
}
