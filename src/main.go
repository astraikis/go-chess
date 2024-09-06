/*
Jackson Astraikis

Go chess engine and terminal board.

Sep 4, 2024
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// State
var board = [64]byte{}
var isWhite = true
var whitesTurn = true
var gameOver = false

func setBoard(fen string) {
	var blankCount = 0
	var fenI = 0

	for i := 0; i < 64; i++ {
		if blankCount > 0 {
			blankCount--
			continue
		}

		var char = rune(fen[fenI])
		if unicode.IsLetter(char) {
			board[i] = byte(char)
		} else if string(char) == "/" {
			i--
		} else {
			blanks, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			blankCount = blanks - 1
		}

		fenI++
	}
}

func main() {
	fmt.Println("Welcome to go-chess!\n")
	fmt.Println("Enter moves like 'c7 > c5'")
	fmt.Println("Enter position like 'c7' to see all legal moves for that piece\n")

	// Check for FEN in command

	setBoard(StartingFen)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("White or black? ")
	scanner.Scan()
	var colorInput = scanner.Text()
	fmt.Println("")

	if strings.ToLower(colorInput) == "black" || strings.ToLower(colorInput) == "b" {
		isWhite = false
	}

	printBoard(nil, "")

	var move string
	for {
		if gameOver {
			break
		}

		if whitesTurn {
			fmt.Print("\nWhites move: ")
			scanner.Scan()
			move = scanner.Text()

			// Parse move
			moveSplit := strings.Fields(move)
			if len(moveSplit) == 1 {
				printLegalMoves(moveSplit[0])
			}

		} else {
			fmt.Print("\nBlacks move: ")
			scanner.Scan()
			move = scanner.Text()
		}

		whitesTurn = !whitesTurn
	}
}
