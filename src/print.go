package main

import (
	"errors"
	"fmt"
)

// Prints board based on global board
// variable.
// 'moves' set identifies squares to
// highlight.
func printBoard(moves map[string]struct{}, movesPos string) {
	for i := 0; i < 8; i++ {
		fmt.Print("    +---+---+---+---+---+---+---+---+")
		if i == 0 {
			fmt.Print("    FEN:")
		} else if i == 2 && blackScore > 0 {
			fmt.Print("    Black:")
		} else if i == 4 && whiteScore > 0 {
			fmt.Print("    White:")
		}
		fmt.Printf("\n%d   ", i+1)

		for j := 0; j < 8; j++ {
			var piece = board[(i*8)+j]
			pos := getPos(i, j)
			_, highlighted := moves[pos]
			if piece == 0 {
				if highlighted {
					fmt.Print("|///")
				} else {
					fmt.Print("|   ")
				}
			} else {
				if movesPos == pos {
					fmt.Printf("|*%s*", string(piece))
				} else {
					if highlighted {
						fmt.Printf("|/%s/", string(piece))
					} else {
						fmt.Printf("| %s ", string(piece))
					}
				}
			}
		}

		fmt.Print("|")
		if i == 0 {
			fmt.Print("    " + currFen)
		} else if i == 2 && blackScore > 0 {
			fmt.Print("    ")
			for piece, count := range capturedPiecesWhite {
				for i := 0; i < count; i++ {
					fmt.Print(piece)
				}
			}
			fmt.Printf(" (%d)", blackScore)
		} else if i == 4 && whiteScore > 0 {
			fmt.Print("    ")
			for piece, count := range capturedPiecesBlack {
				for i := 0; i < count; i++ {
					fmt.Print(piece)
				}
			}
			fmt.Printf(" (%d)", whiteScore)
		}
		fmt.Println("")
	}

	fmt.Println("    +---+---+---+---+---+---+---+---+")
	fmt.Println("\n      a   b   c   d   e   f   g   h")
}

// Prints board with legal moves highlighted.
func printLegalMoves(pos string) error {
	if len(pos) != 2 {
		return errors.New("invalid position")
	}

	var moves = getLegalMoves(pos)
	printBoard(moves, pos)

	return nil
}
