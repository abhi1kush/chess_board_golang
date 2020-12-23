package bitboard

import (
	"math"
)

const (
	white              = true
	black              = false
	whitePieceIdxStart = 0
	whitePieceIdxEnd   = 15
	blackPieceIdxStart = 16
	blackPieceIdxEnd   = 31
)

// Unique pieces.
const (
	king = iota
	knight
	queen
	rook
	bishop
	pawn
)

//Pieces index
const (
	rook1W = iota
	knight1W
	bBishopW
	queenW
	kingW
	wBishopW
	knight2W
	rook2W
	aPawnW
	bPawnW
	cPawnW
	dPawnW
	ePawnW
	fPawnW
	gPawnW
	hPawnW
	aPawnB
	bPawnB
	cPawnB
	dPawnB
	ePawnB
	fPawnB
	gPawnB
	hPawnB
	rook1B
	knight1B
	bBishopB
	queenB
	kingB
	wBishopB
	knight2B
	rook2B
	pieceCount
)

type bitboard struct {
	pieces [pieceCount]uint64
}

type board struct {
	sideToMove        bool
	enPassantAvlbl    bool
	wCastelLongAvlbl  bool
	wCastelShortAvlbl bool
	bCastelLongAvlbl  bool
	bCastelShortAvlbl bool
	bitboard
}

func (b bitboard) initBitBoard() {
	var one uint64 = 1
	b.pieces[rook1W] = one
	b.pieces[knight1W] = one >> 1
	b.pieces[bBishopW] = one >> 2
	b.pieces[queenW] = one >> 3
	b.pieces[kingW] = one >> 4
	b.pieces[wBishopW] = one >> 5
	b.pieces[knight2W] = one >> 6
	b.pieces[rook2W] = one >> 7
	b.pieces[aPawnW] = one >> 8
	b.pieces[bPawnW] = one >> 9
	b.pieces[cPawnW] = one >> 10
	b.pieces[dPawnW] = one >> 11
	b.pieces[ePawnW] = one >> 12
	b.pieces[fPawnW] = one >> 13
	b.pieces[gPawnW] = one >> 14
	b.pieces[hPawnW] = one >> 15
	var maxbit uint64 = uint64(math.Pow(2, 63))
	b.pieces[rook2B] = maxbit
	b.pieces[knight2B] = maxbit << 1
	b.pieces[bBishopB] = maxbit << 2
	b.pieces[kingB] = maxbit << 3
	b.pieces[queenB] = maxbit << 4
	b.pieces[wBishopB] = maxbit << 5
	b.pieces[knight1B] = maxbit << 6
	b.pieces[rook1B] = maxbit << 7
	b.pieces[hPawnB] = maxbit << 8
	b.pieces[gPawnB] = maxbit << 9
	b.pieces[fPawnB] = maxbit << 10
	b.pieces[ePawnB] = maxbit << 11
	b.pieces[dPawnB] = maxbit << 12
	b.pieces[cPawnB] = maxbit << 13
	b.pieces[bPawnB] = maxbit << 14
	b.pieces[aPawnB] = maxbit << 15
}

func (b board) init() {
	b.sideToMove = white
	b.enPassantAvlbl = false
	b.wCastelLongAvlbl = true
	b.wCastelShortAvlbl = true
	b.bCastelLongAvlbl = true
	b.bCastelShortAvlbl = true
	b.bitboard.initBitBoard()
}

func bitToSquare(bitPos uint64) uint64 {
	return uint64(math.Logb(float64(bitPos)))
}

func squareToBit(squareNumber uint64) uint64 {
	var one uint64 = 1
	if squareNumber == 1 {
		return one
	}
	return one >> (squareNumber - 1)
}
