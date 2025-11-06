package pisicne

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sync"
)

var (
	mu      sync.Mutex
	game    *Game
	tmpl    *template.Template
	history []string
)

func main() {
	game = NewGame()

	// Charger le template
	tmpl = template.Must(template.New("index.html").Funcs(template.FuncMap{
		"isWinningCell": func(r, c int) bool {
			return game.WinningCells[[2]int{r, c}]
		},
	}).ParseFiles("templates/index.html"))

	// Charger l'historique depuis fichier
	loadHistory()

	// Servir les images statiques
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	http.HandleFunc("/", handleMenu)
	http.HandleFunc("/start", handleStart)
	http.HandleFunc("/rules", handleRules)
	http.HandleFunc("/info", handleInfo)
	http.HandleFunc("/history", handleHistory)
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/reset", handleReset)

	fmt.Println("Serveur en cours sur http://localhost:8080 ...")
	http.ListenAndServe(":8080", nil)
}

func handleMenu(w http.ResponseWriter, r *http.Request) {
	render(w, "menu")
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	game.Reset()
	mu.Unlock()
	render(w, "jeu")
}

func handleRules(w http.ResponseWriter, r *http.Request) {
	render(w, "regles")
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	render(w, "info")
}

func handleHistory(w http.ResponseWriter, r *http.Request) {
	render(w, "historique")
}

func handlePlay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	col := r.FormValue("col")
	var c int
	fmt.Sscanf(col, "%d", &c)

	mu.Lock()
	if game.Play(c) && game.Winner != 0 {
		// Ajouter dans l'historique
		entry := fmt.Sprintf("Joueur %d a gagné après %d tours", game.Winner, game.TurnCount)
		history = append(history, entry)
		saveHistory()
	}
	mu.Unlock()

	if game.Winner != 0 {
		render(w, "victory")
		return
	}

	render(w, "jeu")
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	game.Reset()
	mu.Unlock()
	render(w, "menu")
}

func render(w http.ResponseWriter, view string) {
	mu.Lock()
	defer mu.Unlock()
	data := map[string]interface{}{
		"View":          view,
		"Board":         game.Board,
		"Winner":        game.Winner,
		"CurrentPlayer": game.CurrentPlayer,
		"History":       history,
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

// --- Historique ---
func loadHistory() {
	f, err := os.Open("history.txt")
	if err != nil {
		history = []string{}
		return
	}
	defer f.Close()

	var lines []string
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			lines = append(lines, string(buf[:n]))
		}
		if err != nil {
			break
		}
	}

	history = []string{}
	for _, line := range lines {
		for _, l := range splitLines(line) {
			if l != "" {
				history = append(history, l)
			}
		}
	}
}

func saveHistory() {
	f, err := os.Create("history.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'écriture de l'historique :", err)
		return
	}
	defer f.Close()
	for _, h := range history {
		fmt.Fprintln(f, h)
	}
}

// Helper pour découper par lignes
func splitLines(s string) []string {
	lines := []string{}
	start := 0
	for i, c := range s {
		if c == '\n' || c == '\r' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}
