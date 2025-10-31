package usecase

import (
	"github.com/fsjorgeluis/tetrix/internal/domain"
)

// Tick moves the piece down and attach the pieces if collision detected.
func Tick(b *domain.Board, current *domain.Piece) bool {
	if current == nil {
		return false
	}
	current.Pos = current.Pos.MoveDown()
	if b.IsCollision(*current) {
		current.Pos = current.Pos.MoveUp()
		_ = b.Merge(current)
		_ = b.ClearFullLines()
		return true
	}
	return false
}

// Drop grants the piece to drop until it collides with the bottom of the board.
func Drop(b *domain.Board, current *domain.Piece) {
	if current == nil {
		return
	}
	for {
		current.Pos = current.Pos.MoveDown()
		if b.IsCollision(*current) {
			current.Pos = current.Pos.MoveUp()
			_ = b.Merge(current)
			_ = b.ClearFullLines()
			break
		}
	}
}
