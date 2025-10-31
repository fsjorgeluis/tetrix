package _interface

import "github.com/fsjorgeluis/tetrix/internal/domain"

type Renderer interface {
	Begin(board *domain.Board, score int)
	DrawPiece(board *domain.Board, piece *domain.Piece)
	DrawNextPiece(piece *domain.Piece, x, y int)
	DrawBoard(board *domain.Board)
	DrawLabel(x, y int, label string)
	Flush()
	Clear()
	Close()
}
