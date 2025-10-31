package usecase

import (
	"testing"

	"github.com/fsjorgeluis/tetrix/infrastructure"
	"github.com/fsjorgeluis/tetrix/internal/domain"
)

func TestNewGameService_InitialState(t *testing.T) {
	spawner := NewDefaultSpawner()
	soundPlayer := infrastructure.NewSoundPlayer()
	board, err := domain.NewBoard(10, 10)
	if err != nil {
		panic(err)
	}
	game, err := NewGameService(board, spawner, soundPlayer)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if game.IsGameOver() {
		t.Errorf("expected game not over at start")
	}
	if game.Level() != 1 {
		t.Errorf("expected level 1 at start, got %d", game.Level())
	}
	boardState := game.GetBoard()
	if len(boardState) != 10 || len(boardState[0]) != 10 {
		t.Errorf("unexpected board size")
	}
}

func TestGameService_TickAndSpawn(t *testing.T) {
	soundPlayer := infrastructure.NewSoundPlayer()
	board, err := domain.NewBoard(10, 10)
	if err != nil {
		panic(err)
	}
	game, _ := NewGameService(board, NewDefaultSpawner(), soundPlayer)

	initialBoard := game.GetBoard()
	game.Tick()
	newBoard := game.GetBoard()

	same := true
	for y := range initialBoard {
		for x := range initialBoard[y] {
			if initialBoard[y][x] != newBoard[y][x] {
				same = false
				break
			}
		}
	}
	if same {
		t.Errorf("expected board to change after Tick")
	}
}

func TestGameService_MoveAndDrop(t *testing.T) {
	soundPlayer := infrastructure.NewSoundPlayer()
	board, err := domain.NewBoard(10, 10)
	if err != nil {
		panic(err)
	}
	game, _ := NewGameService(board, NewDefaultSpawner(), soundPlayer)

	game.MoveLeft()
	game.MoveRight()
	game.RotateCW()
	game.RotateCCW()
	game.Drop()

	if game.IsGameOver() {
		t.Errorf("game should not be over right after drop")
	}
}

func TestScoreIncrement(t *testing.T) {
	board, _ := domain.NewBoard(10, 20)
	gs, _ := NewGameService(board, nil, nil)

	for x := 0; x < board.Width; x++ {
		board.Cells[19][x] = domain.Block
	}

	linesCleared := board.ClearFullLines()
	if linesCleared != 1 {
		t.Errorf("expected 1 line cleared, got %d", linesCleared)
	}

	gs.score += linesCleared

	if gs.Score() != 1 {
		t.Errorf("expected score 1, got %d", gs.Score())
	}
}
