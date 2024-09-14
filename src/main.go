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
	"strings"
)

func main() {
	fmt.Println("Welcome to go-chess!\n")
	fmt.Println("Enter moves like 'c7 > c5'")
	fmt.Println("Enter position like 'c7' to see all legal moves for that piece")
	fmt.Println("Enter position like 'c7?' to see the best move for that piece\n")

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
			fmt.Println("")

			// Parse move
			moveSplit := strings.Fields(move)
			if len(moveSplit) == 1 {
				printLegalMoves(moveSplit[0])
				whitesTurn = !whitesTurn
			}

		} else {
			fmt.Print("\nBlacks move: ")
			scanner.Scan()
			move = scanner.Text()
		}

		whitesTurn = !whitesTurn
	}
}
