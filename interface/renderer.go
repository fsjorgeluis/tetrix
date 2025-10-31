package _interface

import "github.com/fsjorgeluis/tetrix/internal/domain"

type Renderer interface {
	Render(board *domain.Board, piece *domain.Piece)
	Clear()
	Close()
}
