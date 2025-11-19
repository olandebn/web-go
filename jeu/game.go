package main

const (
	Rows = 6
	Cols = 7
)

type Game struct {
	Board         [][]int
	CurrentPlayer int
	TurnCount     int
	Winner        int
	WinningCells  map[[2]int]bool
}

// Nouvelle partie vide
func NewGame() *Game {
	board := make([][]int, Rows)
	for i := range board {
		board[i] = make([]int, Cols)
	}
	return &Game{
		Board:         board,
		CurrentPlayer: 1,
		WinningCells:  make(map[[2]int]bool),
	}
}

// Jouer un coup
func (g *Game) Play(col int) bool {
	if col < 0 || col >= Cols || g.Winner != 0 {
		return false
	}
	for row := Rows - 1; row >= 0; row-- {
		if g.Board[row][col] == 0 {
			g.Board[row][col] = g.CurrentPlayer
			g.TurnCount++
			if g.checkWin(row, col, g.CurrentPlayer) {
				g.Winner = g.CurrentPlayer
			} else {
				g.switchPlayer()
			}
			return true
		}
	}
	return false
}

func (g *Game) switchPlayer() {
	if g.CurrentPlayer == 1 {
		g.CurrentPlayer = 2
	} else {
		g.CurrentPlayer = 1
	}
}

// Vérification victoire
func (g *Game) checkWin(row, col, player int) bool {
	dirs := [][2]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}}
	for _, d := range dirs {
		count := 1
		for i := 1; i < 4; i++ {
			r, c := row+d[0]*i, col+d[1]*i
			if r < 0 || r >= Rows || c < 0 || c >= Cols || g.Board[r][c] != player {
				break
			}
			count++
		}
		for i := 1; i < 4; i++ {
			r, c := row-d[0]*i, col-d[1]*i
			if r < 0 || r >= Rows || c < 0 || c >= Cols || g.Board[r][c] != player {
				break
			}
			count++
		}
		if count >= 4 {
			for i := -3; i <= 3; i++ {
				r, c := row+d[0]*i, col+d[1]*i
				if r >= 0 && r < Rows && c >= 0 && c < Cols && g.Board[r][c] == player {
					g.WinningCells[[2]int{r, c}] = true
				}
			}
			return true
		}
	}
	return false
}

// Réinitialiser
func (g *Game) Reset() {
	for i := range g.Board {
		for j := range g.Board[i] {
			g.Board[i][j] = 0
		}
	}
	g.CurrentPlayer = 1
	g.TurnCount = 0
	g.Winner = 0
	g.WinningCells = make(map[[2]int]bool)
}
