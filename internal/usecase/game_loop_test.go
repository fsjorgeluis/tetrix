package usecase

import (
	"testing"

	"github.com/fsjorgeluis/tetrix/internal/domain"
)

func TestTick_FixesPieceOnCollision(t *testing.T) {
	b := newBoard()
	// put piece at bottom of board
	// so it will collide with the bottom and merge with the piece above
	p := newPieceAt(0, b.Height-1)

	fixed := Tick(b, p)
	if !fixed {
		t.Errorf("expected piece fixed, got %v", fixed)
	}
}

func TestDrop_SettlesPieceAtBottom(t *testing.T) {
	b := newBoard()
	shape := [][]domain.Cell{
		{domain.Block, domain.Block},
		{domain.Block, domain.Block},
	}
	p, _ := domain.NewPiece("O", shape, domain.Block, domain.Position{X: 5, Y: 0})

	Drop(b, p)

	pieceHeight := len(p.Shape)
	expectedBottomY := b.Height - 2
	actualBottomY := p.Pos.Y + pieceHeight - 2

	if actualBottomY != expectedBottomY {
		t.Fatalf("expected piece bottom to be %d, got %d (p.Pos.Y=%d, pieceHeight=%d)",
			expectedBottomY, actualBottomY, p.Pos.Y, pieceHeight)
	}
}
