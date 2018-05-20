package game

type TicTacToe interface {
}

type ticTacToe struct {
}

func NewTicTacToeGame() TicTacToe {
	return &ticTacToe{}
}
