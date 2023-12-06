package showdown

import (
	"fmt"
	"strings"
)

type HumanPlayer struct {
	data *PlayerData
}

func NewHumanPlayer() *HumanPlayer {
	return &HumanPlayer{
		data: NewPlayerData(),
	}
}

func (hp *HumanPlayer) GetData() *PlayerData {
	return hp.data
}

func (hp *HumanPlayer) ExchangeHands(player Player) *ExchangeHands {
	hp.GetData().ExchangeHands(player)
	return NewExchangeHands(hp, player)
}

func (hp *HumanPlayer) ChooseHandCard() *Card {
	// console handle
	fmt.Println(len(hp.GetData().hands), hp.GetData().hands)
	// fmt.Println()
	var idx int
	answerIsInRange := false
	for !answerIsInRange {
		fmt.Printf("Choose Card to Show:")
		fmt.Scan(&idx)
		answerIsInRange = idx >= 0 && idx < len(hp.GetData().hands)
	}
	return hp.GetData().ExtractCard(idx)
}

func (hp *HumanPlayer) DecideName() string {
	fmt.Printf("Please enter player name:")
	var name string
	fmt.Scan(&name)
	return name
}

func (hp *HumanPlayer) DecideShow() bool {
	fmt.Printf("Show the card or Not:(Y/N)?")
	var isShowCard string
	fmt.Scan(&isShowCard)

	return strings.Compare(isShowCard, "Y") == 0
}

func (hp *HumanPlayer) DecideExchange() bool {
	fmt.Printf("Exchange hands or not:(Y/N)?")
	var isExchange string
	fmt.Scan(&isExchange)
	return strings.Compare(isExchange, "Y") == 0
}

func (hp *HumanPlayer) ChooseExchange(self int) int {
	var targetPlayerNumber int
	answerIsInRange := false
	for !answerIsInRange {
		fmt.Printf("Choose hands or not:(0-3) other than %v?", self)
		fmt.Scan(&targetPlayerNumber)
		answerIsInRange = targetPlayerNumber >= 0 && targetPlayerNumber < 4 && self != targetPlayerNumber
	}
	return targetPlayerNumber
}

func (hp *HumanPlayer) AddHands(card *Card) {
	hp.GetData().AddHands(card)
}
