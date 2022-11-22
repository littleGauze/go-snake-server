package internal

var CoinValues = []int{200, 600, 800, 1000, 2000}
var CoinIndex = 0
var CoinActive = 0

type Coin struct {
	GameObject

	index int
	value int
}

func (c *Coin) Init(value int) {
	c.Clazz = "Coin"
	c.value = value
	c.index = CoinIndex

	CoinIndex++
	CoinActive++
}

func (c *Coin) destroy() {
	Game.Board.RemoveObject(c.Position)
	CoinActive -= 1
}

func (c *Coin) HandleCollision(snake *Snake) {
	snake.Points += c.value
	snake.MaxLength += 1
	c.destroy()
}

func CreateRandomCoin() Coin {
	var coin Coin
	coin.Init(1)
	return coin
}
