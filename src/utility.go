package main

import (
	"strconv"
)

// const StartingFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
const StartingFen = "8/8/8/4K3/8/8/8/8 b - - 0 1"

var fileToOffset = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
}

var offsetToFile = map[int]string{
	0: "a",
	1: "b",
	2: "c",
	3: "d",
	4: "e",
	5: "f",
	6: "g",
	7: "h",
}

// Append string to string slice.
func appendString(slice []string, data ...string) []string {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]string, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func getPieceColor(piece string) string {
	if piece == "P" || piece == "R" ||
		piece == "B" || piece == "N" ||
		piece == "Q" || piece == "K" {
		return "white"
	} else if piece == "p" || piece == "r" ||
		piece == "b" || piece == "n" ||
		piece == "q" || piece == "k" {
		return "black"
	}
	return "blank"
}

func getPos(i int, j int) string {
	rank := i + 1
	file := offsetToFile[j]
	return file + strconv.Itoa(rank)
}
