package domain

import "testing"

func mustBoard(t *testing.T, w, h int) *Board {
	t.Helper()
	b, err := NewBoard(w, h)
	if err != nil {
		t.Fatalf("failed creating board: %v", err)
	}
	return b
}

func TestBoardSetAndGetCell(t *testing.T) {
	b := mustBoard(t, 10, 20)
	if err := b.SetCell(2, 3, Block); err != nil {
		t.Fatalf("SetCell returned error: %v", err)
	}
	if got := b.GetCell(2, 3); got != Block {
		t.Fatalf("expected 'X', got '%s'", got)
	}
}

func TestBoardIsCollision_NoCollision(t *testing.T) {
	b := mustBoard(t, 5, 5)
	shape := [][]Cell{
		{Block, Block},
		{Block, Block},
	}
	p, _ := NewPiece("O", shape, Block, Position{X: 1, Y: 1})
	if b.IsCollision(*p) {
		t.Fatalf("expected no collision, but got one")
	}
}

func TestBoardIsCollision_WithCollision(t *testing.T) {
	b := mustBoard(t, 5, 5)
	if err := b.SetCell(2, 2, Block); err != nil {
		t.Fatalf("SetCell returned error: %v", err)
	}

	shape := [][]Cell{
		{Block, Block},
		{Block, Block},
	}
	p, _ := NewPiece("O", shape, Block, Position{X: 1, Y: 1})
	if !b.IsCollision(*p) {
		t.Fatalf("expected collision, got none")
	}
}

func TestBoardClearFullLines(t *testing.T) {
	b := mustBoard(t, 4, 4)
	for x := 0; x < 4; x++ {
		if err := b.SetCell(x, 3, Block); err != nil {
			t.Fatalf("SetCell error: %v", err)
		}
	}
	linesCleared := b.ClearFullLines()
	if linesCleared != 1 {
		t.Fatalf("expected 1 line cleared, got %d", linesCleared)
	}
	for x := 0; x < 4; x++ {
		if b.GetCell(x, 3) != Empty {
			t.Fatalf("expected empty cell at bottom line, found '%s'", b.GetCell(x, 3))
		}
	}
}
