package _interface

import "github.com/fsjorgeluis/tetrix/internal/domain"

type Renderer interface {
	RenderBegin(board *domain.Board)
	RenderPiece(board *domain.Board, piece *domain.Piece)
	RenderBoard(board *domain.Board)
	RenderEnd()
	Clear()
	Close()
}
