package internal

var SpeedCoinColors = []string{"#3366FF", "#FF1400"}
var SpeedCoinIndex = 0
var SpeedCoinActive = 0

type SpeedCoin struct {
	GameObject

	Index int    `json:"index"`
	Color string `json:"color"`
	Speed Speed  `json:"speed"`
}

func (c *SpeedCoin) Init(speed Speed) {
	c.Clazz = "SpeedCoin"
	c.Index = CoinIndex
	c.Color = SpeedCoinColors[speed]
	c.Speed = speed

	SpeedCoinIndex++
	SpeedCoinActive++
}

func (c *SpeedCoin) destroy() {
	Game.Board.RemoveObject(c.Position)
	CoinActive -= 1
}

func (c *SpeedCoin) HandleCollision(snake *Snake) {
	snake.SetSpeed(c.Speed)
	c.destroy()
}

func CreateRandomSpeedCoin(speed Speed) SpeedCoin {
	var coin SpeedCoin
	coin.Init(speed)
	return coin
}
