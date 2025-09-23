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
