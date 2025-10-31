package usecase

import "github.com/fsjorgeluis/tetrix/internal/domain"

// MoveLeft try to move the piece to the left and revert if there is a collision.
func MoveLeft(b *domain.Board, p *domain.Piece) {
	p.Pos = p.Pos.MoveLeft()
	if b.IsCollision(*p) {
		p.Pos = p.Pos.MoveRight()
	}
}

// MoveRight try to move the piece to the right and revert if there is a collision.
func MoveRight(b *domain.Board, p *domain.Piece) {
	p.Pos = p.Pos.MoveRight()
	if b.IsCollision(*p) {
		p.Pos = p.Pos.MoveLeft()
	}
}

// MoveDown try to move the piece down and revert if there is a collision.
func MoveDown(b *domain.Board, p *domain.Piece) {
	p.Pos = p.Pos.MoveDown()
	if b.IsCollision(*p) {
		p.Pos = p.Pos.MoveUp()
	}
}

// RotateCW rotates the piece clockwise and reverts if there is a collision.
func RotateCW(b *domain.Board, p *domain.Piece) {
	p.RotateCW()
	if b.IsCollision(*p) {
		p.RotateCCW()
	}
}

// RotateCCW rotates the piece counter-clockwise and reverts if there is a collision.
func RotateCCW(b *domain.Board, p *domain.Piece) {
	p.RotateCCW()
	if b.IsCollision(*p) {
		p.RotateCW()
	}
}
