package usecase

import (
	"log"

	"github.com/fsjorgeluis/tetrix/infrastructure"
	"github.com/fsjorgeluis/tetrix/internal/domain"
)

type GameService struct {
	board        *domain.Board
	currentPiece *domain.Piece
	nextPiece    *domain.Piece
	spawner      PieceSpawner
	score        int
	level        int
	gameOver     bool
	paused       bool
	sound        *infrastructure.SoundPlayer
}

func NewGameService(
	board *domain.Board,
	spawner PieceSpawner,
	sound *infrastructure.SoundPlayer,
) (*GameService, error) {
	if spawner == nil {
		spawner = NewDefaultSpawner()
	}

	gs := &GameService{
		board:   board,
		spawner: spawner,
		score:   0,
		level:   1,
		sound:   sound,
	}

	if err := gs.spawnNextPiece(); err != nil {
		return nil, err
	}

	return gs, nil
}

func (gs *GameService) spawnNextPiece() error {
	if gs.nextPiece == nil {
		// at first spawn, generate a new piece
		next, err := gs.spawner.NextPiece()
		if err != nil {
			return err
		}
		gs.nextPiece = next
	}
	// swap current piece with next piece
	gs.currentPiece = gs.nextPiece

	next, err := gs.spawner.NextPiece()
	if err != nil {
		return err
	}
	gs.nextPiece = next

	if gs.board.IsGameOver() {
		//log.Println("GAME OVER: spawn blocked")
		gs.gameOver = true
	}
	//if !gs.board.CanPlace(gs.currentPiece) {
	//	gs.gameOver = true
	//	log.Println("GAME OVER: cannot place new piece")
	//}
	return nil
}

func (gs *GameService) CurrentPiece() *domain.Piece {
	return gs.currentPiece
}

func (gs *GameService) NextPiece() *domain.Piece {
	return gs.nextPiece
}

func (gs *GameService) TogglePause() {
	gs.paused = !gs.paused
	// TODO: try to pause music
}

func (gs *GameService) Paused() bool {
	return gs.paused
}

// Tick advances the game state by one tick.
func (gs *GameService) Tick() {
	if gs.gameOver || gs.currentPiece == nil || gs.paused {
		return
	}
	hitBottom := Tick(gs.board, gs.currentPiece)
	if hitBottom {
		_ = gs.board.Merge(gs.currentPiece)
		//gs.board.SavePiece(*gs.currentPiece) // maybe is not needed
		linesCleared := gs.board.ClearFullLines()
		if linesCleared > 0 {
			gs.score += linesCleared
			if gs.sound != nil {
				go gs.sound.PlayEffect("assets/sounds/line_clear.mp3")
			}
		}
		_ = gs.spawnNextPiece()

		//if !gs.board.CanPlace(gs.currentPiece) {
		//	log.Println("GAME OVER: cannot place new piece")
		//	gs.gameOver = true
		//	return
		//}

		if gs.sound != nil {
			go gs.sound.PlayEffect("assets/sounds/shot.mp3")

		}
	}
}

func (gs *GameService) MoveLeft()  { MoveLeft(gs.board, gs.currentPiece) }
func (gs *GameService) MoveRight() { MoveRight(gs.board, gs.currentPiece) }
func (gs *GameService) RotateCW()  { RotateCW(gs.board, gs.currentPiece) }
func (gs *GameService) RotateCCW() { RotateCCW(gs.board, gs.currentPiece) }
func (gs *GameService) MoveDown()  { MoveDown(gs.board, gs.currentPiece) }
func (gs *GameService) Drop() {
	if gs.gameOver || gs.currentPiece == nil {
		return
	}

	Drop(gs.board, gs.currentPiece)
	_ = gs.board.Merge(gs.currentPiece)

	linesCleared := gs.board.ClearFullLines()
	if linesCleared > 0 {
		gs.score += linesCleared
		if gs.sound != nil {
			go gs.sound.PlayEffect("assets/sounds/line_clear.mp3")
		}
	}

	err := gs.spawnNextPiece()
	if err != nil {
		log.Println("GAME OVER: failed to spawn next piece:", err)
		gs.gameOver = true
		return
	}

	if gs.board.IsGameOver() {
		log.Println("GAME OVER: spawn blocked")
		gs.gameOver = true
	}
	//if !gs.board.CanPlace(gs.currentPiece) {
	//	log.Println("GAME OVER: cannot place new piece")
	//	gs.gameOver = true
	//	return
	//}
}

// GetBoard returns a copy of the board with the current piece drawn on top.
func (gs *GameService) GetBoard() [][]domain.Cell {
	copyBoard := make([][]domain.Cell, gs.board.Height)
	for y := 0; y < gs.board.Height; y++ {
		copyBoard[y] = make([]domain.Cell, gs.board.Width)
		copy(copyBoard[y], gs.board.Cells[y])
	}
	if gs.currentPiece != nil {
		for _, pos := range gs.currentPiece.OccupiedCells() {
			if pos.Y >= 0 && pos.Y < gs.board.Height && pos.X >= 0 && pos.X < gs.board.Width {
				copyBoard[pos.Y][pos.X] = gs.currentPiece.Symbol
			}
		}
	}
	return copyBoard
}

func (gs *GameService) IsGameOver() bool { return gs.gameOver }
func (gs *GameService) Score() int       { return gs.score }
func (gs *GameService) Level() int       { return gs.level }
func (gs *GameService) Reset() {
	for y := 0; y < gs.board.Height; y++ {
		for x := 0; x < gs.board.Width; x++ {
			gs.board.Cells[y][x] = domain.Empty
		}
	}
	gs.board.PlacedPieces = []domain.Piece{}

	gs.score = 0
	gs.level = 1
	gs.gameOver = false
	gs.currentPiece = nil
	gs.nextPiece = nil

	if err := gs.spawnNextPiece(); err != nil {
		log.Printf("failed to spawn first piece on reset: %v", err)
		gs.gameOver = true
	}
}
