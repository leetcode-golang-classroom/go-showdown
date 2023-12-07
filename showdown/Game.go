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
func (g *Game) HandleExchange(player Player, playerIdx int) bool {
	isExchange := false
	isShowCard := true
	hasExchangedHand := player.GetData().GetHasExchangedHand()
	if !hasExchangedHand {
		var exchangeResult *ExchangeHands
		// check if need exchange
		isExchange, exchangeResult = g.ExchangeLogic(player, g.players, playerIdx)
		if isExchange {
			g.AddExchanges(exchangeResult)
			isShowCard = player.DecideShow()
		}
	}
	return isShowCard
}
func (g *Game) HandleShowCardAndCalculateMaxLogic(player Player, playerIdx int, turnRecords map[int]*Card, maxCard *Card, isShowCard bool, winnerIdx int) (int, *Card) {
	if isShowCard && len(player.GetData().hands) > 0 {
		currentCard := player.ChooseHandCard()
		turnRecords[playerIdx] = currentCard
		if maxCard == nil {
			maxCard = currentCard
		}
		if maxCard != currentCard && maxCard.Compare(currentCard) {
			maxCard = currentCard
			winnerIdx = playerIdx
		}
	}
	return winnerIdx, maxCard
}
func (g *Game) RunTurn(turn int) *Turn {
	turnRecords := make(map[int]*Card)
	var maxCard *Card
	winnerIdx := 0
	for idx, player := range g.players {
		// handle exchange logic
		isShowCard := g.HandleExchange(player, idx)
		winnerIdx, maxCard = g.HandleShowCardAndCalculateMaxLogic(player, idx, turnRecords, maxCard, isShowCard, winnerIdx)
	}
	return NewTurn(turnRecords, winnerIdx, turn)
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
		turnRecords := g.RunTurn(turn)
		// show result
		turnRecords.ShowTurnResult(g)
		g.GainPoint(turnRecords.winnerIdx)
	}
}
func (g *Game) ShowAllPlayers() {
	for _, player := range g.players {
		fmt.Printf("user: %s, point: %v\n", player.GetData().name, player.GetData().point)
	}
}
func (g *Game) FindWinners() []Player {
	winners := []Player{g.players[0]}
	for idx := 1; idx < len(g.players); idx++ {
		if winners[0].GetData().point < g.players[idx].GetData().point {
			winners = []Player{g.players[idx]}
		} else if winners[0].GetData().point == g.players[idx].GetData().point {
			winners = append(winners, g.players[idx])
		}
	}
	return winners
}
func (g *Game) ShowWinners(winners []Player) {
	if len(winners) == 1 {
		fmt.Printf("winner is %s, point:%v\n", winners[0].GetData().name, winners[0].GetData().point)
		return
	}
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
func (g *Game) JudgeWin() {
	// find out the winner and display
	// show all user
	g.ShowAllPlayers()
	winners := g.FindWinners()
	g.ShowWinners(winners)

}
