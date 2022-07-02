package game

import (
	"errors"
	"fmt"
	"pranav/board"
	"pranav/player"
)

type Game struct {
	Players   [2]player.PLayer
	MainBoard board.Board
	Turn      int
}

func CreateNewGame(p1name, p2name string, n int) *Game {
	var game Game
	game.Players[0] = *player.CreateNewPlayer(p1name, 1)
	game.Players[1] = *player.CreateNewPlayer(p2name, 2)
	game.MainBoard = *board.CreateNewBoard(n)
	game.Turn = 1
	return &game

}

func (g *Game) Play(position int, player player.PLayer) error {
	if (g.MainBoard.Cells[position].Mark != "X" && g.MainBoard.Cells[position].Mark != "O") && position < g.MainBoard.Sixe*g.MainBoard.Sixe && position >= 0 {
		if g.Turn%2 == 0 {
			if player == g.Players[1] {
				g.MainBoard.Cells[position].Mark = g.Players[1].Mark
				g.Turn++
				return nil
			} else {
				fmt.Println("It's Player 2 turn")
			}
		} else {
			if player == g.Players[0] {
				g.MainBoard.Cells[position].Mark = g.Players[0].Mark
				g.Turn++
				return nil
			} else {
				fmt.Println("It's Player 1 turn")
			}

		}
	}
	return errors.New("this position is not valid ")

}
