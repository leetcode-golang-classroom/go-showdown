package showdown

import (
	"fmt"
	"math/rand"
	"time"
)

type AIPlayer struct {
	data *PlayerData
}

func NewAIPlayer() *AIPlayer {
	return &AIPlayer{
		data: NewPlayerData(),
	}
}

func (aip *AIPlayer) GetData() *PlayerData {
	return aip.data
}

func (hp *AIPlayer) DecideName() string {
	rand.NewSource(time.Now().UnixNano())
	randNumber := rand.Intn(100)
	return fmt.Sprintf("AIPlarer:%v", randNumber)
}

func (hp *AIPlayer) DecideShow() bool {
	rand.NewSource(time.Now().UnixNano())
	showOrNot := rand.Intn(2)
	return showOrNot == 1
}

func (hp *AIPlayer) DecideExchange() bool {
	rand.NewSource(time.Now().UnixNano())
	exchangeOrNot := rand.Intn(2)
	return exchangeOrNot == 1
}

func (aip *AIPlayer) ChooseExchange(self int) int {
	rand.NewSource(time.Now().UnixNano())
	nextRound := (self + 1) % 4
	return nextRound
}

func (aip *AIPlayer) AddHands(card *Card) {
	aip.GetData().AddHands(card)
}

func (aip *AIPlayer) ChooseHandCard() *Card {
	rand.NewSource(time.Now().UnixNano())
	idx := 0
	if len(aip.GetData().hands) > 1 {
		idx = rand.Intn(len(aip.GetData().hands))
		return aip.GetData().ExtractCard(idx)
	}
	temp := aip.GetData().hands[0]
	aip.GetData().hands = []*Card{}
	return temp
}

func (aip *AIPlayer) ExchangeHands(player Player) *ExchangeHands {
	aip.GetData().ExchangeHands(player)
	return NewExchangeHands(aip, player)
}
