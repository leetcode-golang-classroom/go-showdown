package showdown

type ExchangeHands struct {
	exchanger           Player
	exchangee           Player
	changeBackCountDown int
}

func NewExchangeHands(exchanger, exchangee Player) *ExchangeHands {
	return &ExchangeHands{
		exchanger:           exchanger,
		exchangee:           exchangee,
		changeBackCountDown: 3,
	}
}

func (exh *ExchangeHands) exchangeBack() {
	exh.exchanger.ExchangeHands(exh.exchangee)
}

func (exh *ExchangeHands) CountDown() {
	if exh.changeBackCountDown > 0 {
		exh.changeBackCountDown--
	}
	if exh.changeBackCountDown == 0 {
		exh.exchangeBack()
	}
}

func (exh *ExchangeHands) IsTimeout() bool {
	return exh.changeBackCountDown == 0
}
