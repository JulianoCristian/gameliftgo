package game

type TicTacToeGame struct {
	xPlayer       *Player
	oPlayer       *Player
	messageWriter MessageWriter
}

func (g *TicTacToeGame) IsGameEmpty() bool {
	return g.xPlayer == nil && g.oPlayer == nil
}

func (g *TicTacToeGame) HandleCommand(command Command) {
	switch command.CommandType {
	case CommandTypeAddToGame:
		if !g.addPlayer(command.SenderID) {
			// Send message game is full
		}
	case CommandTypeRemoveFromGame:
		if !g.removePlayer(command.SenderID) {
			// Send message game is full
		}
	}
}

func (g *TicTacToeGame) addPlayer(playerID string) bool {
	if g.xPlayer != nil && g.xPlayer.PlayerID == playerID {
		return false
	}
	if g.oPlayer != nil && g.oPlayer.PlayerID == playerID {
		return false
	}
	if g.xPlayer == nil || g.oPlayer == nil {
		player := &Player{
			PlayerID: playerID,
		}
		if g.xPlayer == nil {
			g.xPlayer = player
		} else {
			g.oPlayer = player
		}
		return true
	}
	return false
}

func (g *TicTacToeGame) removePlayer(playerID string) bool {
	removed := false
	inProgress := g.xPlayer != nil || g.oPlayer != nil

	if g.xPlayer.PlayerID == playerID {
		g.xPlayer = nil
		removed = true
	} else if g.oPlayer.PlayerID == playerID {
		g.oPlayer = nil
		removed = true
	}

	if inProgress {
		// Send message that game ended to other player in game
	}

	return removed
}

func NewTicTacToeGame(messageWriter MessageWriter) *TicTacToeGame {
	return &TicTacToeGame{
		xPlayer:       nil,
		oPlayer:       nil,
		messageWriter: messageWriter,
	}
}
