package usecase


import (
	"github.com/fsjorgeluis/tetrix/infrastructure"
	"testing"
)

func TestNewGameService_InitialState(t *testing.T) {
	spawner := NewDefaultSpawner()
	sound_player := infrastructure.NewSoundPlayer()
	game, err := NewGameService(10, 20, spawner, sound_player)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if game.IsGameOver() {
		t.Errorf("expected game not over at start")
	}
	if game.Level() != 1 {
		t.Errorf("expected level 1 at start, got %d", game.Level())
	}
	board := game.GetBoard()
	if len(board) != 20 || len(board[0]) != 10 {
		t.Errorf("unexpected board size")
	}
}

func TestGameService_TickAndSpawn(t *testing.T) {
	sound_player := infrastructure.NewSoundPlayer()
	game, _ := NewGameService(10, 20, NewDefaultSpawner(), sound_player)

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
	sound_player := infrastructure.NewSoundPlayer()
	game, _ := NewGameService(10, 20, NewDefaultSpawner(), sound_player)

	game.MoveLeft()
	game.MoveRight()
	game.RotateCW()
	game.RotateCCW()
	game.Drop()

	if game.IsGameOver() {
		t.Errorf("game should not be over right after drop")
	}
}
