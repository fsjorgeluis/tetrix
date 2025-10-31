package usecase

import (
	"testing"

	"github.com/fsjorgeluis/tetrix/internal/domain"
)

func TestDefaultSpawner_NextPiece(t *testing.T) {
	spawner := NewDefaultSpawner()

	piece, err := spawner.NextPiece()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if piece == nil {
		t.Fatalf("expected piece, got nil")
	}
	if len(piece.Shape) == 0 {
		t.Errorf("expected piece shape not empty")
	}
	if piece.Symbol == domain.Empty {
		t.Errorf("expected valid symbol")
	}
	if piece.Pos.Y >= 0 {
		t.Errorf("expected initial Y negative (above board), got %d", piece.Pos.Y)
	}
}
