package game

import (
	"log"
)

type TicTacToeGame struct {
	xPlayer       *Player
	oPlayer       *Player
	messageWriter MessageWriter
	inProgress    bool
	commands      chan Command
	stop          chan bool
}

func (g *TicTacToeGame) Run() {
	go func() {
		for {
			stop := false

			select {
			case command := <-g.commands:
				g.handleCommand(command)
			case <-g.stop:
				stop = true
				break
			}

			if stop {
				break
			}
		}
	}()
}

func (g *TicTacToeGame) Stop() {
	g.stop <- true
}

func (g *TicTacToeGame) HandleCommand(command Command) {
	g.commands <- command
}

func (g *TicTacToeGame) RemovePlayer(playerID string) bool {
	if !g.inProgress {
		return false
	}

	removed := false

	if g.xPlayer != nil && g.xPlayer.PlayerID == playerID {
		g.xPlayer = nil
		removed = true
	} else if g.oPlayer != nil && g.oPlayer.PlayerID == playerID {
		g.oPlayer = nil
		removed = true
	}

	return removed
}

func (g *TicTacToeGame) handleCommand(command Command) {
	log.Printf("handleCommand %s", command.CommandType)
	switch command.CommandType {
	case CommandTypeAddToGame:
		if g.addPlayer(command.SenderID) {
			g.messageWriter.WriteAll(newMessage(MessageTypePlayerAddedToGame, JSON{"playerID": command.SenderID}))

			if g.xPlayer != nil && g.oPlayer != nil {
				g.inProgress = true
				g.messageWriter.WriteAll(newEmptyMessage(MessageTypeGameStarted))
			}
			// Send game state to all
		} else {
			// Game full or already added
		}
	case CommandTypeRemoveFromGame:
		if g.RemovePlayer(command.SenderID) {
			if g.inProgress {
				g.xPlayer = nil
				g.oPlayer = nil
				g.inProgress = false
				g.messageWriter.WriteAll(newEmptyMessage(MessageTypeGameEnded))
			}
			g.messageWriter.WriteAll(newMessage(MessageTypePlayerRemovedFromGame, JSON{"playerID": command.SenderID}))
			// Send game state to all
		} else {
			// User not in game or game not in progress
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

func NewTicTacToeGame(messageWriter MessageWriter) *TicTacToeGame {
	return &TicTacToeGame{
		xPlayer:       nil,
		oPlayer:       nil,
		messageWriter: messageWriter,
		inProgress:    false,
		commands:      make(chan Command),
		stop:          make(chan bool),
	}
}
