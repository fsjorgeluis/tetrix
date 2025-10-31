package domain

import (
	"errors"
)

// Cell represents if a matrix cell is blocked,
// maybe an alternative should be use bool instead
type Cell string

const (
	Empty Cell = "  "
	Block Cell = "[ ]"
)

// Board retains the state of the game.
// Cells[y][x] conventionally (row,col)
type Board struct {
	Width, Height int
	Cells         [][]Cell
	PlacedPieces  []Piece
}

// NewBoard creates a board with the given size
// and initializes all cells to ' ' (empty)
func NewBoard(width, height int) (*Board, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("invalid board size, width and height must be > 0")
	}
	c := make([][]Cell, height)
	for y := range height {
		c[y] = make([]Cell, width)
		for x := range width {
			c[y][x] = Empty
		}
	}
	return &Board{
		Width:        width,
		Height:       height,
		Cells:        c,
		PlacedPieces: []Piece{},
	}, nil
}

// SavePiece saves a new piece in the board, as a placed one
// so it gets rendered whatsoever every tick.
func (b *Board) SavePiece(piece Piece) {
	b.PlacedPieces = append(b.PlacedPieces, piece)
}

// GetCell returns the rune at (x,y).
// if the position is out of bounds, returns 0
func (b *Board) GetCell(x, y int) Cell {
	if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
		return Empty
	}
	return b.Cells[y][x]
}

// SetCell sets the rune at (x,y),
// returns an error if the position is out of bounds
func (b *Board) SetCell(x, y int, val Cell) error {
	if x < 0 || x >= b.Width || y < 0 || y >= b.Height {
		return errors.New("coordinates out of bounds")
	}
	b.Cells[y][x] = val
	return nil
}

// IsEmpty checks if the cell at (x,y) is empty
// if the position is out of bounds, returns false not empty, useful for collisions
func (b *Board) IsEmpty(x, y int) bool {
	if x < 0 || x >= b.Width {
		return false
	}
	if y >= b.Height {
		return false
	}
	if y < 0 {
		return true
	}
	return b.Cells[y][x] == Empty
}

// CanPlace checks if the piece can be placed on the board referencing its current position,
// returns false if the piece is colliding with the board.
// - occupied cells by the piece outside the board e.g. (x out or y >= Height) -> false
// - cells with y < 0 (out of bounds) considered valid
// this allows the pieces to spawn partially outside the board
func (b *Board) CanPlace(p *Piece) bool {
	for _, pos := range p.OccupiedCells() {
		if pos.X < 0 || pos.X >= b.Width || pos.Y >= b.Height {
			return false
		}
		if pos.Y >= 0 && b.Cells[pos.Y][pos.X] != Empty {
			return false
		}
	}
	for _, pos := range p.OccupiedCells() {
		if pos.Y < 0 {
			continue
		}
		if b.Cells[pos.Y][pos.X] != Empty {
			return false
		}
	}
	return true
}

// IsCollision returns true if the piece is colliding with the board,
// false otherwise.
// - occupied cells by the piece outside the board e.g. (x out or y >= Height) -> true
// - cells with y < 0 (out of bounds) considered valid
// this allows the pieces to spawn partially outside the board
func (b *Board) IsCollision(p Piece) bool {
	for _, pos := range p.OccupiedCells() {
		if pos.X < 0 || pos.X >= b.Width || pos.Y >= b.Height {
			return true
		}
		if pos.Y >= 0 && !b.IsEmpty(pos.X, pos.Y) {
			return true
		}
	}
	return false
}

// Merge places the piece on the board,
// assume CanPlace has been called before; however, a minimum check is performed
func (b *Board) Merge(p *Piece) error {
	if !b.CanPlace(p) {
		return errors.New("cannot merge piece: collision or out of bounds")
	}
	for _, pos := range p.OccupiedCells() {
		if pos.Y >= 0 && pos.Y < b.Height && pos.X >= 0 && pos.X < b.Width {
			b.Cells[pos.Y][pos.X] = Block
		}
	}
	return nil
}

// ClearFullLines TODO: change this simple implementation to a more efficient one
// ClearFullLines detects and clears full rows on the board,
// returns the number of cleared rows
func (b *Board) ClearFullLines() int {
	cleared := 0
	writeRow := b.Height - 1

	for y := b.Height - 1; y >= 0; y-- {
		full := true
		for x := 0; x < b.Width; x++ {
			if b.Cells[y][x] == Empty {
				full = false
				break
			}
		}
		if full {
			cleared++
			continue
		}
		copy(b.Cells[writeRow], b.Cells[y])
		writeRow--
	}

	// empty rows at the top
	for y := writeRow; y >= 0; y-- {
		for x := 0; x < b.Width; x++ {
			b.Cells[y][x] = Empty
		}
	}
	// if rows are not cleared, avoid reassigning to prevent unnecessary garbage collection
	if cleared > 0 {
		//log.Printf("Cleared %d line(s)", cleared)
	}
	return cleared
}

// IsGameOver determines if the board is in game over state
// Criteria: every cell is blocked at spawn point e.g.: superior rows are blocked
// if any cell in the first row is blocked, the game is over
func (b *Board) IsGameOver() bool {
	for x := 0; x < b.Width; x++ {
		if b.Cells[0][x] == Block {
			return true
		}
	}
	return false
}
