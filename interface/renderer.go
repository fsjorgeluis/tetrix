package _interface

import (
	"github.com/fsjorgeluis/tetrix/internal/domain"
)

type GameLabel struct {
	X, Y  int
	Text  string
	Color string
}

type Renderer interface {
	Begin(board *domain.Board, score int)
	DrawPiece(board *domain.Board, piece *domain.Piece)
	DrawNextPiece(piece *domain.Piece, x, y int)
	DrawBoard(board *domain.Board)
	DrawLabel(label *GameLabel)
	Flush()
	Clear()
	Close()
}
