package showdown

import "fmt"

type Player interface {
	ChooseHandCard() *Card
	ExchangeHands(player Player) *ExchangeHands
	GetData() *PlayerData
	DecideName() string
	DecideShow() bool
	DecideExchange() bool
	ChooseExchange(self int) int
	AddHands(c *Card)
}

type PlayerData struct {
	point            int
	hasExchangedHand bool
	hands            []*Card
	name             string
}

func (pd *PlayerData) AddHands(card *Card) {
	pd.hands = append(pd.hands, card)
}

func (pd *PlayerData) SetPlayerName(name string) {
	pd.name = name
}

func (pd *PlayerData) SetPoint(point int) {
	pd.point = point
}
func (pd *PlayerData) GetPoint() int {
	return pd.point
}

func (pd *PlayerData) GetHasExchangedHand() bool {
	return pd.hasExchangedHand
}

func (pd *PlayerData) SetHasExchangedHand(hasExchangedHand bool) {
	pd.hasExchangedHand = hasExchangedHand
}

func (pd *PlayerData) SetHands(hands []*Card) {
	pd.hands = hands
}
func (pd *PlayerData) ExchangeHands(player Player) {
	temp := player.GetData().hands
	player.GetData().SetHands(pd.hands)
	pd.hands = temp
}

func (pd *PlayerData) ExtractCard(idx int) *Card {
	temp := pd.hands[idx]
	if idx < len(pd.hands)-1 {
		pd.hands = append(pd.hands[:idx], pd.hands[idx+1:]...)
	} else {
		pd.hands = append([]*Card{}, pd.hands[:idx]...)
	}
	return temp
}
func NewPlayerData() *PlayerData {
	return &PlayerData{
		point:            0,
		hasExchangedHand: false,
		hands:            []*Card{},
		name:             "",
	}
}
func (pd *PlayerData) String() string {
	return fmt.Sprintf("name: %s, point: %v, hasExchangedHand: %v,\n%v", pd.name, pd.point, pd.hasExchangedHand, pd.hands)
}
