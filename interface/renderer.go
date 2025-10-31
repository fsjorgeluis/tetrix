package _interface

import "github.com/fsjorgeluis/tetrix/internal/domain"

type Renderer interface {
	Begin(board *domain.Board, score int)
	DrawPiece(board *domain.Board, piece *domain.Piece)
	DrawBoard(board *domain.Board)
	Flush()
	Clear()
	Close()
}
