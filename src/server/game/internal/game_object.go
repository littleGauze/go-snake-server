package internal

type Drawable interface {
	GetPosition() Position
	SetPosition(pos Position)
	HandleCollision(object *Snake)
}

type GameObject struct {
	Clazz    string   `json:"clazz"`
	Position Position `json:"position"`
}

func (g *GameObject) SetPosition(pos Position) {
	g.Position = pos
}

func (g *GameObject) GetPosition() Position {
	return g.Position
}

type PlayObject struct {
	GameObject
	Name string `json:"name"`
}
