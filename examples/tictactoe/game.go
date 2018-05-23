package main

type Game struct {
	xPlayer *string
	oPlayer *string
}

func (g *Game) AddPlayer(player string) bool {
	if g.xPlayer == nil || g.oPlayer == nil {
		if g.xPlayer == nil {
			g.xPlayer = &player
		} else {
			g.oPlayer = &player
		}
		return true
	}
	return false
}

func (g *Game) RemovePlayer(player string) bool {
	if *g.xPlayer == player {
		g.xPlayer = nil
		return true
	} else if *g.oPlayer == player {
		g.oPlayer = nil
		return true
	}
	return false
}

func (g *Game) IsGameEmpty() bool {
	return g.xPlayer == nil && g.oPlayer == nil
}

func (g *Game) HandleCommand(command Command) {

}

func NewGame() *Game {
	return &Game{
		xPlayer: nil,
		oPlayer: nil,
	}
}
