package main

import (
	"fmt"
	"hangman-web/dictionary"
	"hangman-web/hangman"
	"html/template"
	"log"
	"net/http"
	"sync"
)

var (
	game *hangman.Game
	mu   sync.Mutex
	port = ":8080"
)

func main() {
	err := dictionary.Load("words.txt")
	if err != nil {
		log.Fatalf("Could not load dictionary: %v", err)
	}

	// Routes du serveur
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hangmangame", hangmanGameHandler)
	http.HandleFunc("/guess", guessHandler)
	http.HandleFunc("/hangmangame/restart", restartGameHandler) // Nouvelle route pour recommencer

	// Serveur de fichiers statiques pour les images du pendu
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home.gohtml")
	data := struct {
		Title           string
		ExplicationHome string
	}{
		Title:           "Vous êtes sur HANGMAN",
		ExplicationHome: "Jeu permettant d'essayer de trouver un mot avec 10 tentatives! Après ces 10 tentatives passées, le joueur est considéré comme pendu et à perdu le jeu ! On lui demandera de rejouer ou non !!!",
	}

	tmpl.Execute(w, data)
}

func hangmanGameHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if game == nil {
		word := dictionary.PickWord()
		if word == "" {
			log.Fatalf("Aucun mot choisi dans le dictionnaire")
		}
		var err error
		game, err = hangman.New(10, word)
		if err != nil {
			log.Printf("Erreur lors de l'initialisation du jeu: %v", err)
			http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			return
		}
	}

	remainingAttempts := game.MaxAttempts - game.Attempts

	// Récupérer le nom de l'image du pendu en fonction des tentatives restantes
	hangmanImage := getHangmanImage(remainingAttempts)

	tmpl, err := template.New("hangmangame").ParseFiles("templates/hangmangame.gohtml")
	if err != nil {
		log.Printf("Erreur de parsing du template: %v", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	data := struct {
		TitleGame         string
		State             string
		Word              string
		DisplayedWord     string // Le mot complet ou avec des "_"
		Attempts          int
		MaxAttempts       int
		Guessed           string
		Tried             string // Ajouter cette ligne
		RemainingAttempts int
		HangmanImage      string // Image à afficher
	}{
		TitleGame:         "Jeu HANGMAN",
		State:             game.State,
		Word:              game.Word,
		DisplayedWord:     game.DisplayWord(), // Mot avec "_" ou mot complet
		Attempts:          game.Attempts,
		MaxAttempts:       game.MaxAttempts,
		Guessed:           game.Guessed,
		Tried:             game.Tried, // Transmettre les lettres essayées
		RemainingAttempts: remainingAttempts,
		HangmanImage:      hangmanImage, // Passer l'URL de l'image
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Erreur lors de l'exécution du template: %v", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if game == nil {
		log.Println("Le jeu n'a pas été initialisé.")
		http.Redirect(w, r, "/hangmangame", http.StatusSeeOther)
		return
	}

	guess := r.FormValue("guess")
	log.Printf("Supposition reçue: %s", guess)

	game.MakeAGuess(guess)

	log.Printf("Etat du jeu après supposition: %s", game.State)
	http.Redirect(w, r, "/hangmangame", http.StatusSeeOther)
}

func getHangmanImage(attempts int) string {
	// Calculer l'index de l'image en fonction des tentatives restantes
	if attempts >= 0 && attempts <= 10 {
		return fmt.Sprintf("/static/images/hangman_%d.jpg", 10-attempts)
	}
	return "/static/images/hangman_10.jpg" // Par défaut, l'image de départ (aucun dessin)
}

func restartGameHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Réinitialiser le jeu en sélectionnant un nouveau mot aléatoire
	word := dictionary.PickWord()
	if word == "" {
		log.Fatalf("Aucun mot choisi dans le dictionnaire")
	}
	var err error
	game, err = hangman.New(10, word)
	if err != nil {
		log.Printf("Erreur lors de l'initialisation du jeu: %v", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	// Rediriger vers la page du jeu avec le nouveau mot
	http.Redirect(w, r, "/hangmangame", http.StatusSeeOther)
}
