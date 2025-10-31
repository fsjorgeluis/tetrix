package domain

import "testing"

func TestPositionMoveDown(t *testing.T) {
	pos := Position{X: 5, Y: 5}
	newPos := pos.MoveDown()
	if newPos.Y != 6 || newPos.X != 5 {
		t.Errorf("expected (5,6), got (%d,%d)", newPos.X, newPos.Y)
	}
}

func TestPositionMoveLeftRight(t *testing.T) {
	pos := Position{X: 3, Y: 2}
	left := pos.MoveLeft()
	if left.X != 2 || left.Y != 2 {
		t.Fatalf("expected (2,2), got (%d,%d)", left.X, left.Y)
	}
	right := pos.MoveRight()
	if right.X != 4 || right.Y != 2 {
		t.Fatalf("expected (4,2), got (%d,%d)", right.X, right.Y)
	}
}

func TestPiecePositionBoardIsCollision_NoCollision(t *testing.T) {
	b, _ := NewBoard(5, 5)
	piece := Piece{
		Shape: [][]Cell{
			{Block, Block},
			{Block, Block},
		},
		Pos: Position{X: 1, Y: 1},
	}
	if b.IsCollision(piece) {
		t.Errorf("expected no collision, but got one")
	}
}

func TestPiecePositionBoardIsCollision_WithCollision(t *testing.T) {
	b, _ := NewBoard(5, 5)
	err := b.SetCell(2, 2, Block)
	if err != nil {
		t.Errorf("error setting cell: %s", err)
		return
	}

	piece := Piece{
		Shape: [][]Cell{
			{Block, Block},
			{Block, Block},
		},
		Pos: Position{X: 1, Y: 1},
	}
	if !b.IsCollision(piece) {
		t.Errorf("expected collision, got none")
	}
}

func TestPiecePositionBoardClearFullLines(t *testing.T) {
	b, _ := NewBoard(4, 4)
	for x := 0; x < 4; x++ {
		err := b.SetCell(x, 3, Block)
		if err != nil {
			t.Errorf("error setting cell: %s", err)
			return
		}
	}
	linesCleared := b.ClearFullLines()
	if linesCleared != 1 {
		t.Errorf("expected 1 line cleared, got %d", linesCleared)
	}
	for x := 0; x < 4; x++ {
		if b.GetCell(x, 3) != Empty {
			t.Errorf("expected empty cell at bottom line, found '%s'", b.GetCell(x, 3))
		}
	}
}
