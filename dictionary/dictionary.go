package dictionary

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var words []string

func Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open dictionary file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("could not read dictionary file: %v", err)
	}
	return nil
}

func PickWord() string {
	randomNumber := rand.Intn(100)
	fmt.Println(randomNumber)
	return words[rand.Intn(len(words))]
}
