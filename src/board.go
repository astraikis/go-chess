package main

import (
	"strconv"
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
