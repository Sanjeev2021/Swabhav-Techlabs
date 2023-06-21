package game

import (
	"errors"
	"fmt"

	board "tictactoestruct/Board"
	player "tictactoestruct/Player"
)

type Game struct {
	//Player []*player.Player
	Player []*player.Player
	board  *board.Board
	turn   int
}

func NewGame(p1Name, p2Name string) *Game {
	player1 := player.NewPlayer(p1Name, "X")
	player2 := player.NewPlayer(p2Name, "O")
	b := board.NewBoard()

	return &Game{
		turn:   0,
		Player: []*player.Player{player1, player2},
		board:  b,
	}
}

func (g *Game) Play(cellNum int) (string, error) {
	if g.turn < 0 {
		return "", errors.New("Game has ended")
	}
	if g.board.Cells[cellNum].IsCellMarked() {
		return "", errors.New("cell is already marked")
	}
	currentPlayer := g.Player[g.turn%2]
	g.board.Cells[cellNum].MarkCell(currentPlayer.Symbol)
	g.turn++
	g.board.PrintBoard()
	if g.board.ResultAnalyzer() {
		g.turn = -1
		fmt.Println("The player to win the match is :")
		return currentPlayer.Name, nil
	}
	return "", nil

}
