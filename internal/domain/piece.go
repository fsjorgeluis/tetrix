package domain

import "errors"

// Piece is an abstract piece.
// - Shape is a square matrix (NxN) of 0/1 where 1 = blocked
// square matrix to easily rotate and move
// - Pos is the top-left position of the Shape relative to the board
// - ID is the string identifier the type (I, O, T, S, Z, J, L) useful for
// testing and networking
// Symbol is the representative rune for rendering (agnostic representation)
type Piece struct {
	Shape  [][]Cell
	Pos    Position
	ID     string
	Symbol Cell
}

var Tetrominoes = []struct {
	ID     string
	Shape  [][]Cell
	Symbol Cell
}{
	{"O", [][]Cell{{Block, Block}, {Block, Block}}, Block},
	{"I", [][]Cell{{Block}, {Block}, {Block}, {Block}}, Block},
	{"T", [][]Cell{{Block, Block, Block}, {Empty, Block, Empty}}, Block},
	{"S", [][]Cell{{Empty, Block, Block}, {Block, Block, Empty}}, Block},
	{"Z", [][]Cell{{Block, Block, Empty}, {Empty, Block, Block}}, Block},
	{"J", [][]Cell{{Block, Empty, Empty}, {Block, Block, Block}}, Block},
	{"L", [][]Cell{{Empty, Empty, Block}, {Block, Block, Block}}, Block},
}

// NewPiece creates a piece from a square matrix,
// also validates the mandatory minimum shape of 1x1
func NewPiece(id string, shape [][]Cell, symbol Cell, pos Position) (*Piece, error) {
	if len(shape) == 0 {
		return nil, errors.New("shape cannot be empty")
	}
	w := len(shape[0])
	for _, row := range shape {
		if len(row) != w {
			return nil, errors.New("all rows in shape must have the same width (NxN)")
		}
	}
	p := &Piece{
		Shape:  copyShape(shape),
		Pos:    pos,
		ID:     id,
		Symbol: symbol,
	}
	p.normalize()
	return p, nil
}

// copyShape makes a deep copy of the matrix shape.
func copyShape(s [][]Cell) [][]Cell {
	h := len(s)
	out := make([][]Cell, h)
	for i := range s {
		out[i] = make([]Cell, len(s[i]))
		copy(out[i], s[i])
	}
	return out
}

func (p *Piece) normalize() {
	h := len(p.Shape)
	w := len(p.Shape[0])
	// check empty top rows
	top := 0
topLoop:
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if p.Shape[r][c] == Block {
				break topLoop
			}
		}
		top++
	}
	// check empty bottom rows
	bottom := h - 1
bottomLoop:
	for r := h - 1; r >= 0; r-- {
		for c := 0; c < w; c++ {
			if p.Shape[r][c] == Block {
				break bottomLoop
			}
		}
		bottom--
	}
	// check empty left columns
	left := 0
leftLoop:
	for c := 0; c < w; c++ {
		for r := 0; r < h; r++ {
			if p.Shape[r][c] == Block {
				break leftLoop
			}
		}
		left++
	}
	// check empty right columns
	right := w - 1
rightLoop:
	for c := w - 1; c >= 0; c-- {
		for r := 0; r < h; r++ {
			if p.Shape[r][c] == Block {
				break rightLoop
			}
		}
		right--
	}

	if top > bottom || left > right {
		return
	}

	newH := bottom - top + 1
	newW := right - left + 1
	newShape := make([][]Cell, newH)
	for r := 0; r < newH; r++ {
		newShape[r] = make([]Cell, newW)
		for c := 0; c < newW; c++ {
			newShape[r][c] = p.Shape[top+r][left+c]
		}
	}

	// update position: after removing empty rows/columns,
	// the top-left position must be adjusted
	p.Pos.X += left
	p.Pos.Y += top
	p.Shape = newShape
}

// RotateCW TODO: maybe change the simple rotation system to a more complex one
// like Super Rotation System (SRS)
// RotateCW rotates the piece 90 degrees clockwise
func (p *Piece) RotateCW() {
	h := len(p.Shape)
	if h == 0 {
		return
	}
	w := len(p.Shape[0])
	rot := make([][]Cell, w)
	for i := range rot {
		rot[i] = make([]Cell, h)
	}
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			rot[c][h-1-r] = p.Shape[r][c]
		}
	}
	p.Shape = rot
}

// RotateCCW TODO: maybe change the simple rotation system to a more complex one
// like Super Rotation System (SRS)
// RotateCCW rotates the piece 90 degrees counter-clockwise
func (p *Piece) RotateCCW() {
	h := len(p.Shape)
	if h == 0 {
		return
	}
	w := len(p.Shape[0])
	rot := make([][]Cell, w)
	for i := range rot {
		rot[i] = make([]Cell, h)
	}
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			rot[w-1-c][r] = p.Shape[r][c]
		}
	}
	p.Shape = rot
}

// OccupiedCells returns the absolute positions of all occupied cells
// by the piece in the board.
// Used for collision detection and merge.
func (p *Piece) OccupiedCells() []Position {
	var out []Position
	for r := 0; r < len(p.Shape); r++ {
		for c := 0; c < len(p.Shape[r]); c++ {
			if p.Shape[r][c] == Block {
				out = append(out, Position{
					X: p.Pos.X + c,
					Y: p.Pos.Y + r,
				})
			}
		}
	}
	return out
}
