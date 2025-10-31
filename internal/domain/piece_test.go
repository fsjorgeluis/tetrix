package domain

import (
	"reflect"
	"testing"
)

func TestRotateCWRectangular(t *testing.T) {
	shape := [][]Cell{
		{Block, Block, Empty},
		{Empty, Block, Block},
	}
	p, err := NewPiece("Z", shape, Block, Position{X: 0, Y: 0})
	if err != nil {
		t.Fatalf("unexpected error creating piece: %v", err)
	}

	p.RotateCW()

	expected := [][]Cell{
		{Empty, Block},
		{Block, Block},
		{Block, Empty},
	}

	if !reflect.DeepEqual(p.Shape, expected) {
		t.Fatalf("RotateCW failed. expected %v, got %v", expected, p.Shape)
	}
}

func TestRotateCCWRectangular(t *testing.T) {
	shape := [][]Cell{
		{Block, Block, Empty},
		{Empty, Block, Block},
	}
	p, err := NewPiece("Z", shape, Block, Position{X: 0, Y: 0})
	if err != nil {
		t.Fatalf("unexpected error creating piece: %v", err)
	}

	p.RotateCCW()

	expected := [][]Cell{
		{Empty, Block},
		{Block, Block},
		{Block, Empty},
	}

	if !reflect.DeepEqual(p.Shape, expected) {
		t.Fatalf("RotateCCW failed. expected %v, got %v", expected, p.Shape)
	}
}
