package hangman

import (
	"fmt"
	"strings"

	"math/rand"
)

// Structure du jeu Hangman
type Game struct {
	MaxAttempts     int           // Nombre maximal d'essais
	Attempts        int           // Nombre d'essais effectués
	Word            string        // Mot à deviner
	Guessed         string        // Lettres déjà devinées
	Tried           string        // Toutes les lettres essayées (correctes ou non)
	State           string        // Etat du jeu ("playing", "won", "lost")
	RevealedLetters map[rune]bool // Utilisé pour savoir quelles lettres ont été révélées
}

// Crée un nouveau jeu Hangman
func New(maxAttempts int, word string) (*Game, error) {
	if word == "" {
		return nil, fmt.Errorf("mot vide, impossible de créer le jeu")
	}

	// Calculer le nombre de lettres à révéler
	numRevealed := (len(word) / 2) - 1
	if numRevealed < 0 {
		numRevealed = 0
	}

	// Créer un ensemble pour savoir quelles lettres ont été révélées
	revealedLetters := make(map[rune]bool)
	guessed := "" // Commencez avec un guessed vide

	// Compter les occurrences des lettres dans le mot
	letterCount := make(map[rune]int)
	for _, letter := range word {
		letterCount[letter]++
	}

	// Sélectionner aléatoirement les lettres à révéler sans répéter les lettres
	for len(revealedLetters) < numRevealed {
		// Sélectionner une lettre aléatoire dans le mot
		index := rand.Intn(len(word))
		letter := rune(word[index])

		// Ajouter la lettre à l'ensemble des lettres révélées, si elle n'y est pas déjà
		if !revealedLetters[letter] {
			revealedLetters[letter] = true
			// Ajouter la lettre révélée à la liste guessed
			guessed += string(letter)
		}
	}

	return &Game{
		MaxAttempts:     maxAttempts,
		Attempts:        0,
		Word:            word,
		Guessed:         guessed,
		State:           "playing", // Le jeu commence dans l'état "playing"
		RevealedLetters: revealedLetters,
	}, nil
}

// Méthode pour afficher le mot avec les lettres devinées et celles révélées
func (g *Game) DisplayWord() string {
	var displayWord string
	letterCount := make(map[rune]int)

	// Compter les occurrences de chaque lettre dans le mot
	for _, letter := range g.Word {
		letterCount[letter]++
	}

	// Afficher chaque lettre du mot
	for _, letter := range g.Word {
		// Vérifier si la lettre est déjà devinée ou révélée
		if strings.Contains(g.Guessed, string(letter)) || g.RevealedLetters[letter] {
			displayWord += string(letter)
		} else {
			displayWord += "_"
		}
	}
	// Finalement, afficher le mot avec les lettres devinées et révélées.
	return displayWord
}

// Méthode pour traiter une supposition du joueur
func (g *Game) MakeAGuess(guess string) {
	// Si le jeu est déjà terminé, on ne traite pas de nouvelles suppositions
	if g.State != "playing" {
		return
	}

	// Convertir la supposition en minuscule
	guess = strings.ToLower(guess)

	// Vérification si la supposition est valide (une seule lettre)
	if len(guess) != 1 || !IsLetter(rune(guess[0])) {
		return // Ignorer les suppositions invalides
	}

	// Ajouter la lettre à la liste de toutes les lettres essayées
	g.Tried += " " + guess

	// Si la lettre est dans le mot, on l'ajoute à la liste des lettres devinées
	if strings.Contains(g.Word, guess) {
		g.Guessed += " " + guess
		// Ajouter la lettre révélée à RevealedLetters
		g.RevealedLetters[rune(guess[0])] = true
	} else {
		// Si la lettre n'est pas dans le mot, on incrémente les tentatives
		g.Attempts++
	}

	// Vérifier si le jeu est terminé (gagné ou perdu)
	if g.Attempts >= g.MaxAttempts {
		g.State = "lost" // Le joueur a perdu
	} else if g.isWordGuessed() {
		g.State = "won" // Le joueur a gagné
	}
}

// Vérifie si le joueur a deviné tout le mot
func (g *Game) isWordGuessed() bool {
	for _, letter := range g.Word {
		if !strings.Contains(g.Guessed, string(letter)) {
			return false
		}
	}
	return true
}

func IsLetter(s rune) bool {
	// Vérifier si la rune est une lettre (y compris les lettres accentuées)
	if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') ||
		(s >= 'à' && s <= 'ÿ') || // Plage des lettres accentuées communes
		s == 'ç' || s == 'Ç' || // Ajouter les lettres spéciales comme 'ç' et 'Ç'
		s == 'é' || s == 'è' || s == 'ê' || s == 'ë' ||
		s == 'à' || s == 'â' || s == 'ä' || s == 'ï' ||
		s == 'î' || s == 'ô' || s == 'ö' || s == 'ù' ||
		s == 'û' || s == 'ü' || s == 'ÿ' {
		return true
	}
	return false
}
