package main

import "strconv"

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
		if rankInt == 7 {
			moves[file+strconv.Itoa(rankInt-2)] = struct{}{}
			// moves = appendString(moves, file+strconv.Itoa(rankInt-2))
		}
		moves[file+strconv.Itoa(rankInt-1)] = struct{}{}
		// moves = appendString(moves, file+strconv.Itoa(rankInt-1))
	}

	return moves
}
