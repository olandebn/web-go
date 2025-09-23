package main

const (
	Rows    = 6
	Columns = 7
)

type Player int

const (
	Empty Player = iota
	Player1
	Player2
)

type Game struct {
	Grid      [Rows][Columns]Player
	Current   Player
	GameOver  bool
	Winner    Player
	MoveCount int
}

func NewGame() *Game {
	return &Game{
		Current:   Player1,
		GameOver:  false,
		Winner:    Empty,
		MoveCount: 0,
	}
}

func (g *Game) PlayMove(col int) bool {
	if g.GameOver || col < 0 || col >= Columns {
		return false
	}
	for row := Rows - 1; row >= 0; row-- {
		if g.Grid[row][col] == Empty {
		}
	}
}
