package usecase

import (
	"reflect"
	"testing"

	"github.com/fsjorgeluis/tetrix/internal/domain"
)

func TestMoveLeftAndRight_NoCollision(t *testing.T) {
	b := newBoard()
	p := newPieceAt(5, 5)

	MoveLeft(b, p)
	if p.Pos.X != 4 {
		t.Errorf("expected X=4 after MoveLeft, got %d", p.Pos.X)
	}

	MoveRight(b, p)
	if p.Pos.X != 5 {
		t.Errorf("expected X=5 after MoveRight, got %d", p.Pos.X)
	}
}

func TestMoveLeft_RevertsOnCollision(t *testing.T) {
	b := newBoard()
	p := newPieceAt(0, 0) // already at left edge

	MoveLeft(b, p)
	if p.Pos.X != 0 {
		t.Errorf("expected no movement, got X=%d", p.Pos.X)
	}
}

func TestMoveDown_NoCollision(t *testing.T) {
	board, _ := domain.NewBoard(10, 20)
	p, _ := domain.NewPiece(
		"I",
		[][]domain.Cell{{domain.Block}, {domain.Block}, {domain.Block}},
		domain.Block,
		domain.Position{X: 3, Y: 0})

	before := p.Pos.Y
	MoveDown(board, p)
	if p.Pos.Y != before+1 {
		t.Errorf("expected Y to increase by 1, got %d", p.Pos.Y)
	}
}

func TestMoveDown_WithCollision(t *testing.T) {
	board, _ := domain.NewBoard(10, 20)
	p, _ := domain.NewPiece(
		"O",
		[][]domain.Cell{{domain.Block, domain.Block}, {domain.Block, domain.Block}},
		domain.Block,
		domain.Position{X: 1, Y: 17})
	err := board.Merge(p)
	if err != nil {
		t.Errorf("error merging piece: %s", err)
		return
	}

	// pieza encima
	q, _ := domain.NewPiece(
		"O",
		[][]domain.Cell{{domain.Block, domain.Block}, {domain.Block, domain.Block}},
		domain.Block,
		domain.Position{X: 0, Y: 17})
	MoveDown(board, q)
	if q.Pos.Y != 17 {
		t.Errorf("expected Y not to change due to collision, got %d", q.Pos.Y)
	}
}

func deepCopyShape(shape [][]domain.Cell) [][]domain.Cell {
	newShape := make([][]domain.Cell, len(shape))
	for i := range shape {
		newShape[i] = append([]domain.Cell(nil), shape[i]...)
	}
	return newShape
}

func TestRotateCWAndCCW_NoCollision(t *testing.T) {
	b := newBoard()
	shape := [][]domain.Cell{
		{domain.Block, domain.Empty},
		{domain.Block, domain.Empty},
		{domain.Block, domain.Block},
	}
	p, _ := domain.NewPiece(
		"L",
		shape,
		domain.Block,
		domain.Position{X: 3, Y: 3})

	originalShape := deepCopyShape(p.Shape)

	RotateCW(b, p)
	if reflect.DeepEqual(p.Shape, originalShape) {
		t.Errorf("expected rotation to change shape")
	}

	RotateCCW(b, p)
	if !reflect.DeepEqual(p.Shape, originalShape) {
		t.Errorf("expected rotation CCW to restore original shape")
	}
}
