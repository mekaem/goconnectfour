package main

import (
	"errors"
	"fmt"
)

const (
	rows    = 6
	columns = 7
)

type Player int

const (
	Empty Player = iota
	Player1
	Player2
)

type Game struct {
	board  [rows][columns]Player
	player Player
	turns  int
}

func NewGame() *Game {
	return &Game{player: Player1}
}

func (g *Game) PrintBoard() {
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			fmt.Printf("%d ", g.board[r][c])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Game) DropToken(column int) error {
	if column < 0 || column >= columns {
		return errors.New("invalid column")
	}

	for r := rows - 1; r >= 0; r-- {
		if g.board[r][column] == Empty {
			g.board[r][column] = g.player
			g.turns++
			if g.checkWin(r, column) {
				fmt.Printf("Player %d wins!\n", g.player)
				return nil
			}
			g.switchPlayer()
			return nil
		}
	}
	return errors.New("column is full")
}

func (g *Game) switchPlayer() {
	if g.player == Player1 {
		g.player = Player2
	} else {
		g.player = Player1
	}
}

func (g *Game) checkWin(row, column int) bool {
	directions := [][2]int{
		{0, 1},  // Horizontal
		{1, 0},  // Vertical
		{1, 1},  // Diagonal down-right
		{1, -1}, // Diagonal down-left
	}

	for _, dir := range directions {
		count := 1
		for _, sign := range []int{1, -1} {
			for i := 1; i < 4; i++ {
				r, c := row+i*dir[0]*sign, column+i*dir[1]*sign
				if r < 0 || r >= rows || c < 0 || c >= columns || g.board[r][c] != g.player {
					break
				}
				count++
			}
		}
		if count >= 4 {
			return true
		}
	}
	return false
}

func (g *Game) IsDraw() bool {
	return g.turns == rows*columns
}

func main() {
	game := NewGame()
	game.PrintBoard()

	columns := []int{3, 3, 4, 4, 5, 5, 6}
	for _, col := range columns {
		err := game.DropToken(col)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		game.PrintBoard()
		if game.IsDraw() {
			fmt.Println("The game is a draw!")
			return
		}
	}
}
