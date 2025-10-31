package usecase

import "github.com/fsjorgeluis/tetrix/internal/domain"

func newBoard() *domain.Board {
	b, _ := domain.NewBoard(10, 20)
	return b
}

func newPieceAt(x, y int) *domain.Piece {
	shape := [][]domain.Cell{{domain.Block, domain.Block}}
	p, _ := domain.NewPiece("test", shape, domain.Block, domain.Position{X: x, Y: y})
	return p
}
