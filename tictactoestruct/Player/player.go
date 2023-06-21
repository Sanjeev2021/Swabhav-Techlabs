package player

type Player struct {
	Name   string
	Symbol string
}

func NewPlayer(playerName, Symbol string) *Player {
	newPlayer := &Player{
		Name:   playerName,
		Symbol: Symbol,
	}
	return newPlayer
}
