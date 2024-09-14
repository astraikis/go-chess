package main

import (
	"fmt"
	"strconv"
)

func canAttack(piece string, srcX int, srcY int, targetX int, targetY int) bool {
	// White pawn
	if piece == "P" {
		if (targetX == srcX+1 || targetX == srcX-1) && targetY == srcY-1 {
			fmt.Println(getPieceColor(string(getPieceByPos(targetY, targetX))))
			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				return true
			}
		}
	}

	return false
}

func isValidMove(fen string, move string) {

}

func getPieceByPos(rank int, file int) byte {
	return board[((rank-1)*8)+file]
}

func getLegalMoves(pos string) map[string]struct{} {
	var rank = string(pos[1])
	rankInt, err := strconv.Atoi(string(rank))
	if err != nil {
		panic(err)
	}

	var file = string(pos[0])
	var fileOffset = fileToOffset[file]

	moves := make(map[string]struct{})
	var piece = getPieceByPos(rankInt, fileOffset)

	if string(piece) == "P" {
		// White pawn
		// Forward
		if rankInt == 7 {
			moves[file+strconv.Itoa(rankInt-2)] = struct{}{}
		}
		moves[file+strconv.Itoa(rankInt-1)] = struct{}{}

		// Capture
		if canAttack(string(piece), fileOffset, rankInt, fileOffset-1, rankInt-1) {
			moves[offsetToFile[fileOffset-1]+strconv.Itoa(rankInt-1)] = struct{}{}
		}
		if canAttack(string(piece), fileOffset, rankInt, fileOffset+1, rankInt-1) {
			moves[offsetToFile[fileOffset+1]+strconv.Itoa(rankInt-1)] = struct{}{}
		}

		// En passant
	} else if string(piece) == "R" {
		// White rook
		// Up
		targetY := rankInt - 1
		targetX := fileOffset
		for {
			if targetY == 0 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
				targetY--
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Right
		targetY = rankInt
		targetX = fileOffset + 1
		for {
			if targetX == 8 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetX++
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Down
		targetY = rankInt + 1
		targetX = fileOffset
		for {
			if targetY == 9 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
				targetY++
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Left
		targetY = rankInt
		targetX = fileOffset - 1
		for {
			if targetX == -1 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetX--
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}
	} else if string(piece) == "B" {
		// White bishop
		// Up right
		targetY := rankInt - 1
		targetX := fileOffset + 1
		for {
			if targetY == 0 || targetX == 8 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetY--
				targetX++
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Down right
		targetY = rankInt + 1
		targetX = fileOffset + 1
		for {
			if targetY == 9 || targetX == 8 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetY++
				targetX++
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Down left
		targetY = rankInt + 1
		targetX = fileOffset - 1
		for {
			if targetY == 9 || targetX == -1 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetY++
				targetX--
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Up left
		targetY = rankInt - 1
		targetX = fileOffset - 1
		for {
			if targetY == 0 || targetX == -1 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetY--
				targetX--
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}
	} else if string(piece) == "Q" {
		// Up
		targetY := rankInt - 1
		targetX := fileOffset
		for {
			if targetY == 0 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
				targetY--
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Right
		targetY = rankInt
		targetX = fileOffset + 1
		for {
			if targetX == 8 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetX++
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Down
		targetY = rankInt + 1
		targetX = fileOffset
		for {
			if targetY == 9 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
				targetY++
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Left
		targetY = rankInt
		targetX = fileOffset - 1
		for {
			if targetX == -1 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetX--
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Up right
		targetY = rankInt - 1
		targetX = fileOffset + 1
		for {
			if targetY == 0 || targetX == 8 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetY--
				targetX++
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Down right
		targetY = rankInt + 1
		targetX = fileOffset + 1
		for {
			if targetY == 9 || targetX == 8 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetY++
				targetX++
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Down left
		targetY = rankInt + 1
		targetX = fileOffset - 1
		for {
			if targetY == 9 || targetX == -1 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetY++
				targetX--
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}

		// Up left
		targetY = rankInt - 1
		targetX = fileOffset - 1
		for {
			if targetY == 0 || targetX == -1 {
				break
			}

			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				targetY--
				targetX--
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[offsetToFile[targetX]+strconv.Itoa(targetY)] = struct{}{}
				break
			} else {
				break
			}
		}
	} else if string(piece) == "K" {
		// White king
		// Up
		targetY := rankInt - 1
		targetX := fileOffset

		if targetY != 0 {
			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
				targetY--
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
			}
		}

		// Up right
		targetY = rankInt - 1
		targetX = fileOffset + 1

		if targetY != 0 && targetX != 8 {
			if getPieceColor(string(getPieceByPos(targetY, targetX))) == "blank" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
				targetY--
			} else if getPieceColor(string(getPieceByPos(targetY, targetX))) == "black" {
				moves[file+strconv.Itoa(targetY)] = struct{}{}
			}
		}
	}

	return moves
}
