package domain

// Position represents a coordinate on the board
// X increases to the right; Y increases downwards
// Origin (0, 0) top left corner of the board
type Position struct {
	X int
	Y int
}

func (p Position) MoveBy(dx, dy int) Position {
	return Position{X: p.X + dx, Y: p.Y + dy}
}

// MoveDown return a new Position moved down one unit.
func (p Position) MoveDown() Position {
	return p.MoveBy(0, 1)
}

func (p Position) MoveUp() Position {
	return p.MoveBy(0, -1)
}

// MoveLeft returns a new Position moved left one unit.
func (p Position) MoveLeft() Position {
	return p.MoveBy(-1, 0)
}

// MoveRight returns a new Position moved right one unit.
func (p Position) MoveRight() Position {
	return p.MoveBy(1, 0)
}
