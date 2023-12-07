package showdown

import "fmt"

type Game struct {
	exchangeHands []*ExchangeHands
	deck          *Deck
	players       []Player
}

func NewGame(players []Player) *Game {
	return &Game{
		exchangeHands: []*ExchangeHands{},
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
func (g *Game) AddExchanges(exh *ExchangeHands) {
	g.exchangeHands = append(g.exchangeHands, exh)
}
func (g *Game) ResetExchangesResult(exhs []*ExchangeHands) {
	g.exchangeHands = exhs
}
func (g *Game) ExchangeLogic(targetPlayer Player, players []Player, self int) (bool, *ExchangeHands) {
	isExchange := targetPlayer.DecideExchange()
	var exchangeResult *ExchangeHands
	if isExchange {
		i := targetPlayer.ChooseExchange(self)
		exchangeResult = targetPlayer.ExchangeHands(players[i])
		targetPlayer.GetData().SetHasExchangedHand(true)
	}
	return isExchange, exchangeResult
}
func (g *Game) ExchangeCountDown() {
	// collect not finished task
	notFinishedExchanges := []*ExchangeHands{}
	for _, exchangehand := range g.exchangeHands {
		if exchangehand.IsTimeout() {
			continue
		} else {
			exchangehand.CountDown()
			if !exchangehand.IsTimeout() {
				notFinishedExchanges = append(notFinishedExchanges, exchangehand)
			}
		}
	}
	g.ResetExchangesResult(notFinishedExchanges)
}
func (g *Game) RunTurn() (map[int]*Card, int) {
	turnResult := make(map[int]*Card)
	var maxCard *Card
	maxIdx := 0
	for idx, player := range g.players {
		isExchange := false
		isShowCard := true
		hasExchangedHande := player.GetData().GetHasExchangedHand()
		if !hasExchangedHande {
			var exchangeResult *ExchangeHands
			// check if need exchange
			isExchange, exchangeResult = g.ExchangeLogic(player, g.players, idx)
			if isExchange {
				g.AddExchanges(exchangeResult)
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
	return turnResult, maxIdx
}
func (g *Game) GainPoint(maxIdx int) {
	orginPoint := g.players[maxIdx].GetData().GetPoint()
	g.players[maxIdx].GetData().SetPoint(orginPoint + 1)
}
func (g *Game) TakeTurn() {
	// take turn for each player
	for turn := 1; turn <= 13; turn++ {
		// check exchangehands countdown
		g.ExchangeCountDown()
		turnResult, winnerIdx := g.RunTurn()
		// show result
		g.ShowTurnResult(turnResult, winnerIdx, turn)
		g.GainPoint(winnerIdx)
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
	// show all user
	for _, player := range g.players {
		fmt.Printf("user: %s, point: %v\n", player.GetData().name, player.GetData().point)
	}
	winners := []Player{g.players[0]}
	for idx := 1; idx < len(g.players); idx++ {
		if winners[0].GetData().point < g.players[idx].GetData().point {
			winners = []Player{g.players[idx]}
		} else if winners[0].GetData().point == g.players[idx].GetData().point {
			winners = append(winners, g.players[idx])
		}
	}
	if len(winners) == 1 {
		fmt.Printf("winner is %s, point:%v\n", winners[0].GetData().name, winners[0].GetData().point)
	} else {
		fmt.Print("winners are ")
		for idx, winner := range winners {
			fmt.Printf("%s", winner.GetData().name)
			if idx != len(winners)-1 {
				fmt.Print(",")
			} else {
				fmt.Printf(", point:%v\n", winners[0].GetData().point)
			}
		}
	}

}
