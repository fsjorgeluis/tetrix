package usecase

import (
	"math/rand"
	"time"

	"github.com/fsjorgeluis/tetrix/internal/domain"
)

type PieceSpawner interface {
	NextPiece() (*domain.Piece, error)
}

type DefaultSpawner struct {
	rng *rand.Rand
}

func NewDefaultSpawner() *DefaultSpawner {
	return &DefaultSpawner{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *DefaultSpawner) NextPiece() (*domain.Piece, error) {
	idx := s.rng.Intn(len(domain.Tetrominoes))
	p := domain.Tetrominoes[idx]
	// Pos: starter position is at the top of the board (y=-len(shape))
	pos := domain.Position{X: 3, Y: -len(p.Shape)}
	return domain.NewPiece(p.ID, p.Shape, p.Symbol, pos)
}
