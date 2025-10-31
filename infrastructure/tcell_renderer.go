package infrastructure

import (
	"github.com/fsjorgeluis/tetrix/internal/domain"
	"github.com/gdamore/tcell/v2"
)

type TCellRenderer struct {
	screen tcell.Screen
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
	return &TCellRenderer{screen: s}, nil
}

func (r *TCellRenderer) RenderBegin(board *domain.Board) {
	r.Clear()

	width, height := board.Width, board.Height

	// dibujar bordes
	for y := range height {
		r.drawEmptyCell(0, y)
		r.drawEmptyCell(width-1, y)
	}
	for x := range width {
		r.drawEmptyCell(x, 0)
		r.drawEmptyCell(x, height)
	}

	for y := range height {
		for x := range width {
			if board.Cells[y][x] != domain.Empty {
				r.drawBlock(x, y)
			}
		}
	}
}

func (r *TCellRenderer) RenderBoard(board *domain.Board) {
	for _, piece := range board.PlacedPieces {
		r.RenderPiece(board, &piece)
	}
}

func (r *TCellRenderer) RenderPiece(board *domain.Board, piece *domain.Piece) {
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

func (r *TCellRenderer) RenderEnd() {
	r.screen.Show()
}

// drawBlock draws the piece shape
func (r *TCellRenderer) drawBlock(x, y int) {
	screenX := x * 3
	style := tcell.StyleDefault.Foreground(tcell.ColorYellow)
	r.screen.SetContent(screenX, y, '[', nil, style)
	r.screen.SetContent(screenX+1, y, '█', nil, style)
	r.screen.SetContent(screenX+2, y, ']', nil, style)
}

// drawEmptyCell draws an empty cell like this: [ ]
func (r *TCellRenderer) drawEmptyCell(x, y int) {
	screenX := x * 3
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	r.screen.SetContent(screenX, y, '[', nil, style)
	r.screen.SetContent(screenX+1, y, ' ', nil, style)
	r.screen.SetContent(screenX+2, y, ']', nil, style)
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
