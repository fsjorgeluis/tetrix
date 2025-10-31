package infrastructure

import(
	"github.com/fsjorgeluis/tetrix/internal/domain"
	"testing"
)

func TestTCellRenderer_renderPiece(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		board *domain.Board
		piece *domain.Piece
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := NewTCellRenderer()
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			r.RenderPiece(tt.board, tt.piece)
		})
	}
}

