package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)

var words = []string{
	"passport",
	"driver",
	"children",
	"United States of America",
	"house",
	"zombie",
}

func main() {

	targetWord := getRandomWord()

	guessedLetters := initializeGuessedWords(targetWord)
	hangmanState := 0

	for !(isGameOver(targetWord, guessedLetters, hangmanState)) {

		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()

		if isUsedLetter(guessedLetters, rune(input[0])) {
			fmt.Println("This letter was already used. Try with a different letter.")
			fmt.Println()
		}

		if len(input) > 1 {
			fmt.Println("Invalid input. Please use letters only...")
			continue
		}

		letter := rune(input[0])
		if isCorrectGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}

	printGameState(targetWord, guessedLetters, hangmanState)
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You win!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You lose!")
	} else {
		panic("invalid state. Game is over and there is not winner!")
	}

}

func getRandomWord() string {
	word := words[rand.Intn(len(words))]
	return word
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""
	for _, ch := range targetWord {

		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] {

			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}
		result += " "
	}
	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := os.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}
	return string(data)
}

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	if strings.ContainsRune(targetWord, ' ') {
		guessedLetters[rune(' ')] = true
	}

	return guessedLetters
}

func readInput() string {
	fmt.Print("> ")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.ToLower(strings.TrimSpace(input))
}

func isGameOver(targetWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}

func isUsedLetter(guessedLetters map[rune]bool, letter rune) bool {
	_, isFound := guessedLetters[letter]
	return isFound
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {

	for _, ch := range strings.ToLower(targetWord) {
		if !guessedLetters[ch] {
			return false
		}
	}
	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(strings.ToLower(targetWord), letter)
}
