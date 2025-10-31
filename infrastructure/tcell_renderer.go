package infrastructure

import (
	"fmt"

	"github.com/fsjorgeluis/tetrix/internal/domain"
	"github.com/gdamore/tcell/v2"
)

type TCellRenderer struct {
	screen           tcell.Screen
	offsetX, offsetY int
}

func NewTCellRenderer() (*TCellRenderer, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := s.Init(); err != nil {
		return nil, err
	}
	s.Clear()
	return &TCellRenderer{
		screen:  s,
		offsetX: 2,
		offsetY: 1,
	}, nil
}

func (r *TCellRenderer) Begin(board *domain.Board, score int) {
	r.Clear()

	width, height := board.Width, board.Height

	// draw the borders
	for y := range height {
		r.drawEmptyCell(-1, y)
		r.drawEmptyCell(width, y+1)
	}
	for x := range width {
		r.drawEmptyCell(x, 0)
		r.drawEmptyCell(x, height)
	}

	// draw the cells
	for y := range height {
		for x := range width {
			if board.Cells[y][x] != domain.Empty {
				r.drawBlock(x, y)
			}
		}
	}

	// draw score
	r.DrawScore(score, width)
}

func (r *TCellRenderer) DrawBoard(board *domain.Board) {
	for _, piece := range board.PlacedPieces {
		r.DrawPiece(board, &piece)
	}
}

func (r *TCellRenderer) DrawPiece(board *domain.Board, piece *domain.Piece) {
	width, height := board.Width, board.Height

	if piece == nil {
		return
	}

	for dy, row := range piece.Shape {
		for dx, c := range row {
			if c == domain.Block {
				screenX := piece.Pos.X + dx
				screenY := piece.Pos.Y + dy
				if screenY >= 0 && screenY < height && screenX >= 0 && screenX < width {
					r.drawBlock(screenX, screenY)
				}
			}
		}
	}
}

func (r *TCellRenderer) DrawScore(score, boardWidth int) {
	label := "SCORE: "
	x := boardWidth * 3 // starts after the board
	y := 0              // top-right corner

	style := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	for i, ch := range label {
		r.screen.SetContent(x+i, y, ch, nil, style)
	}

	scoreStr := fmt.Sprintf("%d", score)
	for i, ch := range scoreStr {
		r.screen.SetContent(x+len(label)+i, y, ch, nil, style)
	}
}

func (r *TCellRenderer) Flush() {
	r.screen.Show()
}

// drawBlock draws the piece shape
func (r *TCellRenderer) drawBlock(x, y int) {
	screenX := (x + r.offsetX) * 3
	screenY := y + r.offsetY
	style := tcell.StyleDefault.Foreground(tcell.ColorYellow)
	r.screen.SetContent(screenX, screenY, '[', nil, style)
	r.screen.SetContent(screenX+1, screenY, '█', nil, style)
	r.screen.SetContent(screenX+2, screenY, ']', nil, style)
}

// drawEmptyCell draws an empty cell like this: [ ]
func (r *TCellRenderer) drawEmptyCell(x, y int) {
	screenX := (x + r.offsetX) * 3
	screenY := y + r.offsetY
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	r.screen.SetContent(screenX, screenY, '[', nil, style)
	r.screen.SetContent(screenX+1, screenY, ' ', nil, style)
	r.screen.SetContent(screenX+2, screenY, ']', nil, style)
}

func (r *TCellRenderer) Screen() tcell.Screen {
	return r.screen
}

func (r *TCellRenderer) Clear() {
	r.screen.Clear()
}

func (r *TCellRenderer) Close() {
	r.screen.Fini()
}
