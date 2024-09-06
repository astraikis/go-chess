package main

import (
	"errors"
	"fmt"
	"strconv"
)

func getPos(i int, j int) string {
	rank := i + 1
	file := offsetToFile[j]
	return file + strconv.Itoa(rank)
}

// Prints board based on global board
// variable.
// 'moves' set identifies squares to
// highlight.
func printBoard(moves map[string]struct{}, movesPos string) {
	for i := 0; i < 8; i++ {
		fmt.Println("    +---+---+---+---+---+---+---+---+")
		fmt.Printf("%d   ", i+1)

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
					fmt.Printf("| %s ", string(piece))
				}
			}
		}

		fmt.Println("|")
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
