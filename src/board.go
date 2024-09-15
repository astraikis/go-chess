package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// State
var board = [64]byte{}
var isWhite = true
var whitesTurn = true
var gameOver = false
var currFen = ""
var capturedPiecesWhite = map[string]int{}
var capturedPiecesBlack = map[string]int{}
var blackScore = 0
var whiteScore = 0

func setCapturedPieces() {
	capturedPiecesWhite = map[string]int{
		"P": 8,
		"R": 2,
		"N": 2,
		"B": 2,
		"Q": 1,
		"K": 1,
	}
	capturedPiecesBlack = map[string]int{
		"p": 8,
		"r": 2,
		"n": 2,
		"b": 2,
		"q": 1,
		"k": 1,
	}
	for i := 0; i < 64; i++ {
		switch string(board[i]) {
		case "p":
			capturedPiecesBlack["p"]--
		case "r":
			capturedPiecesBlack["r"]--
		case "n":
			capturedPiecesBlack["n"]--
		case "b":
			capturedPiecesBlack["b"]--
		case "q":
			capturedPiecesBlack["q"]--
		case "k":
			capturedPiecesBlack["k"]--
		case "P":
			capturedPiecesWhite["P"]--
		case "R":
			capturedPiecesWhite["R"]--
		case "N":
			capturedPiecesWhite["N"]--
		case "B":
			capturedPiecesWhite["B"]--
		case "Q":
			capturedPiecesWhite["Q"]--
		case "K":
			capturedPiecesWhite["K"]--
		default:
			continue
		}
	}
}

func setPieceScores() {
	for piece, count := range capturedPiecesWhite {
		blackScore += count * pieceToValue[piece]
	}
	for piece, count := range capturedPiecesBlack {
		whiteScore += count * pieceToValue[piece]
	}
}

func movePiece(src string, dest string) {
	fmt.Println(src, dest)
}

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
