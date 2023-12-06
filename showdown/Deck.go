package showdown

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	cards  []*Card
	maxLen int
}

func InitCard() []*Card {
	cards := []*Card{}
	for _, suit := range Suits {
		for _, rank := range Ranks {
			cards = append(cards, &Card{rank: rank, suit: suit})
		}
	}
	return cards
}
func (d *Deck) ShowDeck() {
	for _, card := range d.cards {
		fmt.Println(card)
	}
}
func NewDeck() *Deck {
	return &Deck{
		cards:  InitCard(),
		maxLen: 52,
	}
}
func (s *Deck) HasCards() bool {
	return len(s.cards) > 0
}
func (s *Deck) Shuffle() {
	// shuffle current totalCards
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(s.cards), func(i, j int) {
		s.cards[i], s.cards[j] = s.cards[j], s.cards[i]
	})
}

func (s *Deck) DrawCard(player Player) {
	rand.NewSource(time.Now().UnixNano())
	idx := rand.Intn(len(s.cards))
	temp := s.cards[idx]
	if idx >= 1 {
		s.cards = append(s.cards[:idx-1], s.cards[idx:]...)
	} else {
		s.cards = append([]*Card{}, s.cards[1:]...)
	}

	player.AddHands(temp)
}
